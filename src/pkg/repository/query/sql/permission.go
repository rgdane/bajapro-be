package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type permissionRepository struct {
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

func NewPermissionRepository() adapter.PermissionRepository {
	return &permissionRepository{db: config.DB}
}

// --- Chainable Configs ---

func (repo *permissionRepository) clone() *permissionRepository {
	clone := *repo
	return &clone
}

func (repo *permissionRepository) WithTx(tx *gorm.DB) adapter.PermissionRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *permissionRepository) WithPreloads(preloads ...string) adapter.PermissionRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *permissionRepository) WithAssociations(associations ...string) adapter.PermissionRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *permissionRepository) WithReplacements(replacements map[string]interface{}) adapter.PermissionRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *permissionRepository) WithJoins(joins ...string) adapter.PermissionRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *permissionRepository) WithWhere(query interface{}, args ...interface{}) adapter.PermissionRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *permissionRepository) WithOrder(order string) adapter.PermissionRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *permissionRepository) WithLimit(limit int) adapter.PermissionRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *permissionRepository) WithCursor(cursor int) adapter.PermissionRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- Builder Helper ---

func (repo *permissionRepository) getQueryBuilder() *builder.QueryBuilder[models.Permission] {
	qb := builder.NewQueryBuilder[models.Permission](repo.db).
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

// --- CRUD ---

func (repo *permissionRepository) InsertPermission(data *models.Permission) (*models.Permission, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *permissionRepository) UpdatePermission(id int64, updates map[string]interface{}) (*models.Permission, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *permissionRepository) UpdateManyPermissions(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *permissionRepository) RemovePermission(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *permissionRepository) RemoveManyPermissions(ids []int64) error {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("id IN ?", ids)
		}).Delete(nil)
}

func (repo *permissionRepository) FindPermission() ([]models.Permission, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *permissionRepository) FindPermissionByID(id int64) (*models.Permission, error) {
	return repo.getQueryBuilder().FindByID(id)
}
