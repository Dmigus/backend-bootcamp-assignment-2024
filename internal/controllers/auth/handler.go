package auth

import (
	"backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/login"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/register"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	uuidGoogle "github.com/google/uuid"
)

type (
	LoginService interface {
		DummyLogin(ut models.UserRole) (string, error)
		Login(ctx context.Context, userId models.UserId, password string) (string, error)
	}
	RegisterService interface {
		Register(ctx context.Context, email, password string, role models.UserRole) (*models.UserId, error)
	}
	ServerHandler struct {
		loginService    LoginService
		registerService RegisterService
	}
)

func NewServerHandler(loginService LoginService, registerService RegisterService) *ServerHandler {
	return &ServerHandler{loginService: loginService, registerService: registerService}
}

func (s *ServerHandler) GetDummyLogin(w http.ResponseWriter, _ *http.Request, params GetDummyLoginParams) {
	var userType models.UserRole
	switch params.UserType {
	case Client:
		userType = models.Client
	case Moderator:
		userType = models.Moderator
	default:
		http.Error(w, "invalid user_type", http.StatusBadRequest)
		return
	}
	token, err := s.loginService.DummyLogin(userType)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	respBody := struct {
		AuthToken string `json:"token"`
	}{token}
	writeJsonResponse(w, respBody)
}

func (s *ServerHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id       *string `json:"id"`
		Password *string `json:"password"`
	}
	err := renting.ReadJsonBody(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Id == nil {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	if req.Password == nil {
		http.Error(w, "password required", http.StatusBadRequest)
		return
	}
	userId, err := userIdFromString(*req.Id)
	if err != nil {
		http.Error(w, "wrong id format", http.StatusBadRequest)
		return
	}
	token, err := s.loginService.Login(r.Context(), userId, *req.Password)
	if err != nil {
		if errors.Is(err, login.ErrInvalidCredentials) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		write5xxResponse(w, err.Error())
		return
	}
	respBody := struct {
		AuthToken string `json:"token"`
	}{token}
	writeJsonResponse(w, respBody)
}

func (s *ServerHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	req, err := postRegisterRequestFromR(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var userType models.UserRole
	switch UserType(*req.UserType) {
	case Client:
		userType = models.Client
	case Moderator:
		userType = models.Moderator
	default:
		http.Error(w, "invalid user_type", http.StatusBadRequest)
		return
	}
	uuid, err := s.registerService.Register(r.Context(), *req.Email, *req.Password, userType)
	if err != nil {
		if errors.Is(err, register.ErrUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		write5xxResponse(w, err.Error())
		return
	}
	respBody := struct {
		UserId string `json:"user_id"`
	}{userIdToString(*uuid)}
	writeJsonResponse(w, respBody)
}

func writeJsonResponse(w http.ResponseWriter, v interface{}) {
	serialized, err := json.Marshal(v)
	if err != nil {
		write5xxResponse(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(serialized)
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

type postRegisterRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	UserType *string `json:"user_type"`
}

func postRegisterRequestFromR(r *http.Request) (*postRegisterRequest, error) {
	var req postRegisterRequest
	err := renting.ReadJsonBody(r, &req)
	if err != nil {
		return nil, fmt.Errorf("invalid body")
	}
	if req.Email == nil {
		return nil, fmt.Errorf("email required")
	}
	if req.Password == nil {
		return nil, fmt.Errorf("password required")
	}
	if req.UserType == nil {
		return nil, fmt.Errorf("user_type required")
	}
	return &req, nil
}

func userIdFromString(uuidStr string) (models.UserId, error) {
	uuidG, err := uuidGoogle.Parse(uuidStr)
	if err != nil {
		return models.UserId{}, err
	}
	bytes, _ := uuidG.MarshalBinary()
	var uuid models.UserId
	copy(uuid[:], bytes)
	return uuid, nil
}

func userIdToString(uuid models.UserId) string {
	uuidG, _ := uuidGoogle.FromBytes(uuid[:])
	return uuidG.String()
}
