package renting

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
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

func (h *Renting) GetFlats(ctx context.Context, id int) ([]models.Flat, error) {
	//queries := New(h.getDBTX(ctx))
	//
	//if err != nil {
	//	if errors.Is(err, pgx.ErrNoRows) {
	//		err = pkgerrors.Wrap(models.ErrHouseNotFound, "house not found by id "+strconv.Itoa(id))
	//	}
	//	return nil, err
	//}
	return []models.Flat{}, nil
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

func houseServiceToDto(req housecreate.HouseCreateRequest) createHouseParams {
	params := createHouseParams{Address: req.Address, Year: int32(req.Year)}
	if req.Developer != nil {
		params.Developer = pgtype.Text{String: *req.Developer, Valid: true}
	}
	return params
}

func houseDtoToService(dto House) models.House {
	result := models.House{
		Id:        int(dto.ID),
		Address:   dto.Address,
		Year:      int(dto.Year),
		CreatedAt: dto.CreatedAt.Time,
	}
	if dto.Developer.Valid {
		result.Developer = &dto.Developer.String
	}
	if dto.UpdateAt.Valid {
		result.UpdateAt = &dto.UpdateAt.Time
	}
	return result
}
