package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"auth/common"
	"auth/grpc"
	"auth/repository"

	grpcserver "google.golang.org/grpc"
)

func Start() {
	load()
}

func load() {
	// Initialize Zap logger
	logger, _ := zap.NewProduction() // Use zap.NewDevelopment() for development
	defer logger.Sync()              // Flush logs before exit

	sugar := logger.Sugar()

	app := fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zap.WarnLevel)),
			}
		}),
		fx.Supply(sugar), // Provide the SugaredLogger to the DI container
		common.Module,
		repository.Module,
		grpc.Module,
		fx.Invoke(func(*grpcserver.Server) {}),
	)

	app.Run()
}
