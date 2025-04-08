package controller

import (
	grpcservice "auth/grpc/service"
	util "auth/grpc/util"
	"auth/proto/auth"
	"auth/proto/gcommon"
	"context"
)

type DepartmentGrpcController struct {
	auth.UnimplementedDepartmentServiceServer
	departmentService grpcservice.DepartmentGrpcService
}

func NewDepartmentGrpcController(
	departmentService grpcservice.DepartmentGrpcService,
) *DepartmentGrpcController {
	return &DepartmentGrpcController{
		departmentService: departmentService,
	}
}

func (c *DepartmentGrpcController) CreateDepartment(ctx context.Context, req *auth.CreateDepartmentRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.departmentService.CreateDepartment)
}

func (c *DepartmentGrpcController) ListDepartment(ctx context.Context, req *auth.ListDepartmentRequest) (*auth.ListDepartmentResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.departmentService.ListDepartment)
}
