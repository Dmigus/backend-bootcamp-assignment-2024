// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package houses

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Flat struct {
	ID      int64
	HouseID int64
	Price   int32
	Rooms   int32
	Status  string
}

type House struct {
	ID        int64
	Address   string
	Year      int32
	Developer pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdateAt  pgtype.Timestamp
}

type User struct {
	ID           pgtype.UUID
	Email        string
	Salt         string
	PasswordHash string
	Role         string
}
