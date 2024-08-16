package housecreate

import (
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	"context"
	"encoding/json"
	"net/http"
)

type (
	HouseService interface {
		CreateHouse(ctx context.Context, req housecreate.HouseCreateRequest) (*models.House, error)
	}
	Handler struct {
		service HouseService
	}
)

func NewHandler(service HouseService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	var req renting.PostHouseCreateJSONBody
	err := renting.ReadJsonBody(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	serviceReq := housecreate.HouseCreateRequest{Address: req.Address, Year: req.Year, Developer: req.Developer}
	serviceResp, err := h.service.CreateHouse(r.Context(), serviceReq)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	serialized, err := serializeHouseCreateResp(serviceResp)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}

func serializeHouseCreateResp(resp *models.House) ([]byte, error) {
	dto := renting.House{
		Address:   resp.Address,
		Year:      resp.Year,
		Developer: resp.Developer,
		Id:        resp.Id,
		CreatedAt: &resp.CreatedAt,
		UpdateAt:  resp.UpdateAt,
	}
	return json.Marshal(dto)
}
