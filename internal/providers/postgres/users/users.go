package users

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/login"
	registerUsecase "backend-bootcamp-assignment-2024/internal/services/auth/usecases/register"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Users struct {
	queries *Queries
}

func NewUsers(tx DBTX) *Users {
	return &Users{queries: New(tx)}
}

func (u *Users) Add(ctx context.Context, req registerUsecase.Request) error {
	params := registerParams{
		ID:           pgtype.UUID{Bytes: req.Id, Valid: true},
		Email:        req.Email,
		Salt:         req.Salt,
		PasswordHash: req.HashedPassword,
		Role:         req.Role.String(),
	}
	err := u.queries.register(ctx, params)
	if err != nil {
		if isDuplicateViolation(err) {
			return registerUsecase.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

func (u *Users) GetAuthData(ctx context.Context, userId models.UserId) (*login.AuthData, error) {
	id := pgtype.UUID{Bytes: userId, Valid: true}
	data, err := u.queries.getAuthData(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, login.ErrUserNotFound
		}
		return nil, err
	}
	role, err := convertRole(data.Role)
	if err != nil {
		return nil, err
	}
	return &login.AuthData{
		Salt: data.Salt,
		Hash: data.PasswordHash,
		Role: role,
	}, nil
}

func convertRole(roleStr string) (models.UserRole, error) {
	switch roleStr {
	case models.ClientRoleName:
		return models.Client, nil
	case models.ModeratorRoleName:

		return models.Moderator, nil
	default:
		return models.UserRole(0), models.ErrUnknownRole
	}
}

func isDuplicateViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}
