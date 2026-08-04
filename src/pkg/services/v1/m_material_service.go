package services

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type MMaterialService interface {
	WithTx(tx *gorm.DB) MMaterialService

	GetCourseMaterialTree(
		courseID int64,
		levelID int64,
	) ([]dto.CourseMaterialTreeResponseDto, error)
	CreateMMaterial(input *models.MMaterials) (*models.MMaterials, error)
	UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error)
	DeleteMMaterial(id int64) error
	GetAllMMaterials(filter dto.MMaterialFilterDto) ([]models.MMaterials, error)
	GetMMaterialByID(id int64, filter dto.MMaterialFilterDto) (*models.MMaterials, error)
	GetMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error)
	GetDB() *gorm.DB
	BulkCreateMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error)
	BulkUpdateMMaterials(ids []int64, updates map[string]interface{}) error
	BulkDeleteMMaterials(ids []int64) error
}

type mMaterialService struct {
	repo              sql.MMaterialRepository
	lessonRepo        sql.MLessonRepository
	essayQuestionRepo sql.EssayQuestionRepository
	tx                *gorm.DB
}

func NewMMaterialService(
	repo sql.MMaterialRepository,
	lessonRepo sql.MLessonRepository,
	essayQuestionRepo sql.EssayQuestionRepository,
) MMaterialService {
	return &mMaterialService{
		repo:              repo,
		lessonRepo:        lessonRepo,
		essayQuestionRepo: essayQuestionRepo,
	}
}

func (s *mMaterialService) WithTx(tx *gorm.DB) MMaterialService {
	return &mMaterialService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mMaterialService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mMaterialService) GetCourseMaterialTree(
	courseID int64,
	levelID int64,
) ([]dto.CourseMaterialTreeResponseDto, error) {
	lessons, err := s.lessonRepo.
		FindCourseMaterialTree(
			courseID,
			levelID,
		)
	if err != nil {
		return nil,
			gorm_err.TranslateGormError(err)
	}

	essayQuestions, err := s.essayQuestionRepo.
		FindActiveEssayQuestions()
	if err != nil {
		return nil,
			gorm_err.TranslateGormError(err)
	}

	essayBySubLessonID := make(
		map[int64][]models.EssayQuestion,
	)

	essayByCodeQuestionID := make(
		map[int64][]models.EssayQuestion,
	)

	for _, question := range essayQuestions {
		if question.SubLessonID != 0 {
			essayBySubLessonID[
				question.SubLessonID] = append(
				essayBySubLessonID[
					question.SubLessonID],
				question,
			)

			continue
		}

		if question.CodeQuestionID != 0 {
			essayByCodeQuestionID[
				question.CodeQuestionID] = append(
				essayByCodeQuestionID[
					question.CodeQuestionID],
				question,
			)
		}
	}

	response := make(
		[]dto.CourseMaterialTreeResponseDto,
		0,
		len(lessons),
	)

	for _, lesson := range lessons {
		subLessons := make(
			[]dto.CourseSubLessonResponseDto,
			0,
			len(lesson.SubLessons),
		)

		for _, subLesson := range lesson.SubLessons {
			materials := make(
				[]dto.CourseMaterialResponseDto,
				0,
				len(subLesson.Materials),
			)

			for _, material := range subLesson.Materials {
				materials = append(
					materials,
					dto.CourseMaterialResponseDto{
						ID: material.ID,
						SubLessonID: material.SubLessonID,
						Title: material.Title,
						Materials: material.Materials,
						URLVideo: material.URLVideo,
						ContentPosition: material.ContentPosition,
						PromptLim: material.PromptLLM,
						Published: material.Published,
						IsActive: material.IsActive,
						CreatedAt: material.CreatedAt,
						UpdatedAt: material.UpdatedAt,
					},
				)
			}

			var codeQuestionResponse *dto.CourseCodeQuestionResponseDto
			var essayQuestionModels []models.EssayQuestion

			if len(subLesson.CodeQuestions) > 0 {
				codeQuestion := subLesson.CodeQuestions[0]

				codeQuestionResponse =
					&dto.CourseCodeQuestionResponseDto{
						ID: codeQuestion.ID,
						SubLessonID: codeQuestion.SubLessonID,
						CodeQuestion: codeQuestion.CodeQuestion,
						Image: codeQuestion.Image,
						Score: codeQuestion.Score,
						Hint: codeQuestion.Hint,
						IsActive: codeQuestion.IsActive,
						CreatedAt: codeQuestion.CreatedAt,
						UpdatedAt: codeQuestion.UpdatedAt,
					}

				essayQuestionModels = append(
					essayQuestionModels,
					essayByCodeQuestionID[
						codeQuestion.ID]...,
				)
			}

			// Essay question yang memiliki sub_lesson_id
			// lebih diprioritaskan, sesuai logika Next.js.
			if questions, exists :=
				essayBySubLessonID[subLesson.ID]; exists {
				essayQuestionModels = questions
			}

			essayResponses := make(
				[]dto.CourseEssayQuestionResponseDto,
				0,
				len(essayQuestionModels),
			)

			for _, essay := range essayQuestionModels {
				essayResponses = append(
					essayResponses,
					dto.CourseEssayQuestionResponseDto{
						ID: essay.ID,
						SubLessonID: essay.SubLessonID,
						CodeQuestionID: essay.CodeQuestionID,
						EssayQuestion: essay.EssayQuestion,
						Answer: essay.Answer,
						Answer2: essay.Answer2,
						Answer3: essay.Answer3,
						Answer4: essay.Answer4,
						IsActive: essay.IsActive,
						CreatedAt: essay.CreatedAt,
						UpdatedAt: essay.UpdatedAt,
					},
				)
			}

			subLessons = append(
				subLessons,
				dto.CourseSubLessonResponseDto{
					ID: subLesson.ID,
					LessonID: subLesson.LessonID,
					Title: subLesson.Title,
					OrderPosition: subLesson.OrderPosition,
					IsActive: subLesson.IsActive,
					CreatedAt: subLesson.CreatedAt,
					UpdatedAt: subLesson.UpdatedAt,
					Materials: materials,
					CodeQuestion: codeQuestionResponse,
					EssayQuestions: essayResponses,
				},
			)
		}

		response = append(
			response,
			dto.CourseMaterialTreeResponseDto{
				ID: lesson.ID,
				CourseID: lesson.CourseID,
				LevelID: lesson.LevelID,
				Title: lesson.Title,
				Description: lesson.Description,
				Position: lesson.Position,
				ImgThumbnail: lesson.ImgThumbnail,
				Published: lesson.Published,
				IsActive: lesson.IsActive,
				CreatedAt: lesson.CreatedAt,
				UpdatedAt: lesson.UpdatedAt,
				SubLessons: subLessons,
			},
		)
	}

	return response, nil
}

func (s *mMaterialService) CreateMMaterial(input *models.MMaterials) (*models.MMaterials, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertMMaterial(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error) {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = repo.WithUnscoped()
	}

	if _, err := repo.FindMMaterialByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateMMaterial(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) DeleteMMaterial(id int64) error {
	err := s.repo.RemoveMMaterial(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mMaterialService) GetAllMMaterials(filter dto.MMaterialFilterDto) ([]models.MMaterials, error) {
	repo := s.repo

	if filter.Preload {
		repo = repo.WithPreloads("SubLesson")
	}

	if filter.Name != "" {
		repo = repo.WithWhere("title ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.Sort != "" && filter.Order != "" {
		orderClause := filter.Sort + " " + filter.Order
		repo = repo.WithOrder(orderClause)
	}

	if filter.ShowDeleted {
		repo = repo.WithUnscoped().WithWhere("m_materials.deleted_at IS NOT NULL")
	}
	fmt.Println("filter", filter)

	data, err := repo.FindMMaterials()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) GetMMaterialByID(id int64, filter dto.MMaterialFilterDto) (*models.MMaterials, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("SubLesson")
	}
	data, err := repo.FindMMaterialByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) BulkCreateMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error) {
	datas, err := s.repo.InsertManyMMaterials(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return datas, nil
}

func (s *mMaterialService) BulkUpdateMMaterials(ids []int64, updates map[string]interface{}) error {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = s.repo.WithUnscoped()
	}

	err := repo.UpdateManyMMaterials(ids, updates)
	return gorm_err.TranslateGormError(err)
}

func (s *mMaterialService) GetMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error) {
	data, err := s.repo.FindMMaterialsByIDs(ids)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) BulkDeleteMMaterials(ids []int64) error {
	err := s.repo.RemoveManyMMaterials(ids)
	return gorm_err.TranslateGormError(err)
}
