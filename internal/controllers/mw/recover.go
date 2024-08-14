package mw

import (
	"backend-bootcamp-assignment-2024/internal/controllers/auth"
	"encoding/json"
	"net/http"
)

const unknownError = "unknown error"

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				body := recoverToRespBody(err)
				jsonBody, _ := json.Marshal(body)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(jsonBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func recoverToRespBody(err any) auth.N5xx {
	var responseMessage string
	switch err.(type) {
	case error:
		responseMessage = err.(error).Error()
	case string:
		responseMessage = err.(string)
	default:
		responseMessage = unknownError
	}
	respCode := http.StatusInternalServerError
	return auth.N5xx{
		Code:    &respCode,
		Message: responseMessage,
	}
}
