package graphdb

import (
	"context"
	"fmt"
	"jk-api/internal/config"
	adapter "jk-api/pkg/repository/adapter/graphdb"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type graphRepository struct {
	driver        neo4j.DriverWithContext
	sessionConfig neo4j.SessionConfig
	ctx           context.Context

	statements []string
	params     map[string]interface{}
}

// --- Constructor ---
func NewGraphRepository() adapter.GraphRepository {
	return &graphRepository{
		driver:        config.GetNeo4j(),
		sessionConfig: neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite},
		ctx:           context.Background(),
		params:        make(map[string]interface{}),
		statements:    []string{},
	}
}

// --- Clone helper untuk chainable pattern ---
func (r *graphRepository) clone() *graphRepository {
	clone := *r
	clone.params = make(map[string]interface{})
	for k, v := range r.params {
		clone.params[k] = v
	}
	clone.statements = append([]string{}, r.statements...)
	return &clone
}

// --- Builder methods (return GraphRepository agar bisa chaining) ---
func (r *graphRepository) WithMatch(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "MATCH "+query)
	return clone
}

func (r *graphRepository) WithMerge(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "MERGE "+query)
	return clone
}

func (r *graphRepository) WithCreate(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "CREATE "+query)
	return clone
}

func (r *graphRepository) WithDelete(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "DELETE "+query)
	return clone
}

func (r *graphRepository) WithWhere(query string, params map[string]interface{}) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "WHERE "+query)
	for k, v := range params {
		clone.params[k] = v
	}
	return clone
}

func (r *graphRepository) WithSet(query string, params map[string]interface{}) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "SET "+query)
	for k, v := range params {
		clone.params[k] = v
	}
	return clone
}

func (r *graphRepository) WithReturn(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "RETURN "+query)
	return clone
}

// --- ✅ Baru: WithParams ---
func (r *graphRepository) WithParams(params map[string]interface{}) adapter.GraphRepository {
	clone := r.clone()
	for k, v := range params {
		clone.params[k] = v
	}
	return clone
}

// --- Build and join ---
func (r *graphRepository) buildQuery() string {
	out := ""
	for _, s := range r.statements {
		out += s + "\n"
	}
	return out
}

func (r *graphRepository) WithOptionalMatch(query string) adapter.GraphRepository {
	clone := r.clone()
	clone.statements = append(clone.statements, "OPTIONAL MATCH "+query)
	return clone
}

// --- Executor ---
func (r *graphRepository) RunRead() ([]neo4j.Record, error) {
	query := r.buildQuery()
	session := r.driver.NewSession(r.ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(r.ctx)

	result, err := session.ExecuteRead(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(r.ctx, query, r.params)
		if err != nil {
			return nil, err
		}
		var records []neo4j.Record
		for res.Next(r.ctx) {
			records = append(records, *res.Record())
		}
		return records, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]neo4j.Record), nil
}

func (r *graphRepository) RunWrite() error {
	query := r.buildQuery()
	session := r.driver.NewSession(r.ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(r.ctx)

	_, err := session.ExecuteWrite(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(r.ctx, query, r.params)
		return nil, err
	})
	return err
}

// --- CRUD Methods ---

func (r *graphRepository) FindNodes(label string) ([]neo4j.Record, error) {
	return r.
		WithMatch("(n:" + label + ")").
		WithReturn("n").
		RunRead()
}

func (r *graphRepository) FindNodeByID(label string, id string) (*neo4j.Record, error) {
	query := fmt.Sprintf("MATCH (n:%s {id: $id}) RETURN n", label)
	session := r.driver.NewSession(r.ctx, r.sessionConfig)
	defer session.Close(r.ctx)

	result, err := session.ExecuteRead(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(r.ctx, query, map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		if res.Next(r.ctx) {
			return res.Record(), nil
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	record := result.(*neo4j.Record)
	return record, nil
}

func (r *graphRepository) CreateNode(label string, data map[string]interface{}) error {
	return r.WithCreate("(n:"+label+")").
		WithSet("n = $data", map[string]interface{}{"data": data}).
		RunWrite()
}

func (r *graphRepository) MergeNode(label string, id string, data map[string]interface{}) error {
	return r.WithMerge("(n:"+label+" {id: $id})").
		WithSet("n += $data", map[string]interface{}{
			"id":   id,
			"data": data,
		}).
		RunWrite()
}

func (r *graphRepository) UpdateNode(label string, id string, updates map[string]interface{}) error {
	return r.WithMatch("(n:"+label+" {id: $id})").
		WithSet("n += $updates", map[string]interface{}{
			"id":      id,
			"updates": updates,
		}).
		RunWrite()
}

func (r *graphRepository) DeleteNode(label string, id string) error {
	return r.WithMatch("(n:" + label + " {id: $id})").
		WithReturn("n").
		RunWrite()
}