package repository

import (
	"auth/ent"
	"auth/proto/auth"
	"context"
)

type (

	// AuthRepository I write all in this repository, better to split it into multiple repositories

	AuthRepository interface {
		FindByUsername(ctx context.Context, username string) (*ent.Auth, error)
		UpdatePassword(ctx context.Context, tx *ent.Tx, username string, password string) error
		UpdateAccountState(ctx context.Context, tx *ent.Tx, username string, state int32) error
		ExistsByEmail(ctx context.Context, email string) (bool, error)
		CreateEmployee(ctx context.Context, tx *ent.Tx, request *auth.CreateEmployeeRequest) error
		CreateDepartment(ctx context.Context, tx *ent.Tx, request *auth.CreateDepartmentRequest) error
		FindAllEmployee(ctx context.Context, req *auth.ListEmployeeRequest) ([]*ent.Employee, int, error)
		FindAllDepartment(ctx context.Context, req *auth.ListDepartmentRequest) ([]*ent.Department, int, error)
	}
)
