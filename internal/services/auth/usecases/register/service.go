package register

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type (
	Request struct {
		Id             models.UserId
		Email          string
		Salt           string
		HashedPassword string
		Role           models.UserRole
	}
	Repository interface {
		Add(ctx context.Context, req Request) error
	}
	SaltGenerator interface {
		NewSalt() string
	}
	UserIdGenerator interface {
		NewUserId() models.UserId
	}
	PasswordHasher interface {
		Hash(password string, salt string) (string, error)
	}
	Service struct {
		saltGenerator   SaltGenerator
		repo            Repository
		userIdGenerator UserIdGenerator
		passwordHasher  PasswordHasher
	}
)

func (s *Service) Register(ctx context.Context, email, password string, role models.UserRole) (*models.UserId, error) {
	salt := s.saltGenerator.NewSalt()
	hashedPassword, err := s.passwordHasher.Hash(password, salt)
	if err != nil {
		return nil, err
	}
	uuid := s.userIdGenerator.NewUserId()
	req := Request{
		Id:             s.userIdGenerator.NewUserId(),
		Email:          email,
		Salt:           salt,
		HashedPassword: hashedPassword,
		Role:           role,
	}
	err = s.repo.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return &uuid, nil
}
