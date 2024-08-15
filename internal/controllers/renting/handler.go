package renting

import (
	"backend-bootcamp-assignment-2024/internal/services/renting/house"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type (
	HouseService interface {
		CreateHouse(ctx context.Context, req house.HouseCreateRequest) (*house.HouseCreateResponse, error)
		GetHouseByID(ctx context.Context, id int) (*house.HouseCreateResponse, error)
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
	var req PostHouseCreateJSONBody
	err := readJsonBody(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	serviceReq := house.HouseCreateRequest{Address: req.Address, Year: req.Year, Developer: *req.Developer}
	serviceResp, err := s.houseService.CreateHouse(r.Context(), serviceReq)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	resp := House{
		Address:   serviceResp.Address,
		Year:      serviceResp.Year,
		Developer: &serviceResp.Developer,
		Id:        serviceResp.Id,
		CreatedAt: &serviceResp.CreatedAt,
		UpdateAt:  &serviceResp.UpdateAt,
	}
	serialized, err := json.Marshal(resp)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)
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

func readJsonBody(r *http.Request, dst any) error {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyData, dst)
}

func write5xxResponse(w http.ResponseWriter, message string) {
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
