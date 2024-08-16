package getflats

import (
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"errors"
	"net/http"
)

type (
	FlatsService interface {
		GetFlats(ctx context.Context, id int) ([]models.Flat, error)
	}
	Handler struct {
		service FlatsService
	}
)

func NewHandler(service FlatsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetHouseId(w http.ResponseWriter, r *http.Request, id renting.HouseId) {
	serviceResp, err := h.service.GetFlats(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrHouseNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			renting.Write5xxResponse(w, err.Error())
		}
		return
	}
	serialized, err := serializeFlats(serviceResp)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}

func serializeFlats(flats []models.Flat) ([]byte, error) {
	return []byte{}, nil
}
