package mw

import (
	"backend-bootcamp-assignment-2024/internal/services/auth"
	"net/http"
	"strings"
)

type (
	roleRecognizer interface {
		GetRole(tokenStr string) (auth.UserType, error)
	}
	AuthenticatedMiddleware struct {
		roleRecognizer roleRecognizer
		next           http.Handler
	}
	ModeratorOnlyMiddleware struct {
		roleRecognizer roleRecognizer
		next           http.Handler
	}
)

func NewModeratorOnlyMiddleware(roleRecognizer roleRecognizer, next http.Handler) *ModeratorOnlyMiddleware {
	return &ModeratorOnlyMiddleware{roleRecognizer: roleRecognizer, next: next}
}

func NewAuthenticatedMiddleware(roleRecognizer roleRecognizer, next http.Handler) *AuthenticatedMiddleware {
	return &AuthenticatedMiddleware{roleRecognizer, next}
}

func (mw *AuthenticatedMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, present := getTokenFromR(r)
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
	token, present := getTokenFromR(r)
	if !present {
		writeUnauthorized(w)
		return
	}
	role, err := mw.roleRecognizer.GetRole(token)
	if err == nil && role == auth.Moderator {
		mw.next.ServeHTTP(w, r)
	} else {
		writeUnauthorized(w)
	}
}

func getTokenFromR(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", false
	}
	token, _ := strings.CutPrefix(authHeader, "Bearer")
	return token, true
}

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
