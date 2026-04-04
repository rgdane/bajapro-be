package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MUserService interface {
	WithTx(tx *gorm.DB) MUserService

	CreateMUser(input *models.MUsers) (*models.MUsers, error)
	UpdateMUser(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MUsers, error)
	DeleteMUser(id int64, isPermanent bool) error
	GetAllMUsers(filter dto.MUserFilterDto) ([]models.MUsers, error)
	GetMUserByID(id int64, filter dto.MUserFilterDto) (*models.MUsers, error)
	GetDB() *gorm.DB
	BulkCreateMUsers(data []*models.MUsers) ([]*models.MUsers, error)
	BulkUpdateMUsers(ids []int64, updates map[string]interface{}, associatons map[string]interface{}) error
	BulkDeleteMUsers(ids []int64, isPermanent bool) error
}

type mUserService struct {
	repo sql.MUsersRepository
}

func NewMUserService(repo sql.MUsersRepository) MUserService {
	return &mUserService{repo: repo}
}

func (s *mUserService) WithTx(tx *gorm.DB) MUserService {
	clone := *s
	clone.repo = s.repo.WithTx(tx)
	return &clone
}

func (s *mUserService) CreateMUser(input *models.MUsers) (*models.MUsers, error) {
	return s.repo.InsertMUser(input)
}

func (s *mUserService) UpdateMUser(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MUsers, error) {
	return s.repo.UpdateMUser(id, updates)
}

func (s *mUserService) DeleteMUser(id int64, isPermanent bool) error {
	return s.repo.RemoveMUser(id)
}

func (s *mUserService) GetAllMUsers(filter dto.MUserFilterDto) ([]models.MUsers, error) {
	repo := s.repo

	// Apply filters
	if filter.RoleID != nil {
		repo = repo.WithWhere("role_id = ?", *filter.RoleID)
	}
	if filter.ClassID != nil {
		repo = repo.WithWhere("class_id = ?", *filter.ClassID)
	}
	if filter.Name != nil {
		repo = repo.WithWhere("name ILIKE ?", "%"+*filter.Name+"%")
	}
	if filter.Email != nil {
		repo = repo.WithWhere("email ILIKE ?", "%"+*filter.Email+"%")
	}
	if filter.IsApprovedByAdmin != nil {
		repo = repo.WithWhere("is_approved_by_admin = ?", *filter.IsApprovedByAdmin)
	}
	if filter.InstansiSekolah != nil {
		repo = repo.WithWhere("instansi_sekolah ILIKE ?", "%"+*filter.InstansiSekolah+"%")
	}
	if filter.IsActive != nil {
		repo = repo.WithWhere("isactive = ?", *filter.IsActive)
	}

	return repo.FindMUser()
}

func (s *mUserService) GetMUserByID(id int64, filter dto.MUserFilterDto) (*models.MUsers, error) {
	return s.repo.FindMUserByID(id)
}

func (s *mUserService) GetDB() *gorm.DB {
	return s.repo.GetDB()
}

func (s *mUserService) BulkCreateMUsers(data []*models.MUsers) ([]*models.MUsers, error) {
	return s.repo.InsertManyMUsers(data)
}

func (s *mUserService) BulkUpdateMUsers(ids []int64, updates map[string]interface{}, associatons map[string]interface{}) error {
	return s.repo.UpdateManyMUsers(ids, updates)
}

func (s *mUserService) BulkDeleteMUsers(ids []int64, isPermanent bool) error {
	return s.repo.RemoveManyMUsers(ids)
}
