package orders

import (
	"backend-bootcamp-assignment-2024/internal/services/renting/house"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type Houses struct {
	queries *Queries
}

func NewHouses(tx DBTX) *Houses {
	return &Houses{queries: New(tx)}
}

func (h *Houses) Create(ctx context.Context, req house.HouseCreateRequest) (house.HouseCreateResponse, error) {
	_ = createHouseParams{Address: req.Address, Year: int32(req.Year), Developer: pgtype.Text{String: req.Developer}}
	return house.HouseCreateResponse{}, nil
}
