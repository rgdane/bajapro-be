package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type roleRepository struct {
	db           *gorm.DB
	preloads     []string
	associations []string
	replacements map[string]interface{}
	joins        []string
	whereClauses []func(*gorm.DB) *gorm.DB
	order        string
	limit        *int
	cursor       *int
}

func NewRoleRepository() adapter.RoleRepository {
	return &roleRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *roleRepository) clone() *roleRepository {
	clone := *repo
	return &clone
}

func (repo *roleRepository) WithTx(tx *gorm.DB) adapter.RoleRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *roleRepository) WithPreloads(preloads ...string) adapter.RoleRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *roleRepository) WithAssociations(associations ...string) adapter.RoleRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *roleRepository) WithReplacements(replacements map[string]interface{}) adapter.RoleRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *roleRepository) WithJoins(joins ...string) adapter.RoleRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *roleRepository) WithWhere(query interface{}, args ...interface{}) adapter.RoleRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *roleRepository) WithOrder(order string) adapter.RoleRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *roleRepository) WithLimit(limit int) adapter.RoleRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *roleRepository) WithCursor(cursor int) adapter.RoleRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *roleRepository) getQueryBuilder() *builder.QueryBuilder[models.Role] {
	qb := builder.NewQueryBuilder[models.Role](repo.db).
		WithPreloads(repo.preloads...).
		WithAssociations(repo.associations...).
		WithReplacements(repo.replacements).
		WithJoins(repo.joins...).
		WithOrder(repo.order)

	for _, w := range repo.whereClauses {
		qb = qb.WithWhere(w)
	}
	if repo.limit != nil {
		qb = qb.WithLimit(*repo.limit)
	}
	if repo.cursor != nil {
		qb = qb.WithCursor(*repo.cursor)
	}
	return qb
}

// --- 🔧 CRUD Methods ---

func (repo *roleRepository) InsertRole(data *models.Role) (*models.Role, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *roleRepository) UpdateRole(id int64, updates map[string]interface{}) (*models.Role, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *roleRepository) UpdateManyRoles(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *roleRepository) RemoveRole(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *roleRepository) RemoveManyRoles(ids []int64) error {
	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).Delete(nil)
}

func (repo *roleRepository) FindRole() ([]models.Role, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *roleRepository) FindRoleByID(id int64) (*models.Role, error) {
	return repo.getQueryBuilder().FindByID(id)
}
