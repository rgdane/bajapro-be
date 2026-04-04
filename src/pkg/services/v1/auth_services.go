package services

import (
	"log"
	"os"

	"jk-api/internal/database/models"
	"jk-api/internal/errors/bcrypt_err"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (*models.MUsers, error)
	GetProfile(token string) (*models.MUsers, error)
	GenerateToken(userID int64, name string) (string, error)
	DecodeToken(token string) (jwt.MapClaims, error)
}

type authService struct {
	repo sql.MUsersRepository
}

func NewAuthService(userRepo sql.MUsersRepository) *authService {
	return &authService{repo: userRepo}
}

func (s *authService) Login(email, password string) (*models.MUsers, error) {
	user, err := s.repo.
		WithPreloads("Role", "Class").
		FindMUserByEmail(email)

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	log.Printf("User ID: %d, Role ID: %d", user.ID, user.RoleID)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, bcrypt_err.TranslateBcryptError(err)
	}

	return user, nil
}

func (s *authService) GetProfile(token string) (user *models.MUsers, err error) {
	claims, err := s.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	userID, ok := claims["user_id"].(float64)

	if !ok {
		return nil, err
	}
	user, err = s.repo.WithPreloads("Role", "Class").FindMUserByID(int64(userID))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) GenerateToken(userID int64, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"name":    name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) DecodeToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return claims, err
}
