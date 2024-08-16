package housecreate

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
)

type (
	HouseCreateRequest struct {
		Address   string
		Year      int
		Developer *string
	}

	Repository interface {
		Create(context.Context, HouseCreateRequest) (*models.House, error)
	}
	HouseCreateService struct {
		repo Repository
	}
)

func NewHouseService(repo Repository) *HouseCreateService {
	return &HouseCreateService{repo: repo}
}

func (hs *HouseCreateService) CreateHouse(ctx context.Context, req HouseCreateRequest) (*models.House, error) {
	return hs.repo.Create(ctx, req)
}
