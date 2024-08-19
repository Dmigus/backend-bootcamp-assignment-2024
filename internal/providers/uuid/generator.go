package uuid

import (
	"backend-bootcamp-assignment-2024/internal/models"
	uuidGoogle "github.com/google/uuid"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) NewUserId() models.UserId {
	uuidG, _ := uuidGoogle.NewRandom()
	bytes, _ := uuidG.MarshalBinary()
	var uuid models.UserId
	copy(uuid[:], bytes)
	return uuid
}
