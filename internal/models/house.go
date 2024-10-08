package models

import (
	"errors"
	"time"
)

var ErrHouseNotFound = errors.New("house is not found")

type House struct {
	Id        int
	Address   string
	Year      int
	Developer *string
	CreatedAt time.Time
	UpdateAt  *time.Time
}
