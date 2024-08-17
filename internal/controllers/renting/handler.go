package renting

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

type (
	houseCreateHandler interface {
		PostHouseCreate(w http.ResponseWriter, r *http.Request)
	}
	getFlatsHandler interface {
		GetHouseId(w http.ResponseWriter, r *http.Request, id HouseId)
	}
	createFlatHandler interface {
		PostFlatCreate(w http.ResponseWriter, r *http.Request)
	}
	updateFlatHandler interface {
		PostFlatUpdate(w http.ResponseWriter, r *http.Request)
	}
	// ServerHandler это контейнер, в который положим все нужные хэндлеры, чтобы соответствовать сгенерированному ServerInterface
	ServerHandler struct {
		houseCreateHandler houseCreateHandler
		getFlatsHandler    getFlatsHandler
		createFlatHandler  createFlatHandler
		updateFlatHandler  updateFlatHandler
	}
)

func (s *ServerHandler) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	s.createFlatHandler.PostFlatCreate(w, r)
}

func (s *ServerHandler) PostFlatUpdate(w http.ResponseWriter, r *http.Request) {
	s.updateFlatHandler.PostFlatUpdate(w, r)
}

func (s *ServerHandler) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	s.houseCreateHandler.PostHouseCreate(w, r)
}

func (s *ServerHandler) GetHouseId(w http.ResponseWriter, r *http.Request, id HouseId) {
	s.getFlatsHandler.GetHouseId(w, r, id)
}

func (s *ServerHandler) PostHouseIdSubscribe(w http.ResponseWriter, r *http.Request, id HouseId) {
	//TODO implement me
	panic("implement me")
}

func NewServerHandler(
	houseCreateHandler houseCreateHandler,
	getFlatsHandler getFlatsHandler,
	createFlatHandler createFlatHandler,
	updateFlatHandler updateFlatHandler,
) *ServerHandler {
	return &ServerHandler{
		houseCreateHandler: houseCreateHandler,
		getFlatsHandler:    getFlatsHandler,
		createFlatHandler:  createFlatHandler,
		updateFlatHandler:  updateFlatHandler,
	}
}

func ReadJsonBody(r *http.Request, dst any) error {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyData, dst)
}

func Write5xxResponse(w http.ResponseWriter, message string) {
	respCode := http.StatusInternalServerError
	body := N5xx{
		Code:    &respCode,
		Message: message,
	}
	jsonBody, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(jsonBody)
}

func FlatModelToDto(resp *models.Flat) Flat {
	return Flat{
		Id:      resp.Id,
		HouseId: resp.HouseId,
		Price:   resp.Price,
		Rooms:   resp.Rooms,
		Status:  statusModelToDto(resp.Status),
	}
}

func statusModelToDto(status models.FlatStatus) Status {
	switch status {
	case models.Created:
		return Created
	case models.OnModerate:
		return OnModeration
	case models.Approved:
		return Approved
	case models.Declined:
		return Declined
	}
	return ""
}
