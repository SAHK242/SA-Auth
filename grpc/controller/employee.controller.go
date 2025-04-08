package controller

import (
	grpcservice "auth/grpc/service"
	util "auth/grpc/util"
	"auth/proto/auth"
	"auth/proto/gcommon"
	"context"
)

type EmployeeGrpcController struct {
	auth.UnimplementedEmployeeServiceServer
	employeeService grpcservice.EmployeeGrpcService
}

func NewEmployeeGrpcController(
	employeeService grpcservice.EmployeeGrpcService,
) *EmployeeGrpcController {
	return &EmployeeGrpcController{
		employeeService: employeeService,
	}
}

func (c *EmployeeGrpcController) CreateEmployee(ctx context.Context, req *auth.CreateEmployeeRequest) (*gcommon.EmptyResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.employeeService.CreateEmployee)
}

func (c *EmployeeGrpcController) ListEmployee(ctx context.Context, req *auth.ListEmployeeRequest) (*auth.ListEmployeeResponse, error) {
	return util.WithSafeErrBiFunction(ctx, req, c.employeeService.ListEmployee)
}
