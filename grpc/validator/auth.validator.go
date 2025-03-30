package grpcvalidator

import (
	"auth/proto/auth"
	"auth/proto/gcommon"
	"auth/repository"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authGrpcValidator struct {
	logger         *zap.SugaredLogger
	authRepository repository.AuthRepository
}

func (v *authGrpcValidator) ValidateCreateEmployeeRequest(ctx context.Context, req *auth.CreateEmployeeRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	emailExist, err := v.authRepository.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return status.Error(codes.Internal, "failed to check email existence")
	}
	if emailExist {
		return status.Error(codes.AlreadyExists, "email already exists")
	}
	if req.Type != auth.EmployeeType_EMPLOYEE_TYPE_DOCTOR && req.Type != auth.EmployeeType_EMPLOYEE_TYPE_NURSE {
		return status.Error(codes.InvalidArgument, "invalid employee type")
	}

	if req.FirstName == "" && req.LastName == "" {
		return status.Error(codes.InvalidArgument, "first name and last name are required")
	}

	if req.Gender != gcommon.Gender_GENDER_FEMALE && req.Gender != gcommon.Gender_GENDER_MALE && req.Gender != gcommon.Gender_GENDER_OTHER {
		return status.Error(codes.InvalidArgument, "gender type is invalid")
	}

	if req.StartDate == 0 {
		return status.Error(codes.InvalidArgument, "start date is required")
	}

	if req.DegreeName == "" {
		return status.Error(codes.InvalidArgument, "degree name is required")
	}

	if req.DegreeYear == 0 {
		return status.Error(codes.InvalidArgument, "degree year is required")
	}

	if req.DepartmentId == "" {
		return status.Error(codes.InvalidArgument, "department id is required")
	}
	return nil
}

func NewAuthGrpcValidator(
	logger *zap.SugaredLogger,
	authRepository repository.AuthRepository,
) AuthValidator {
	return &authGrpcValidator{
		logger:         logger,
		authRepository: authRepository,
	}
}
