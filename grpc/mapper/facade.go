package grpcmapper

import (
	"auth/ent"
	"auth/proto/auth"
)

type (
	AuthGrpcMapper interface {
		ConvertUser(from *ent.Employee) *auth.User
		ConvertUserSlice(from []*ent.Employee) []*auth.User
	}

	DepartmentGrpcMapper interface {
		ConvertDepartment(from *ent.Department) *auth.Department
		ConvertDepartmentSlice(from []*ent.Department) []*auth.Department
	}
)
