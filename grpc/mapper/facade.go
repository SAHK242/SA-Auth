package grpcmapper

import (
	"auth/ent"
	"auth/proto/auth"
)

type (
	AuthGrpcMapper interface {
		ConvertUser(from *ent.Employee) *auth.User
		ConvertUserSlice(from []*ent.Employee) []*auth.User
		ConvertDepartment(from *ent.Department) *auth.Department
		ConvertDepartmentSlice(from []*ent.Department) []*auth.Department
	}
)
