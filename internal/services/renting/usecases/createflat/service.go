package createflat

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
)

type (
	Request struct {
		HouseId int
		Price   int
		Rooms   int
	}
	FlatsRepo interface {
		CreateFlat(ctx context.Context, flat Request) (*models.Flat, error)
	}
	HousesRepo interface {
		HouseUpdated(ctx context.Context, houseId int) error
	}
	TxManager interface {
		WithinTransaction(ctx context.Context, f func(context.Context) bool) error
	}
	Service struct {
		flatsRepo  FlatsRepo
		housesRepo HousesRepo
		txManager  TxManager
	}
)

func NewService(flatsRepo FlatsRepo, housesRepo HousesRepo, txManager TxManager) *Service {
	return &Service{flatsRepo: flatsRepo, housesRepo: housesRepo, txManager: txManager}
}

func (f *Service) CreateFlat(ctx context.Context, flat Request) (*models.Flat, error) {
	var createdFlat *models.Flat
	var serviceErr error
	txErr := f.txManager.WithinTransaction(ctx, func(txCtx context.Context) bool {
		createdFlat, serviceErr = f.createFlatAndUpdateHouse(txCtx, flat)
		if serviceErr != nil {
			return false
		}
		return true
	})
	if serviceErr != nil {
		return nil, serviceErr
	}
	if txErr != nil {
		return nil, txErr
	}
	return createdFlat, nil
}

func (f *Service) createFlatAndUpdateHouse(ctx context.Context, flat Request) (*models.Flat, error) {
	createdFlat, err := f.flatsRepo.CreateFlat(ctx, flat)
	if err != nil {
		return nil, err
	}
	err = f.housesRepo.HouseUpdated(ctx, flat.HouseId)
	if err != nil {
		return nil, err
	}
	return createdFlat, nil
}
