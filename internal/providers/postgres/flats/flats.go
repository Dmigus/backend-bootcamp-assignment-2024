package flats

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/samber/lo"
)

type Flats struct {
	defaultDBTX DBTX
}

func NewFlats(defaultDBTX DBTX) *Flats {
	return &Flats{defaultDBTX: defaultDBTX}
}

func (f *Flats) CreateFlat(ctx context.Context, req createflat.Request) (*models.Flat, error) {
	queries := New(f.getDBTX(ctx))
	params := createFlatParams{HouseID: int64(req.HouseId), Price: int32(req.Price), Rooms: int32(req.Rooms)}
	flat, err := queries.createFlat(ctx, params)
	if err != nil {
		if isFKViolation(err) {
			return nil, models.ErrHouseNotFound
		}
		return nil, err
	}
	flatService := flatDtoToService(flat)
	return &flatService, nil
}

func (f *Flats) GetFlats(ctx context.Context, houseId int) ([]models.Flat, error) {
	queries := New(f.getDBTX(ctx))
	flats, err := queries.getFlats(ctx, int64(houseId))
	if err != nil {
		return nil, err
	}
	flatsModel := lo.Map(flats, func(flat Flat, _ int) models.Flat {
		return flatDtoToService(flat)
	})
	return flatsModel, nil
}

func (f *Flats) GetApprovedFlats(ctx context.Context, houseId int) ([]models.Flat, error) {
	queries := New(f.getDBTX(ctx))
	flats, err := queries.getApprovedFlats(ctx, int64(houseId))
	if err != nil {
		return nil, err
	}
	flatsModel := lo.Map(flats, func(flat Flat, _ int) models.Flat {
		return flatDtoToService(flat)
	})
	return flatsModel, nil
}

func (f *Flats) GetForUpdate(ctx context.Context, id int) (*models.Flat, error) {
	queries := New(f.getDBTX(ctx))
	flats, err := queries.getFlatForUpdate(ctx, int64(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrFlatNotFound
		}
		return nil, err
	}
	flatService := flatDtoToService(flats)
	return &flatService, nil
}

func (f *Flats) UpdateStatus(ctx context.Context, id int, status models.FlatStatus) (*models.Flat, error) {
	queries := New(f.getDBTX(ctx))
	params := updateStatusParams{ID: int64(id), Status: status.String()}
	updatedFlat, err := queries.updateStatus(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrFlatNotFound
		}
		return nil, err
	}
	flatService := flatDtoToService(updatedFlat)
	return &flatService, nil
}

func (f *Flats) getDBTX(ctx context.Context) DBTX {
	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		return f.defaultDBTX
	}
	dbtx, ok := tx.(DBTX)
	if ok {
		return dbtx
	}
	return f.defaultDBTX
}

func isFKViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23503" {
		return true
	}
	return false
}
