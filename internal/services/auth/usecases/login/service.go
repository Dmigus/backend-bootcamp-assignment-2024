package login

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"errors"
)

var (
	ErrNoSuchUserType     = errors.New("userType invalid")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type (
	JWTCodec interface {
		Encode(claims models.AuthClaims) (string, error)
		Decode(token string) (*models.AuthClaims, error)
	}
	AuthData struct {
		Salt []byte
		Hash []byte
		Role models.UserRole
	}
	Repository interface {
		GetAuthData(ctx context.Context, userId models.UserId) (*AuthData, error)
	}
	PasswordHasher interface {
		CheckPasswordHash(salt []byte, password string, hash []byte) bool
	}
	Service struct {
		codec          JWTCodec
		repo           Repository
		passwordHasher PasswordHasher
	}
)

func NewService(codec JWTCodec, repo Repository, passwordHasher PasswordHasher) *Service {
	return &Service{codec: codec, repo: repo, passwordHasher: passwordHasher}
}

func (a *Service) Login(ctx context.Context, userId models.UserId, password string) (string, error) {
	authData, err := a.repo.GetAuthData(ctx, userId)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	if a.passwordHasher.CheckPasswordHash(authData.Salt, password, authData.Hash) {
		claims := models.AuthClaims{Role: authData.Role}
		return a.codec.Encode(claims)
	}
	return "", ErrInvalidCredentials
}

func (a *Service) DummyLogin(ut models.UserRole) (string, error) {
	claims := models.AuthClaims{Role: ut}
	return a.codec.Encode(claims)
}

// GetRole проверяет токен и возвращает роль пользователя, если токен валидный. Если токен невалидный, то ErrInvalidToken
func (a *Service) GetRole(tokenStr string) (models.UserRole, error) {
	claims, err := a.codec.Decode(tokenStr)
	if err != nil {
		return models.UserRole(0), err
	}
	return claims.Role, nil
}
