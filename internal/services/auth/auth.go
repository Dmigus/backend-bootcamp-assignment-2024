package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type UserType int

const (
	Client UserType = iota + 1
	Moderator
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

func (a *AuthService) DummyLogin(ut UserType) (string, error) {
	claims := jwt.MapClaims{}
	switch ut {
	case Client:
		claims["role"] = clientRoleName
	case Moderator:
		claims["role"] = moderatorRoleName
	default:
		return "", ErrNoSuchUserType
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(a.key)
}

// GetRole проверяет токен и возвращает роль пользователя, если токен валидный. Если токен невалидный, то ErrInvalidToken
func (a *AuthService) GetRole(tokenStr string) (UserType, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return a.key, nil
	})
	if err != nil {
		return UserType(0), ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role := claims["role"]
		switch role {
		case clientRoleName:
			return Client, nil
		case moderatorRoleName:
			return Moderator, nil
		default:
			return UserType(0), ErrInvalidToken
		}
	}
	return UserType(0), ErrInvalidToken
}
