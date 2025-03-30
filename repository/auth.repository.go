package repository

import (
	"auth/ent"
	"auth/ent/auth"
	auth2 "auth/proto/auth"
	"auth/proto/gcommon"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)
import db "database/sql"

type authRepository struct {
	client *ent.Client
	db     *db.DB
}

func (a authRepository) CreateEmployee(ctx context.Context, tx *ent.Tx, request *auth2.CreateEmployeeRequest) error {
	defaultPassword := "123456"
	username := strings.Split(request.GetEmail(), "@")[0]

	authEmployee, err := tx.Auth.Create().
		SetEmail(request.GetEmail()).
		SetUsername(username).
		SetPassword(defaultPassword).
		SetState(int32(gcommon.AccountState_ACCOUNT_STATE_LOCKED)).
		Save(ctx)
	if err != nil {
		return err
	}

	prefix := "DR"
	if request.GetType() != auth2.EmployeeType_EMPLOYEE_TYPE_DOCTOR {
		prefix = "NU"
	}

	// Count existing entries
	var count int
	if request.GetType() == auth2.EmployeeType_EMPLOYEE_TYPE_DOCTOR {
		count, err = tx.Doctor.Query().Count(ctx)
	} else {
		count, err = tx.Nurse.Query().Count(ctx)
	}
	if err != nil {
		return err
	}

	// Generate the formatted code with 7-digit padding
	code := fmt.Sprintf("%s%07d", prefix, count+1)
	employee, err := tx.Employee.Create().
		SetEmployeeID(authEmployee.ID).
		SetFirstName(request.GetFirstName()).
		SetLastName(request.GetLastName()).
		SetGender(int32(request.GetGender())).
		SetDateOfBirth(time.UnixMilli(request.GetDateOfBirth())).
		SetAddress(request.GetAddress()).
		SetStartDate(time.UnixMilli(request.GetStartDate())).
		SetPhoneNumber(request.GetPhoneNumber()).
		SetDegreeName(request.GetDegreeName()).
		SetDegreeYear(request.GetDegreeYear()).
		SetDepartmentID(uuid.MustParse(request.GetDepartmentId())).
		SetEmployeeType(int32(request.GetType())).
		SetCode(code).
		Save(ctx)

	if err != nil {
		return err
	}

	if request.GetType() == auth2.EmployeeType_EMPLOYEE_TYPE_DOCTOR {
		_, err = tx.Doctor.Create().SetEmployee(employee).Save(ctx)
	} else {
		_, err = tx.Nurse.Create().SetEmployee(employee).Save(ctx)
	}
	return err
}

func (a authRepository) UpdateAccountState(ctx context.Context, tx *ent.Tx, username string, state int32) error {
	return tx.Auth.Update().Where(auth.UsernameEQ(username)).SetState(state).Exec(ctx)
}

func (a authRepository) FindByUsername(ctx context.Context, username string) (*ent.Auth, error) {
	return a.client.Auth.Query().Where(auth.UsernameEQ(username)).WithEmployee(
		func(query *ent.EmployeeQuery) {
			query.WithDoctor().WithNurse().WithDepartment()
		},
	).First(ctx)
}

func (a authRepository) UpdatePassword(ctx context.Context, tx *ent.Tx, username string, password string) error {
	return tx.Auth.Update().Where(auth.UsernameEQ(username)).SetPassword(password).Exec(ctx)
}

func (a authRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return a.client.Auth.Query().Where(auth.EmailEQ(email)).Exist(ctx)
}

func (a authRepository) CreateDepartment(ctx context.Context, tx *ent.Tx, request *auth2.CreateDepartmentRequest) error {
	_, err := tx.Department.Create().
		SetName(request.GetName()).
		Save(ctx)
	return err
}

func NewAuthRepository(client *ent.Client, db *db.DB) AuthRepository {
	return &authRepository{client: client, db: db}
}
