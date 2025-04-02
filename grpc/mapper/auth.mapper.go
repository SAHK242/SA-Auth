package grpcmapper

import (
	"auth/ent"
	"auth/proto/auth"
	"auth/proto/gcommon"
)

type authGrpcMapper struct{}

func (a authGrpcMapper) ConvertUser(from *ent.Employee) *auth.User {
	if from == nil {
		return nil
	}
	to := &auth.User{
		Id:          from.ID.String(),
		FirstName:   from.FirstName,
		LastName:    from.LastName,
		Code:        from.Code,
		Gender:      gcommon.Gender(from.Gender),
		DateOfBirth: from.DateOfBirth.Unix(),
		Address:     from.Address,
		StartDate:   from.StartDate.Unix(),
		PhoneNumber: from.PhoneNumber,
		DegreeName:  from.DegreeName,
		DegreeYear:  from.DegreeYear,
	}

	if from.EndDate != nil {
		to.EndDate = from.EndDate.Unix()
	}

	if from.Edges.Department != nil {
		department := &auth.Department{
			Id:   from.Edges.Department.ID.String(),
			Name: from.Edges.Department.Name,
		}
		to.Department = department
	}
	return to
}

func (a authGrpcMapper) ConvertUserSlice(from []*ent.Employee) []*auth.User {
	to := make([]*auth.User, 0)
	for _, item := range from {
		to = append(to, a.ConvertUser(item))
	}
	return to
}

func (a authGrpcMapper) ConvertDepartment(from *ent.Department) *auth.Department {
	if from == nil {
		return nil
	}
	to := &auth.Department{
		Id:   from.ID.String(),
		Name: from.Name,
	}
	return to
}

func (a authGrpcMapper) ConvertDepartmentSlice(from []*ent.Department) []*auth.Department {
	to := make([]*auth.Department, 0)
	for _, item := range from {
		to = append(to, a.ConvertDepartment(item))
	}
	return to
}

func NewAuthGrpcMapper() AuthGrpcMapper {
	return &authGrpcMapper{}
}
