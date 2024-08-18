package models

import "errors"

type UserRole int

var ErrUnknownRole = errors.New("unknown role")

const (
	Client UserRole = iota + 1
	Moderator
	ClientRoleName    = "client"
	ModeratorRoleName = "moderator"
)

func (r UserRole) String() string {
	switch r {
	case Client:
		return "client"
	case Moderator:
		return "moderator"
	default:
		return ""
	}
}

type AuthClaims struct {
	Role UserRole
	Name string
}
