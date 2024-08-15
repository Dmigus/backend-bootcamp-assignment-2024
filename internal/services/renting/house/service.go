package house

import (
	"context"
	"errors"
	"time"
)

var ErrHouseNotFound = errors.New("house not found")

type (
	HouseCreateRequest struct {
		Address   string
		Year      int
		Developer string
	}
	HouseCreateResponse struct {
		Id        int
		Address   string
		Year      int
		Developer string
		CreatedAt time.Time
		UpdateAt  time.Time
	}
	houseRepository interface {
		Create(context.Context, HouseCreateRequest) (HouseCreateResponse, error)
		Get(context.Context, int) (HouseCreateResponse, error)
	}
	HouseService struct {
		repo houseRepository
	}
)

func (hs *HouseService) CreateHouse(ctx context.Context, req HouseCreateRequest) (HouseCreateResponse, error) {
	return hs.repo.Create(ctx, req)
}

func (hs *HouseService) GetHouseByID(ctx context.Context, id int) (HouseCreateResponse, error) {
	return hs.repo.Get(ctx, id)
}
