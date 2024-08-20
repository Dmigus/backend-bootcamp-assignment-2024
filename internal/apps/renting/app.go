// Package renting содержит backend-сервис, который обрабатывает запросы всех пользователей

package renting

import (
	authController "backend-bootcamp-assignment-2024/internal/controllers/auth"
	"backend-bootcamp-assignment-2024/internal/controllers/mw"
	rentingController "backend-bootcamp-assignment-2024/internal/controllers/renting"
	createflatHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/createflat"
	getflatsHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/getflats"
	housecreateHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/housecreate"
	updateflatHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/updateflat"
	"backend-bootcamp-assignment-2024/internal/providers/hash"
	"backend-bootcamp-assignment-2024/internal/providers/jwt"
	"backend-bootcamp-assignment-2024/internal/providers/postgres"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/flats"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/houses"
	"backend-bootcamp-assignment-2024/internal/providers/postgres/users"
	"backend-bootcamp-assignment-2024/internal/providers/salt"
	"backend-bootcamp-assignment-2024/internal/providers/uuid"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/login"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/register"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/getflats"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/updateflat"
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

var Dependencies = fx.Provide(
	// services
	fx.Annotate(
		login.NewService,
		fx.As(new(authController.LoginService)),
		fx.As(new(mw.RoleRecognizer)),
		fx.As(new(getflatsHandler.RoleRecognizer)),
	),
	fx.Annotate(
		register.NewService,
		fx.As(new(authController.RegisterService)),
	),
	fx.Annotate(
		housecreate.NewHouseService,
		fx.As(new(housecreateHandler.HouseService)),
	),
	fx.Annotate(
		getflats.NewService,
		fx.As(new(getflatsHandler.FlatsService)),
	),
	fx.Annotate(
		createflat.NewService,
		fx.As(new(createflatHandler.FlatsService)),
	),
	fx.Annotate(
		updateflat.NewService,
		fx.As(new(updateflatHandler.FlatService)),
	),

	// repo
	fx.Annotate(
		houses.NewHouses,
		fx.As(new(housecreate.Repository)),
		fx.As(new(createflat.HousesRepo)),
	),
	fx.Annotate(
		flats.NewFlats,
		fx.As(new(getflats.Repository)),
		fx.As(new(createflat.FlatsRepo)),
		fx.As(new(updateflat.Repository)),
	),
	fx.Annotate(
		postgres.NewTxManger,
		fx.As(new(createflat.TxManager)),
		fx.As(new(updateflat.TxManager)),
	),
	fx.Annotate(
		CreateConnToPostgres,
		fx.As(new(houses.DBTX)),
		fx.As(new(flats.DBTX)),
		fx.As(new(postgres.TxBeginner)),
		fx.As(new(users.DBTX)),
	),
	fx.Annotate(
		users.NewUsers,
		fx.As(new(login.Repository)),
		fx.As(new(register.Repository)),
	),

	//auth
	fx.Annotate(
		jwtCodec,
		fx.As(new(login.JWTCodec)),
	),
	fx.Annotate(
		hash.NewBCryptHasher,
		fx.As(new(login.PasswordHasher)),
		fx.As(new(register.PasswordHasher)),
	),
	fx.Annotate(
		salt.NewGenerator,
		fx.As(new(register.SaltGenerator)),
	),
	fx.Annotate(
		uuid.NewGenerator,
		fx.As(new(register.UserIdGenerator)),
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
	getflatsHandler.NewHandler,
	housecreateHandler.NewHandler,
	createflatHandler.NewHandler,
	updateflatHandler.NewHandler,
	generalMux,

	httpServer,
)

var Module = fx.Module("renting",
	Dependencies,
	fx.Invoke(func(*http.Server) {}),
)

func CreateConnToPostgres(config *Config) (*pgxpool.Pool, error) {
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

func jwtCodec(config *Config) *jwt.Codec {
	key := []byte(config.CipherKey)
	return jwt.NewCodec(key)
}

func authHandler(loginService authController.LoginService, registerService authController.RegisterService) http.Handler {
	serverHandler := authController.NewServerHandler(loginService, registerService)
	return authController.Handler(serverHandler)
}

type rentingHandlerParams struct {
	fx.In
	GetFlatsHandler    *getflatsHandler.Handler
	HouseCreateHandler *housecreateHandler.Handler
	CreateFlatHandler  *createflatHandler.Handler
	UpdateFlatHandler  *updateflatHandler.Handler
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
