package models

import "errors"

type UserRole int

var ErrUnknownRole = errors.New("unknown role")

const (
	Client UserRole = iota + 1
	Moderator
)
