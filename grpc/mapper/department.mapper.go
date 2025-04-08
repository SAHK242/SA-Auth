package grpcmapper

import (
	"auth/ent"
	"auth/proto/auth"
)

type departmentGrpcMapper struct{}

func (a departmentGrpcMapper) ConvertDepartment(from *ent.Department) *auth.Department {
	if from == nil {
		return nil
	}
	to := &auth.Department{
		Id:   from.ID.String(),
		Name: from.Name,
	}
	return to
}

func (a departmentGrpcMapper) ConvertDepartmentSlice(from []*ent.Department) []*auth.Department {
	to := make([]*auth.Department, 0)
	for _, item := range from {
		to = append(to, a.ConvertDepartment(item))
	}
	return to
}

func NewDepartmentGrpcMapper() DepartmentGrpcMapper {
	return &departmentGrpcMapper{}
}
