package getflats

import (
	"backend-bootcamp-assignment-2024/internal/controllers/mw"
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"encoding/json"
	"github.com/samber/lo"
	"net/http"
)

type (
	FlatsService interface {
		GetFlats(ctx context.Context, id int, role models.UserRole) ([]models.Flat, error)
	}
	RoleRecognizer interface {
		GetRole(tokenStr string) (models.UserRole, error)
	}
	Handler struct {
		roleRecognizer RoleRecognizer
		service        FlatsService
	}
)

func NewHandler(service FlatsService, roleRecognizer RoleRecognizer) *Handler {
	return &Handler{service: service, roleRecognizer: roleRecognizer}
}

func (h *Handler) GetHouseId(w http.ResponseWriter, r *http.Request, id renting.HouseId) {
	token, present := mw.GetTokenFromRequest(r)
	if !present {
		renting.Write5xxResponse(w, "token is not present")
		return
	}
	role, err := h.roleRecognizer.GetRole(token)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
		return
	}
	serviceResp, err := h.service.GetFlats(r.Context(), id, role)
	if err != nil {
		renting.Write5xxResponse(w, err.Error())
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
	dtos := lo.Map(flats, func(flat models.Flat, _ int) renting.Flat {
		return renting.Flat{
			HouseId: flat.HouseId,
			Id:      flat.Id,
			Price:   flat.Price,
			Rooms:   flat.Rooms,
			Status:  statusModelToDto(flat.Status),
		}
	})
	return json.Marshal(dtos)
}

func statusModelToDto(status models.FlatStatus) renting.Status {
	switch status {
	case models.Created:
		return renting.Created
	case models.OnModerate:
		return renting.OnModeration
	case models.Approved:
		return renting.Approved
	case models.Declined:
		return renting.Declined
	}
	return ""
}
