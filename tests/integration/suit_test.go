//go:build integration

package integration

import (
	"backend-bootcamp-assignment-2024/internal/apps/renting"
	createflatHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/createflat"
	getflatsHandler "backend-bootcamp-assignment-2024/internal/controllers/renting/getflats"
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"time"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const migrationsDir = "../../migrations"

type Suite struct {
	suite.Suite
	Pg           *postgres.PostgresContainer
	ConnToDB     *sql.DB
	module       fx.Option
	cachedPgPool *pgxpool.Pool
}

func (s *Suite) SetupSuite() {
	ctx := context.Background()
	postgresContainer, err := postgres.Run(ctx, "docker.io/postgres:16.2-bullseye",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	s.Require().NoError(err)
	s.Pg = postgresContainer

	dsn, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err)
	s.ConnToDB, err = sql.Open("pgx", dsn)
	s.Require().NoError(err)
	err = migrateUp(ctx, s.ConnToDB)
	s.Require().NoError(err)
	config := &renting.Config{
		PostgresDSN: dsn,
	}
	s.cachedPgPool, err = renting.CreateConnToPostgres(config)
	s.Require().NoError(err)

	s.module = fx.Module("testing module",
		fx.Provide(func() *renting.Config {
			return config
		}),
		renting.Dependencies,
		fx.Replace(s.cachedPgPool),
	)
}

func (s *Suite) TearDownSuite() {
	s.cachedPgPool.Close()
	ctx := context.Background()
	err := s.Pg.Terminate(ctx)
	s.Assert().NoError(err)
}

func migrateUp(ctx context.Context, conn *sql.DB) error {
	return goose.UpContext(ctx, conn, migrationsDir)
}

func (s *Suite) SetupTest() {
	_, err := s.ConnToDB.Exec("TRUNCATE flat;")
	s.Assert().NoError(err)
	_, err = s.ConnToDB.Exec("DELETE FROM house;")
	s.Assert().NoError(err)
}

func (s *Suite) TestFlatCreate() {
	appWithTest := fxtest.New(
		s.T(),
		s.module,
		fx.Invoke(func(service createflatHandler.FlatsService) {
			ctx := context.Background()
			_, err := s.ConnToDB.ExecContext(ctx,
				"INSERT INTO house (id, address, year, developer, created_at) VALUES ($1, $2, $3, $4, clock_timestamp()::timestamp)",
				10, "street", 2000, "developer")
			s.Require().NoError(err)

			req := createflat.Request{HouseId: 10, Price: 100, Rooms: 100}
			_, err = service.CreateFlat(context.Background(), req)
			s.Assert().NoError(err)
			flatRow := s.ConnToDB.QueryRowContext(ctx, "SELECT price, rooms, status FROM flat;")
			var price, rooms int
			var status string
			err = flatRow.Scan(&price, &rooms, &status)
			s.Require().NoError(err)
			s.Assert().Equal(100, price)
			s.Assert().Equal(100, rooms)
			s.Assert().Equal(models.Created.String(), status)
		}),
	)
	appWithTest.RequireStart().RequireStop()
}

func (s *Suite) TestGetFlats() {
	appWithTest := fxtest.New(
		s.T(),
		s.module,
		fx.Invoke(func(service getflatsHandler.FlatsService) {
			ctx := context.Background()
			_, err := s.ConnToDB.ExecContext(ctx,
				"INSERT INTO house (id, address, year, developer, created_at) VALUES ($1, $2, $3, $4, clock_timestamp()::timestamp)",
				10, "street", 2000, "developer")
			s.Require().NoError(err)
			_, err = s.ConnToDB.ExecContext(ctx,
				"INSERT INTO flat (house_id, price, rooms, status) VALUES (10, 200, 3, 'Created'), (10, 300, 4, 'Approved')")
			s.Require().NoError(err)
			returnedForClient, err := service.GetFlats(ctx, 10, models.Client)
			s.Require().NoError(err)
			returnedForModerator, err := service.GetFlats(ctx, 10, models.Moderator)
			s.Require().NoError(err)
			s.Assert().Len(returnedForClient, 1)
			s.Assert().Len(returnedForModerator, 2)
			s.Assert().Equal(models.Approved, returnedForClient[0].Status)
			s.Assert().Equal(300, returnedForClient[0].Price)
			s.Assert().Equal(4, returnedForClient[0].Rooms)
		}),
	)
	appWithTest.RequireStart().RequireStop()
}
