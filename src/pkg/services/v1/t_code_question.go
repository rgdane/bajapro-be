package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TCodeQuestionService interface {
	WithTx(tx *gorm.DB) TCodeQuestionService
	GetCodeQuestionsBySubLessonID(subLessonID int64) ([]models.CodeQuestion, error)
	GetCodeQuestionByID(id int64) (*models.CodeQuestion, error)
	CreateCodeQuestion(data *models.CodeQuestion) (*models.CodeQuestion, error)
	RunCodeTest(input *dto.CodeTestRunDto) (*dto.CodeTestRunResponseDto, error)
	SubmitCodeTest(
		userID int64,
		input *dto.SubmitCodeTestDto,
	) (*dto.SubmitCodeTestResponseDto, error)
	GetDB() *gorm.DB
}

type tCodeQuestionService struct {
	repo                     sql.TCodeQuestionRepository
	userRepository           sql.UserRepository
	tCodeAnswerRepository    sql.TCodeAnswerRepository
	essayQuestionRepository  sql.EssayQuestionRepository
	tEssayAnswerRepository   sql.TEssayAnswerRepository
	tStudentCourseRepository sql.TStudentCourseRepository
	mBadgeSettingsRepository sql.MBadgeSettingsRepository
	tStudentProgressService  TStudentProgressService

	tx *gorm.DB
}

func NewTCodeQuestionService(
	repo sql.TCodeQuestionRepository,
	userRepository sql.UserRepository,
	tCodeAnswerRepository sql.TCodeAnswerRepository,
	essayQuestionRepository sql.EssayQuestionRepository,
	tEssayAnswerRepository sql.TEssayAnswerRepository,
	tStudentCourseRepository sql.TStudentCourseRepository,
	mBadgeSettingsRepository sql.MBadgeSettingsRepository,
	tStudentProgressService TStudentProgressService,
) TCodeQuestionService {
	return &tCodeQuestionService{
		repo:                     repo,
		userRepository:           userRepository,
		tCodeAnswerRepository:    tCodeAnswerRepository,
		essayQuestionRepository:  essayQuestionRepository,
		tEssayAnswerRepository:   tEssayAnswerRepository,
		tStudentCourseRepository: tStudentCourseRepository,
		mBadgeSettingsRepository: mBadgeSettingsRepository,
		tStudentProgressService:  tStudentProgressService,
	}
}

func (s *tCodeQuestionService) WithTx(tx *gorm.DB) TCodeQuestionService {
	return &tCodeQuestionService{
		repo:                     s.repo.WithTx(tx),
		userRepository:           s.userRepository.WithTx(tx),
		tCodeAnswerRepository:    s.tCodeAnswerRepository.WithTx(tx),
		essayQuestionRepository:  s.essayQuestionRepository.WithTx(tx),
		tEssayAnswerRepository:   s.tEssayAnswerRepository.WithTx(tx),
		tStudentCourseRepository: s.tStudentCourseRepository.WithTx(tx),
		mBadgeSettingsRepository: s.mBadgeSettingsRepository.WithTx(tx),
		tStudentProgressService:  s.tStudentProgressService.WithTx(tx),
		tx:                       tx,
	}
}

func (s *tCodeQuestionService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tCodeQuestionService) GetCodeQuestionsBySubLessonID(subLessonID int64) ([]models.CodeQuestion, error) {
	data, err := s.repo.FindTCodeQuestionsBySubLessonID(subLessonID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeQuestionService) CreateCodeQuestion(data *models.CodeQuestion) (*models.CodeQuestion, error) {
	data, err := s.repo.CreateTCodeQuestion(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeQuestionService) GetCodeQuestionByID(id int64) (*models.CodeQuestion, error) {
	data, err := s.repo.FindTCodeQuestionByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeQuestionService) RunCodeTest(
	input *dto.CodeTestRunDto,
) (*dto.CodeTestRunResponseDto, error) {

	// Cari user berdasarkan email
	user, err := s.userRepository.FindUserByEmail(input.User)
	if err != nil {
		return nil, err
	}

	// // Mapping request Go -> Python
	// pythonRequest := dto.PythonRunCodeRequestDto{
	// 	Code:       input.Code,
	// 	QuestionID: input.CodeQuestionID,
	// 	User:       input.User,
	// }

	// Marshal request
	form := url.Values{}
	form.Set("user", input.User)
	form.Set("code", input.Code)
	form.Set(
		"question_id",
		strconv.FormatInt(input.CodeQuestionID, 10),
	)

	resp, err := http.PostForm(
		config.AppConfig.CompilerURL+"/api/run",
		form,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Cek status response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("compiler service returned status %d", resp.StatusCode)
	}

	// Decode response Python
	var result dto.CodeTestRunResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Cari jawaban user untuk soal ini
	answer, err := s.tCodeAnswerRepository.FindByUserIDAndCodeQuestionID(
		user.ID,
		input.CodeQuestionID,
	)

	if errors.Is(err, gorm.ErrRecordNotFound) {

		// INSERT
		_, err = s.tCodeAnswerRepository.CreateTCodeAnswer(&models.TCodeAnswer{
			UserID:         user.ID,
			CodeQuestionID: input.CodeQuestionID,
			IsCodeRight:    result.Result.Passed,
			ExploringScore: result.Result.Point,
		})
		if err != nil {
			return nil, err
		}

	} else if err != nil {

		return nil, err

	} else {

		// UPDATE
		answer.IsCodeRight = result.Result.Passed
		answer.ExploringScore = result.Result.Point

		_, err = s.tCodeAnswerRepository.UpdateTCodeAnswer(answer)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (s *tCodeQuestionService) SubmitCodeTest(
	userID int64,
	input *dto.SubmitCodeTestDto,
) (*dto.SubmitCodeTestResponseDto, error) {

	//----------------------------------------------------
	// 1. Ambil Code Answer
	//----------------------------------------------------

	codeAnswer, err := s.tCodeAnswerRepository.FindByUserIDAndCodeQuestionID(
		userID,
		input.CodeQuestionID,
	)

	if err != nil {
		return nil, err
	}

	//----------------------------------------------------
	// 2. Update Exploring Score
	//----------------------------------------------------

	codeAnswer.ExploringScore = input.Score
	codeAnswer.IsCodeRight = input.Score >= 100

	if _, err := s.tCodeAnswerRepository.UpdateTCodeAnswer(codeAnswer); err != nil {
		return nil, err
	}

	//----------------------------------------------------
	// 3. Nilai Essay
	//----------------------------------------------------

	totalEssayScore := 0

	for i, essayID := range input.EssayQuestionID {

		essayQuestion, err := s.essayQuestionRepository.FindEssayQuestionByID(essayID)
		if err != nil {
			return nil, err
		}

		//------------------------------------------
		// Request ke Django
		//------------------------------------------

		form := url.Values{}

		form.Add("esay_answer", essayQuestion.Answer)
		form.Add("esay_answer2", essayQuestion.Answer2)
		form.Add("esay_answer3", essayQuestion.Answer3)
		form.Add("esay_answer4", essayQuestion.Answer4)
		form.Add("user_answer", input.Answer[i])

		resp, err := http.PostForm(
			config.AppConfig.CompilerURL+"/compiler/generate/grade",
			form,
		)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		var grade struct {
			Output int `json:"output"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&grade); err != nil {
			return nil, err
		}

		totalEssayScore += grade.Output

		//------------------------------------------
		// Simpan Jawaban Essay
		//------------------------------------------

		answer, err := s.tEssayAnswerRepository.FindTEssayAnswersByEssayQuestionIDAndUserID(
			userID,
			essayID,
		)

		if errors.Is(err, gorm.ErrRecordNotFound) {

			model := &models.TEssayAnswer{
				UserID:          userID,
				EssayQuestionID: essayID,
				Answer:          input.Answer[i],
			}

			switch i {

			case 0:
				model.KonteksPenjelasan = grade.Output

			case 1:
				model.Keruntutan = grade.Output

			case 2:
				model.Kebenaran = grade.Output

			}

			if _, err := s.tEssayAnswerRepository.CreateTEssayAnswer(model); err != nil {
				return nil, err
			}

		} else {

			answer.Answer = input.Answer[i]

			switch i {

			case 0:
				answer.KonteksPenjelasan = grade.Output

			case 1:
				answer.Keruntutan = grade.Output

			case 2:
				answer.Kebenaran = grade.Output

			}

			if _, err := s.tEssayAnswerRepository.UpdateTEssayAnswerByID(answer); err != nil {
				return nil, err
			}

		}

	}

	//----------------------------------------------------
	// 4. Update Total Score Student Course
	//----------------------------------------------------

	studentCourse, err := s.tStudentCourseRepository.FindByUserIDAndCourseID(
		userID,
		input.CourseID,
	)

	if err != nil {
		return nil, err
	}

	studentCourse.TotalScore =
		codeAnswer.ExploringScore +
			totalEssayScore

	badge, err := s.mBadgeSettingsRepository.FindMBadgeByScore(studentCourse.TotalScore)
	if err != nil {
		return nil, err
	}
	fmt.Println("Total Score:", studentCourse.TotalScore)
	fmt.Println("Badge ID:", badge.ID)
	fmt.Println("Badge Name:", badge.Name)

	studentCourse.BadgeID = badge.ID

	if _, err := s.tStudentCourseRepository.UpdateTotalScoreByID(studentCourse); err != nil {
		return nil, err
	}
	fmt.Println("Saved Badge ID:", studentCourse.BadgeID)
	//----------------------------------------------------
	// 5. Update Student Progress
	//----------------------------------------------------

	progress := &models.TStudentProgress{
		UserID:      userID,
		SubLessonID: input.SubLessonID,
		Status:      "completed",
	}

	_, err = s.tStudentProgressService.CompleteTStudentProgress(progress)
	if err != nil {
		return nil, err
	}

	//----------------------------------------------------

	return &dto.SubmitCodeTestResponseDto{
		Success: true,
		Message: "Submit berhasil",
	}, nil

}
