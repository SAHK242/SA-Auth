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
		Id:          from.EmployeeID.String(),
		FirstName:   from.FirstName,
		LastName:    from.LastName,
		Code:        from.Code,
		Gender:      gcommon.Gender(from.Gender),
		DateOfBirth: from.DateOfBirth.Unix(),
		Address:     from.Address,
		StartDate:   from.StartDate.Unix(),
		EndDate:     from.EndDate.Unix(),
		PhoneNumber: from.PhoneNumber,
		DegreeName:  from.DegreeName,
		DegreeYear:  from.DegreeYear,
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

func NewAuthGrpcMapper() AuthGrpcMapper {
	return &authGrpcMapper{}
}
