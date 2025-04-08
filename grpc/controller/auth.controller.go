package controller

import (
	grpcservice "auth/grpc/service"
	util "auth/grpc/util"
	"auth/proto/auth"
	"auth/proto/gcommon"
	"context"
)

type AuthGrpcController struct {
	auth.UnimplementedAuthServiceServer
	authService grpcservice.AuthGrpcService
}

func NewAuthGrpcController(
	authService grpcservice.AuthGrpcService,
) *AuthGrpcController {
	return &AuthGrpcController{
		authService: authService,
	}
}

func (c *AuthGrpcController) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.authService.Login)
}

func (c *AuthGrpcController) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.authService.ChangePassword)
}
