package updateflat

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"errors"
)

var ErrWrongFlatStatus = errors.New("wrong flat status for update")

type (
	Repository interface {
		GetForUpdate(context.Context, int) (*models.Flat, error)
		UpdateStatus(context.Context, int, models.FlatStatus) (*models.Flat, error)
	}
	TxManager interface {
		WithinTransaction(ctx context.Context, f func(context.Context) bool) error
	}
	Service struct {
		repo Repository
		txm  TxManager
	}
)

func NewService(repo Repository, txm TxManager) *Service {
	return &Service{repo: repo, txm: txm}
}

func (s *Service) UpdateStatus(ctx context.Context, id int, newStatus models.FlatStatus) (*models.Flat, error) {
	var serviceErr error
	var updatedFlat *models.Flat
	txErr := s.txm.WithinTransaction(ctx, func(txCtx context.Context) bool {
		updatedFlat, serviceErr = s.performStatusChange(txCtx, id, newStatus)
		return serviceErr == nil
	})
	if serviceErr != nil {
		return nil, serviceErr
	}
	if txErr != nil {
		return nil, txErr
	}
	return updatedFlat, nil
}

func (s *Service) performStatusChange(ctx context.Context, id int, newStatus models.FlatStatus) (*models.Flat, error) {
	flat, err := s.repo.GetForUpdate(ctx, id)
	if err != nil {
		return nil, err
	}
	if !isStatusUpdateAllowed(flat.Status, newStatus) {
		err = ErrWrongFlatStatus
		return nil, err
	}
	return s.repo.UpdateStatus(ctx, id, newStatus)
}

func isStatusUpdateAllowed(old, new models.FlatStatus) bool {
	if old == models.Created && new == models.OnModerate {
		return true
	}
	if old == models.OnModerate && new == models.Approved {
		return true
	}
	if old == models.OnModerate && new == models.Declined {
		return true
	}
	return false
}
