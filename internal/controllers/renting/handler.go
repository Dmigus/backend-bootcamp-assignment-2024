package renting

import (
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
	ServerHandler struct {
		houseCreateHandler houseCreateHandler
		getFlatsHandler    getFlatsHandler
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
) *ServerHandler {
	return &ServerHandler{
		houseCreateHandler: houseCreateHandler,
		getFlatsHandler:    getFlatsHandler,
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
