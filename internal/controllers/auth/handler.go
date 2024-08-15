package auth

import (
	"backend-bootcamp-assignment-2024/internal/services/auth"
	"encoding/json"
	"net/http"
)

type (
	Service interface {
		DummyLogin(ut auth.UserType) (string, error)
	}
	ServerHandler struct {
		service Service
	}
)

func NewServerHandler(service Service) *ServerHandler {
	return &ServerHandler{service: service}
}

func (s *ServerHandler) GetDummyLogin(w http.ResponseWriter, _ *http.Request, params GetDummyLoginParams) {
	var userType auth.UserType
	switch params.UserType {
	case Client:
		userType = auth.Client
	case Moderator:
		userType = auth.Moderator
	default:
		http.Error(w, "invalid user_type", http.StatusBadRequest)
		return
	}
	token, err := s.service.DummyLogin(userType)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	structured := struct {
		AuthToken string `json:"auth_token"`
	}{token}
	serialized, err := json.Marshal(structured)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}

func (s *ServerHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
