package updateflat

import (
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/updateflat"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type (
	FlatService interface {
		UpdateStatus(ctx context.Context, id int, newStatus models.FlatStatus) (*models.Flat, error)
	}
	Handler struct {
		service FlatService
	}
)

func NewHandler(service FlatService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) PostFlatUpdate(w http.ResponseWriter, r *http.Request) {
	var req renting.PostFlatUpdateJSONBody
	err := renting.ReadJsonBody(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if req.Status == nil {
		http.Error(w, "status is required", http.StatusBadRequest)
		return
	}
	status, err := dtoStatusToModel(*req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	serviceResp, err := h.service.UpdateStatus(r.Context(), req.Id, status)
	if err != nil {
		if errors.Is(err, updateflat.ErrWrongFlatStatus) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			renting.Write5xxResponse(w, err.Error())
		}
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

func dtoStatusToModel(dto renting.Status) (models.FlatStatus, error) {
	switch dto {
	case renting.Created:
		return models.Created, nil
	case renting.OnModeration:
		return models.OnModerate, nil
	case renting.Approved:
		return models.Approved, nil
	case renting.Declined:
		return models.Declined, nil
	default:
		return models.FlatStatus(0), errors.New("invalid status")
	}

}
