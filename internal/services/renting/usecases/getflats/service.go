package getflats

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
)

type (
	Repository interface {
		GetFlats(context.Context, int) ([]models.Flat, error)
		GetApprovedFlats(context.Context, int) ([]models.Flat, error)
	}
	Service struct {
		repo Repository
	}
)

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetFlats(ctx context.Context, id int, role models.UserRole) ([]models.Flat, error) {
	switch role {
	case models.Moderator:
		return s.repo.GetFlats(ctx, id)
	case models.Client:
		return s.repo.GetApprovedFlats(ctx, id)
	default:
		return nil, models.ErrUnknownRole
	}
}
