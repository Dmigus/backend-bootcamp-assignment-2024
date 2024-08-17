package createflat

import (
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat"
	"context"
	"encoding/json"
	"net/http"
)

type (
	FlatsService interface {
		CreateFlat(ctx context.Context, flat createflat.Request) (*models.Flat, error)
	}
	Handler struct {
		service FlatsService
	}
)

func NewHandler(service FlatsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	var req renting.PostFlatCreateJSONRequestBody
	err := renting.ReadJsonBody(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if req.Rooms == nil {
		http.Error(w, "rooms is required", http.StatusBadRequest)
		return
	}
	serviceReq := createflat.Request{HouseId: req.HouseId, Price: req.Price, Rooms: *req.Rooms}
	serviceResp, err := h.service.CreateFlat(r.Context(), serviceReq)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	dto := renting.FlatModelToDto(serviceResp)
	serialized, err := json.Marshal(dto)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)

}
