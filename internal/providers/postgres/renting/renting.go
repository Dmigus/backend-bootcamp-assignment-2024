package renting

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	"context"
	"github.com/samber/lo"
)

type Renting struct {
	defaultDBTX DBTX
}

func NewRenting(tx DBTX) *Renting {
	return &Renting{defaultDBTX: tx}
}

func (h *Renting) Create(ctx context.Context, req housecreate.HouseCreateRequest) (*models.House, error) {
	params := houseServiceToDto(req)
	queries := New(h.getDBTX(ctx))
	newHouse, err := queries.createHouse(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := houseDtoToService(newHouse)
	return &resp, nil
}

func (h *Renting) GetFlats(ctx context.Context, houseId int) ([]models.Flat, error) {
	queries := New(h.getDBTX(ctx))
	flats, err := queries.getFlats(ctx, int64(houseId))
	if err != nil {
		return nil, err
	}
	flatsModel := lo.Map(flats, func(flat Flat, _ int) models.Flat {
		return flatDtoToService(flat)
	})
	return flatsModel, nil
}

func (h *Renting) GetApprovedFlats(ctx context.Context, houseId int) ([]models.Flat, error) {
	queries := New(h.getDBTX(ctx))
	flats, err := queries.getApprovedFlats(ctx, int64(houseId))
	if err != nil {
		return nil, err
	}
	flatsModel := lo.Map(flats, func(flat Flat, _ int) models.Flat {
		return flatDtoToService(flat)
	})
	return flatsModel, nil
}

func (h *Renting) getDBTX(ctx context.Context) DBTX {
	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		return h.defaultDBTX
	}
	dbtx, ok := tx.(DBTX)
	if ok {
		return dbtx
	}
	return h.defaultDBTX
}
