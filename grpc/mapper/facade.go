package grpcmapper

import (
	"auth/ent"
	"auth/proto/auth"
)

type (
	AuthGrpcMapper interface {
		ConvertUser(from *ent.Employee) *auth.User
	}
)
