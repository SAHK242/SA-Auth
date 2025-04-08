package grpcvalidator

import (
	"auth/proto/auth"
	"auth/repository"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type departmentGrpcValidator struct {
	logger         *zap.SugaredLogger
	authRepository repository.AuthRepository
}

func (d departmentGrpcValidator) ValidateCreateDepartmentRequest(ctx context.Context, req *auth.CreateDepartmentRequest) error {
	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "department name is required")
	}

	return nil
}

func NewDepartmentValidator(
	logger *zap.SugaredLogger,
) DepartmentValidator {
	return &departmentGrpcValidator{
		logger: logger,
	}
}
