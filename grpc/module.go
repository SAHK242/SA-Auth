package grpc

import (
	mapper "auth/grpc/mapper"
	"auth/grpc/server"
	service "auth/grpc/service"
	"auth/interceptor"

	"go.uber.org/fx"

	validator "auth/grpc/validator"

	"auth/grpc/controller"
)

var Module = fx.Module("grpc",
	interceptor.Module,
	mapper.Module,
	validator.Module,
	service.Module,
	controller.Module,
	server.Module,
)
