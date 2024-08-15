package renting

import (
	authController "backend-bootcamp-assignment-2024/internal/controllers/auth"
	"backend-bootcamp-assignment-2024/internal/controllers/mw"
	rentingController "backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/houses"
	"backend-bootcamp-assignment-2024/internal/services/auth"
	"backend-bootcamp-assignment-2024/internal/services/renting/house"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("renting",
	fx.Provide(
		fx.Annotate(
			authService,
			fx.As(new(authController.Service)),
			fx.As(new(mw.RoleRecognizer)),
		),
		fx.Annotate(
			house.NewHouseService,
			fx.As(new(rentingController.HouseService)),
		),
		fx.Annotate(
			houses.NewHouses,
			fx.As(new(house.Repository)),
		),
		fx.Annotate(
			createConnToPostgres,
			fx.As(new(houses.DBTX)),
		),
		fx.Annotate(
			authHandler,
			fx.ResultTags(`name:"authHandler"`),
		),
		fx.Annotate(
			rentingHandler,
			fx.ResultTags(`name:"rentingHandler"`),
		),
		generalMux,
		httpServer,
	),
	fx.Invoke(func(*http.Server) {}),
)

func createConnToPostgres(config *Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.PostgresDSN)
	if err != nil {
		return nil, err
	}
	conn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return conn, nil
}

func authService(config *Config) *auth.AuthService {
	key := []byte(config.CipherKey)
	return auth.NewAuthService(key)
}

func authHandler(service authController.Service) http.Handler {
	serverHandler := authController.NewServerHandler(service)
	return authController.Handler(serverHandler)
}

func rentingHandler(service rentingController.HouseService) http.Handler {
	serverHandler := rentingController.NewServerHandler(service)
	return rentingController.Handler(serverHandler)
}

type generalMuxParams struct {
	fx.In
	RoleRecognizer mw.RoleRecognizer
	AuthHandler    http.Handler `name:"authHandler"`
	RentingHandler http.Handler `name:"rentingHandler"`
}

func generalMux(params generalMuxParams) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /dummyLogin", params.AuthHandler)
	mux.Handle("POST /login", params.AuthHandler)
	mux.Handle("POST /register", params.AuthHandler)
	mux.Handle("POST /house/", mw.NewModeratorOnlyMiddleware(params.RoleRecognizer, params.RentingHandler))
	return mw.Recovery(mux)
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
