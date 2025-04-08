package grpcmapper

import "go.uber.org/fx"

var Module = fx.Provide(NewAuthGrpcMapper, NewDepartmentGrpcMapper)
