package graphdb

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type GraphRepository interface {
	// --- CRUD operations ---
	FindNodes(label string) ([]neo4j.Record, error)
	FindNodeByID(label string, id string) (*neo4j.Record, error)
	CreateNode(label string, data map[string]interface{}) error
	UpdateNode(label string, id string, updates map[string]interface{}) error
	DeleteNode(label string, id string) error
	MergeNode(label string, id string, data map[string]interface{}) error

	// --- Fluent builder ---
	WithMatch(match string) GraphRepository
	WithMerge(merge string) GraphRepository
	WithCreate(create string) GraphRepository
	WithDelete(query string) GraphRepository
	WithWhere(where string, params map[string]interface{}) GraphRepository
	WithSet(set string, params map[string]interface{}) GraphRepository
	WithReturn(returns string) GraphRepository
	WithParams(params map[string]interface{}) GraphRepository
	WithOptionalMatch(query string) GraphRepository

	// --- Execution ---
	RunRead() ([]neo4j.Record, error)
	RunWrite() error
}
