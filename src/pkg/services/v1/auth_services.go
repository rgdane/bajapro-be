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
	Login(email, password string) (*models.User, error)
	GetProfile(token string) (*models.User, error)
	GenerateToken(userID int64, name string) (string, error)
	DecodeToken(token string) (jwt.MapClaims, error)
}

type authService struct {
	repo sql.UserRepository
}

func NewAuthService(userRepo sql.UserRepository) *authService {
	return &authService{repo: userRepo}
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.
		WithPreloads("HasRoles.HasPermissions").
		FindUserByEmail(email)

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	log.Printf("User ID: %d, Roles count: %d", user.ID, len(user.HasRoles))
	for i, role := range user.HasRoles {
		log.Printf("Role %d: %+v", i, role)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, bcrypt_err.TranslateBcryptError(err)
	}

	return user, nil
}

func (s *authService) GetProfile(token string) (user *models.User, err error) {
	claims, err := s.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	userID, ok := claims["user_id"].(float64)

	if !ok {
		return nil, err
	}
	user, err = s.repo.WithPreloads("HasRoles.HasPermissions", "HasSquads.HasProject", "HasTitle.HasPosition").FindUserByID(int64(userID))
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
