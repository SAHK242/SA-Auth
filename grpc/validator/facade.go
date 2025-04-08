package grpcvalidator

import (
	"auth/proto/auth"
	"context"
)

type (
	AuthValidator interface {
		ValidateCreateEmployeeRequest(ctx context.Context, req *auth.CreateEmployeeRequest) error
	}

	DepartmentValidator interface {
		ValidateCreateDepartmentRequest(ctx context.Context, req *auth.CreateDepartmentRequest) error
	}
)
