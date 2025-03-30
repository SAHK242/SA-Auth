package grpcservice

import (
	"auth/proto/auth"
	"auth/proto/gcommon"
	"context"
)

type (
	AuthGrpcService interface {
		Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error)
		ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*gcommon.EmptyResponse, error)
		CreateEmployee(ctx context.Context, request *auth.CreateEmployeeRequest) (*gcommon.EmptyResponse, error)
		CreateDepartment(ctx context.Context, request *auth.CreateDepartmentRequest) (*gcommon.EmptyResponse, error)
		ListEmployee(ctx context.Context, request *auth.ListEmployeeRequest) (*auth.ListEmployeeResponse, error)
		ListDepartment(ctx context.Context, request *auth.ListDepartmentRequest) (*auth.ListDepartmentResponse, error)
	}
)
