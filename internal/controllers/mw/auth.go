package mw

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"net/http"
	"strings"
)

type (
	RoleRecognizer interface {
		GetRole(tokenStr string) (models.UserRole, error)
	}
	AuthenticatedMiddleware struct {
		roleRecognizer RoleRecognizer
		next           http.Handler
	}
	ModeratorOnlyMiddleware struct {
		roleRecognizer RoleRecognizer
		next           http.Handler
	}
)

func NewModeratorOnlyMiddleware(roleRecognizer RoleRecognizer, next http.Handler) *ModeratorOnlyMiddleware {
	return &ModeratorOnlyMiddleware{roleRecognizer: roleRecognizer, next: next}
}

func NewAuthenticatedMiddleware(roleRecognizer RoleRecognizer, next http.Handler) *AuthenticatedMiddleware {
	return &AuthenticatedMiddleware{roleRecognizer, next}
}

func (mw *AuthenticatedMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, present := GetTokenFromRequest(r)
	if !present {
		writeUnauthorized(w)
		return
	}
	_, err := mw.roleRecognizer.GetRole(token)
	if err == nil {
		mw.next.ServeHTTP(w, r)
	} else {
		writeUnauthorized(w)
	}
}

func (mw *ModeratorOnlyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, present := GetTokenFromRequest(r)
	if !present {
		writeUnauthorized(w)
		return
	}
	role, err := mw.roleRecognizer.GetRole(token)
	if err == nil && role == models.Moderator {
		mw.next.ServeHTTP(w, r)
	} else {
		writeUnauthorized(w)
	}
}

func GetTokenFromRequest(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", false
	}
	token, _ := strings.CutPrefix(authHeader, "Bearer ")
	return token, true
}

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
