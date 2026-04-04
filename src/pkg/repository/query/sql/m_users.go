package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mUsersRepository struct {
	db           *gorm.DB
	preloads     []string
	associations []string
	replacements map[string]interface{}
	joins        []string
	whereClauses []func(*gorm.DB) *gorm.DB
	order        string
	limit        *int
	cursor       *int
	unscoped     bool
}

func NewMUsersRepository() adapter.MUsersRepository {
	return &mUsersRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mUsersRepository) clone() *mUsersRepository {
	clone := *repo
	return &clone
}

func (repo *mUsersRepository) WithTx(tx *gorm.DB) adapter.MUsersRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mUsersRepository) WithPreloads(preloads ...string) adapter.MUsersRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mUsersRepository) WithAssociations(associations ...string) adapter.MUsersRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mUsersRepository) WithReplacements(replacements map[string]interface{}) adapter.MUsersRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mUsersRepository) WithJoins(joins ...string) adapter.MUsersRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mUsersRepository) WithWhere(query interface{}, args ...interface{}) adapter.MUsersRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mUsersRepository) WithOrder(order string) adapter.MUsersRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mUsersRepository) WithLimit(limit int) adapter.MUsersRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mUsersRepository) WithCursor(cursor int) adapter.MUsersRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *mUsersRepository) WithUnscoped() adapter.MUsersRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🧱 Builder ---

func (repo *mUsersRepository) getQueryBuilder() *builder.QueryBuilder[models.MUsers] {
	db := repo.db.Table("m_users")
	if repo.unscoped {
		db = db.Unscoped()
	}

	qb := builder.NewQueryBuilder[models.MUsers](db).
		WithPreloads(repo.preloads...).
		WithAssociations(repo.associations...).
		WithReplacements(repo.replacements).
		WithJoins(repo.joins...).
		WithOrder(repo.order)

	for _, where := range repo.whereClauses {
		qb = qb.WithWhere(where)
	}
	if repo.limit != nil {
		qb = qb.WithLimit(*repo.limit)
	}
	if repo.cursor != nil {
		qb = qb.WithCursor(*repo.cursor)
	}
	return qb
}

// --- 🔧 CRUD ---

func (repo *mUsersRepository) InsertMUser(data *models.MUsers) (*models.MUsers, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mUsersRepository) InsertManyMUsers(data []*models.MUsers) ([]*models.MUsers, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mUsersRepository) UpdateMUser(id int64, updates map[string]interface{}) (*models.MUsers, error) {
	data, err := repo.getQueryBuilder().UpdateByID(id, updates)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mUsersRepository) UpdateManyMUsers(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mUsersRepository) RemoveMUser(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mUsersRepository) RemoveManyMUsers(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

// --- 🔍 Finders ---

func (repo *mUsersRepository) FindMUser() ([]models.MUsers, error) {
	data, err := repo.getQueryBuilder().FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mUsersRepository) FindMUserByID(id int64) (*models.MUsers, error) {
	data, err := repo.getQueryBuilder().FindByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mUsersRepository) FindMUserByEmail(email string) (*models.MUsers, error) {
	// Explicitly set the table and build the query from scratch to ensure correct table targeting
	db := repo.db.Table("m_users")
	
	for _, preload := range repo.preloads {
		db = db.Preload(preload)
	}
	
	var user models.MUsers
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *mUsersRepository) GetDB() *gorm.DB {
	return repo.db
}
