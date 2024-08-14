package renting

import (
	"backend-bootcamp-assignment-2024/internal/services/renting/house"
	"context"
	"net/http"
)

type (
	HouseService interface {
		CreateHouse(ctx context.Context, req house.HouseCreateRequest) (house.HouseCreateResponse, error)
		GetHouseByID(ctx context.Context, id int) (house.HouseCreateResponse, error)
	}
	ServerHandler struct {
		houseService HouseService
	}
)

func (s *ServerHandler) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerHandler) PostFlatUpdate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerHandler) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerHandler) GetHouseId(w http.ResponseWriter, r *http.Request, id HouseId) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerHandler) PostHouseIdSubscribe(w http.ResponseWriter, r *http.Request, id HouseId) {
	//TODO implement me
	panic("implement me")
}

func NewServerHandler(houseService HouseService) *ServerHandler {
	return &ServerHandler{houseService: houseService}
}
