package models

import "errors"

type FlatStatus int

var ErrFlatNotFound = errors.New("flat is not found")

const (
	Created FlatStatus = iota + 1
	OnModerate
	Approved
	Declined
)

func (f FlatStatus) String() string {
	switch f {
	case Created:
		return "Created"
	case OnModerate:
		return "OnModerate"
	case Approved:
		return "Approved"
	case Declined:
		return "Declined"
	}
	return ""
}

type Flat struct {
	Id      int
	HouseId int
	Price   int
	Rooms   int
	Status  FlatStatus
}
