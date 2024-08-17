package renting

import (
	authController "backend-bootcamp-assignment-2024/internal/controllers/auth"
	"backend-bootcamp-assignment-2024/internal/controllers/mw"
	rentingController "backend-bootcamp-assignment-2024/internal/controllers/renting"
	"backend-bootcamp-assignment-2024/internal/controllers/renting/createflat"
	"backend-bootcamp-assignment-2024/internal/controllers/renting/getflats"
	"backend-bootcamp-assignment-2024/internal/controllers/renting/housecreate"
	"backend-bootcamp-assignment-2024/internal/controllers/renting/updateflat"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/flats"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/houses"
	"backend-bootcamp-assignment-2024/internal/services/auth"
	createflat2 "backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat"
	getflats2 "backend-bootcamp-assignment-2024/internal/services/renting/usecases/getflats"
	housecreate2 "backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	updateflat2 "backend-bootcamp-assignment-2024/internal/services/renting/usecases/updateflat"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("renting",
	fx.Provide(
		// services
		fx.Annotate(
			authService,
			fx.As(new(authController.Service)),
			fx.As(new(mw.RoleRecognizer)),
			fx.As(new(getflats.RoleRecognizer)),
		),
		fx.Annotate(
			housecreate2.NewHouseService,
			fx.As(new(housecreate.HouseService)),
		),
		fx.Annotate(
			getflats2.NewService,
			fx.As(new(getflats.FlatsService)),
		),
		fx.Annotate(
			createflat2.NewService,
			fx.As(new(createflat.FlatsService)),
		),
		fx.Annotate(
			updateflat2.NewService,
			fx.As(new(updateflat.FlatService)),
		),

		// repo
		fx.Annotate(
			houses.NewHouses,
			fx.As(new(housecreate2.Repository)),
			fx.As(new(createflat2.HousesRepo)),
		),
		fx.Annotate(
			flats.NewFlats,
			fx.As(new(getflats2.Repository)),
			fx.As(new(createflat2.FlatsRepo)),
			fx.As(new(updateflat2.Repository)),
		),
		fx.Annotate(
			postgres.NewTxManger,
			fx.As(new(createflat2.TxManager)),
			fx.As(new(updateflat2.TxManager)),
		),
		fx.Annotate(
			createConnToPostgres,
			fx.As(new(houses.DBTX)),
			fx.As(new(flats.DBTX)),
			fx.As(new(postgres.TxBeginner)),
		),

		// controllers
		fx.Annotate(
			authHandler,
			fx.ResultTags(`name:"authHandler"`),
		),
		fx.Annotate(
			rentingHandler,
			fx.ResultTags(`name:"rentingHandler"`),
		),
		getflats.NewHandler,
		housecreate.NewHandler,
		createflat.NewHandler,
		updateflat.NewHandler,
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

type rentingHandlerParams struct {
	fx.In
	GetFlatsHandler    *getflats.Handler
	HouseCreateHandler *housecreate.Handler
	CreateFlatHandler  *createflat.Handler
	UpdateFlatHandler  *updateflat.Handler
}

func rentingHandler(handlers rentingHandlerParams) http.Handler {
	serverHandler := rentingController.NewServerHandler(
		handlers.HouseCreateHandler,
		handlers.GetFlatsHandler,
		handlers.CreateFlatHandler,
		handlers.UpdateFlatHandler,
	)
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
	rentingModeratorOnly := mw.NewModeratorOnlyMiddleware(params.RoleRecognizer, params.RentingHandler)
	rentingAuthenticated := mw.NewAuthenticatedMiddleware(params.RoleRecognizer, params.RentingHandler)
	mux.Handle("POST /house/create", rentingModeratorOnly)
	mux.Handle("GET /house/{id}", rentingAuthenticated)
	mux.Handle("POST /flat/create", rentingAuthenticated)
	mux.Handle("POST /flat/update", rentingModeratorOnly)
	mux.Handle("POST /house/{id}/subscribe", rentingAuthenticated)
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
