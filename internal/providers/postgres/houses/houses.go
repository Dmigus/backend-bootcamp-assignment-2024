package houses

import (
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/services/renting/house"
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	pkgerrors "github.com/pkg/errors"
)

type Houses struct {
	defaultDBTX DBTX
}

func NewHouses(tx DBTX) *Houses {
	return &Houses{defaultDBTX: tx}
}

func (h *Houses) Create(ctx context.Context, req house.HouseCreateRequest) (*house.HouseCreateResponse, error) {
	// TODO: разобраться с null значениями
	params := createHouseParams{Address: req.Address, Year: int32(req.Year), Developer: pgtype.Text{String: req.Developer}}
	queries := New(h.getDBTX(ctx))
	newHouse, err := queries.createHouse(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := houseDtoToService(newHouse)
	return &resp, nil
}

func (h *Houses) Get(ctx context.Context, id int) (*house.HouseCreateResponse, error) {
	queries := New(h.getDBTX(ctx))
	returned, err := queries.getOrder(ctx, int64(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = pkgerrors.Wrap(house.ErrHouseNotFound, "house not found by id "+strconv.Itoa(id))
		}
		return nil, err
	}
	resp := houseDtoToService(returned)
	return &resp, nil
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

func houseDtoToService(dto House) house.HouseCreateResponse {
	return house.HouseCreateResponse{
		Id:        int(dto.ID),
		Address:   dto.Address,
		Year:      int(dto.Year),
		Developer: dto.Developer.String,
		CreatedAt: dto.CreatedAt.Time,
		UpdateAt:  dto.UpdateAt.Time,
	}
}
