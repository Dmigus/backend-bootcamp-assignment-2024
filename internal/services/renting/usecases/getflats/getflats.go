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
	GetFlatsService struct {
		repo Repository
	}
)

func NewGetFlatsService(repo Repository) *GetFlatsService {
	return &GetFlatsService{repo: repo}
}

func (s *GetFlatsService) GetFlats(ctx context.Context, id int, role models.UserRole) ([]models.Flat, error) {
	switch role {
	case models.Moderator:
		return s.repo.GetFlats(ctx, id)
	case models.Client:
		return s.repo.GetApprovedFlats(ctx, id)
	default:
		return nil, models.ErrUnknownRole
	}
}
