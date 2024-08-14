package renting

import (
	authController "backend-bootcamp-assignment-2024/internal/controllers/auth"
	"backend-bootcamp-assignment-2024/internal/services/auth"
	"context"
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("renting",
	fx.Provide(
		fx.Annotate(
			authService,
			fx.As(new(authController.Service)),
		),
		httpHandler,
		httpServer,
	),
	fx.Invoke(func(*http.Server) {}),
)

func authService(config *Config) *auth.AuthService {
	key := []byte(config.CipherKey)
	return auth.NewAuthService(key)
}

func httpHandler(service authController.Service) http.Handler {
	serverHandler := authController.NewServerHandler(service)
	return authController.Handler(serverHandler)
}

func httpServer(lc fx.Lifecycle, config *Config, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: handler,
	}
	lc.Append(fx.StartStopHook(
		func() {
			go func() {
				_ = server.ListenAndServe()
			}()
		},
		func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	))
	return server
}
