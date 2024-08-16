package models

type FlatStatus int

const (
	Created FlatStatus = iota + 1
	OnModerate
	Approved
	Declined
)

type Flat struct {
}
