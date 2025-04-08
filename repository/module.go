package repository

import (
	"auth/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"auth/ent"
)

type (
	Datasource struct {
		Connection string
	}

	DatasourceProps struct {
		fx.In
		Client *ent.Client
		Db     *sql.DB
		Logger *zap.SugaredLogger
	}
)

var Module = fx.Provide(
	NewDatasource,
	NewEntClient,
	NewNativeClient,
	NewAuthRepository,
	NewEmployeeRepository,
	NewDepartmentRepository,
)

func NewDatasource(config config.Config) *Datasource {
	user := config.MustGetString("DB_USER")
	pass := config.MustGetString("DB_PASS")
	name := config.MustGetString("DB_NAME")
	host := config.MustGetString("DB_HOST")
	port := config.MustGetInt("DB_PORT")

	conn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		user, pass, name, host, port,
	)

	return &Datasource{Connection: conn}
}

func NewEntClient(lifecycle fx.Lifecycle, config config.Config, logger *zap.SugaredLogger, datasource *Datasource) *ent.Client {
	entLogger := logger.WithOptions(zap.AddCallerSkip(10))

	options := func() []ent.Option {
		if config.MustGetString("PROFILE") == "local" {
			return []ent.Option{
				ent.Debug(),
				ent.Log(func(a ...any) {
					entLogger.Debugln(a...)
				}),
			}
		} else {
			return []ent.Option{}
		}
	}

	client, err := ent.Open("postgres", datasource.Connection, options()...)
	if err != nil {
		panic(err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client
}

func NewNativeClient(lifecycle fx.Lifecycle, datasource *Datasource) *sql.DB {
	db, err := sql.Open("postgres", datasource.Connection)
	if err != nil {
		panic(err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db
}
