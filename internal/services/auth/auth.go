package auth

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const (
	clientRoleName    = "client"
	moderatorRoleName = "moderator"
)

var (
	ErrNoSuchUserType = errors.New("userType invalid")
	ErrInvalidToken   = errors.New("invalid token")
)

type AuthService struct {
	key []byte
}

func NewAuthService(key []byte) *AuthService {
	return &AuthService{key: key}
}

func (a *AuthService) DummyLogin(ut models.UserRole) (string, error) {
	claims := jwt.MapClaims{}
	switch ut {
	case models.Client:
		claims["role"] = clientRoleName
	case models.Moderator:
		claims["role"] = moderatorRoleName
	default:
		return "", ErrNoSuchUserType
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(a.key)
}

// GetRole проверяет токен и возвращает роль пользователя, если токен валидный. Если токен невалидный, то ErrInvalidToken
func (a *AuthService) GetRole(tokenStr string) (models.UserRole, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return a.key, nil
	})
	if err != nil {
		return models.UserRole(0), ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role := claims["role"]
		switch role {
		case clientRoleName:
			return models.Client, nil
		case moderatorRoleName:
			return models.Moderator, nil
		default:
			return models.UserRole(0), ErrInvalidToken
		}
	}
	return models.UserRole(0), ErrInvalidToken
}
