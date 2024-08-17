package houses

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Houses struct {
	defaultDBTX DBTX
}

func NewHouses(tx DBTX) *Houses {
	return &Houses{defaultDBTX: tx}
}

func (h *Houses) Create(ctx context.Context, req housecreate.HouseCreateRequest) (*models.House, error) {
	params := houseServiceToDto(req)
	queries := New(h.getDBTX(ctx))
	newHouse, err := queries.createHouse(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := houseDtoToService(newHouse)
	return &resp, nil
}

func (h *Houses) HouseUpdated(ctx context.Context, houseId int) error {
	queries := New(h.getDBTX(ctx))
	err := queries.checkHouseExistence(ctx, int64(houseId))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ErrHouseNotFound
		}
		return err
	}
	return queries.houseUpdated(ctx, int64(houseId))
}

func (h *Houses) getDBTX(ctx context.Context) DBTX {
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
