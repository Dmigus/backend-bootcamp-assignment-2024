// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAuthData = `-- name: getAuthData :one
SELECT salt, password_hash, role
FROM "user"
WHERE id = $1
`

type getAuthDataRow struct {
	Salt         []byte
	PasswordHash string
	Role         string
}

func (q *Queries) getAuthData(ctx context.Context, id pgtype.UUID) (getAuthDataRow, error) {
	row := q.db.QueryRow(ctx, getAuthData, id)
	var i getAuthDataRow
	err := row.Scan(&i.Salt, &i.PasswordHash, &i.Role)
	return i, err
}

const register = `-- name: register :exec
INSERT INTO "user" (id, email, salt, password_hash, role)
VALUES ($1, $2, $3, $4, $5)
`

type registerParams struct {
	ID           pgtype.UUID
	Email        string
	Salt         []byte
	PasswordHash string
	Role         string
}

func (q *Queries) register(ctx context.Context, arg registerParams) error {
	_, err := q.db.Exec(ctx, register,
		arg.ID,
		arg.Email,
		arg.Salt,
		arg.PasswordHash,
		arg.Role,
	)
	return err
}
