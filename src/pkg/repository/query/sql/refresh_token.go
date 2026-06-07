package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"
	"time"

	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db           *gorm.DB
	whereClauses []func(*gorm.DB) *gorm.DB
	order        string
	limit        *int
	unscoped     bool
}

func NewRefreshTokenRepository() adapter.RefreshTokenRepository {
	return &refreshTokenRepository{db: config.DB}
}
// --- 🔁 Chainable Configs ---

func (repo *refreshTokenRepository) clone() *refreshTokenRepository {
	clone := *repo
	return &clone
}

func (repo *refreshTokenRepository) WithTx(tx *gorm.DB) adapter.RefreshTokenRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *refreshTokenRepository) WithWhere(query interface{}, args ...interface{}) adapter.RefreshTokenRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *refreshTokenRepository) WithOrder(order string) adapter.RefreshTokenRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *refreshTokenRepository) WithLimit(limit int) adapter.RefreshTokenRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *refreshTokenRepository) WithUnscoped() adapter.RefreshTokenRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🧱 Builder ---

func (repo *refreshTokenRepository) getQueryBuilder() *builder.QueryBuilder[models.RefreshToken] {
	db := repo.db

	if repo.unscoped {
		db = db.Unscoped()
	}

	qb := builder.NewQueryBuilder[models.RefreshToken](db).
		WithOrder(repo.order)

	for _, where := range repo.whereClauses {
		qb = qb.WithWhere(where)
	}

	if repo.limit != nil {
		qb = qb.WithLimit(*repo.limit)
	}

	return qb
}

// --- 🔧 CRUD ---

func (repo *refreshTokenRepository) Insert(data *models.RefreshToken) (*models.RefreshToken, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *refreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("token = ?", token)
		}).
		FindOne()
}

func (repo *refreshTokenRepository) DeleteByToken(token string) error {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("token = ?", token)
		}).
		DeleteWhere()
}

func (repo *refreshTokenRepository) DeleteByUserID(userID int64) error {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)
		}).
		DeleteWhere()
}

func (repo *refreshTokenRepository) DeleteExpired() error {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("expires_at < ?", time.Now())
		}).
		DeleteWhere()
}
