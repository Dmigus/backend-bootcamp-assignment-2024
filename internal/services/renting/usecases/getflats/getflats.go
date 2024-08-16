package getflats

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
)

type (
	Repository interface {
		GetFlats(context.Context, int) ([]models.Flat, error)
	}
	GetFlatsService struct {
		repo Repository
	}
)

func NewGetFlatsService(repo Repository) *GetFlatsService {
	return &GetFlatsService{repo: repo}
}

func (s *GetFlatsService) GetFlats(ctx context.Context, id int) ([]models.Flat, error) {
	return s.repo.GetFlats(ctx, id)
}
