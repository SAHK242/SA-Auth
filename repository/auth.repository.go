package repository

import (
	"auth/ent"
	"auth/ent/auth"
	"auth/ent/department"
	"auth/ent/employee"
	auth2 "auth/proto/auth"
	"auth/proto/gcommon"
	"context"
	"entgo.io/ent/dialect/sql"
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

func (a authRepository) FindAllEmployee(ctx context.Context, req *auth2.ListEmployeeRequest) ([]*ent.Employee, int, error) {
	baseQuery := a.client.Employee.Query().WithAuth().WithDepartment()
	if req.Search != "" {
		baseQuery.Where(
			employee.Or(
				employee.CodeContainsFold(req.Search),
				employee.HasAuthWith(auth.EmailContainsFold(req.Search)),
				employee.PhoneNumberContainsFold(req.Search),
				func(selector *sql.Selector) {
					enNamePredicate := sql.ExprP("concat(first_name,' ',last_name) ilike " + "'%" + req.Search + "%'")
					viNamePredicate := sql.ExprP("concat(last_name,' ',first_name) ilike " + "'%" + req.Search + "%'")
					selector.Where(sql.Or(enNamePredicate, viNamePredicate))
				},
			),
		)
	}

	if req.DepartmentId != "" {
		baseQuery.Where(employee.DepartmentIDEQ(uuid.MustParse(req.DepartmentId)))
	}

	if req.EmployeeId != "" {
		baseQuery.Where(employee.IDEQ(uuid.MustParse(req.EmployeeId)))
	}

	if req.EmployeeType != 0 {
		baseQuery.Where(employee.EmployeeTypeEQ(int32(req.EmployeeType)))
	}

	pagination := ToPagination(req.Pageable)

	var records []*ent.Employee
	var total int
	var err error

	if pagination.PagingIgnored {
		records, err = baseQuery.Order(ToOrderSpecifier(pagination.Sort, employee.FieldFirstName, Ascending)).All(ctx)

		if err != nil {
			return []*ent.Employee{}, 0, err
		}

		return records, len(records), err
	}

	total, err = baseQuery.Clone().Count(ctx)

	if err != nil {
		return []*ent.Employee{}, 0, err
	}

	records, err = baseQuery.Clone().
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(ToOrderSpecifier(pagination.Sort, employee.FieldFirstName, Ascending)).
		All(ctx)

	if err != nil {
		return []*ent.Employee{}, 0, err
	}
	return records, total, nil
}

func (a authRepository) FindAllDepartment(ctx context.Context, req *auth2.ListDepartmentRequest) ([]*ent.Department, int, error) {
	baseQuery := a.client.Department.Query()
	if req.Search != "" {
		baseQuery.Where(department.NameContainsFold(req.Search))
	}

	pagination := ToPagination(req.Pageable)

	var records []*ent.Department
	var total int
	var err error

	if pagination.PagingIgnored {
		records, err = baseQuery.Order(ToOrderSpecifier(pagination.Sort, employee.FieldFirstName, Ascending)).All(ctx)

		if err != nil {
			return []*ent.Department{}, 0, err
		}

		return records, len(records), err
	}

	total, err = baseQuery.Clone().Count(ctx)

	if err != nil {
		return []*ent.Department{}, 0, err
	}

	records, err = baseQuery.Clone().
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(ToOrderSpecifier(pagination.Sort, employee.FieldFirstName, Ascending)).
		All(ctx)

	if err != nil {
		return []*ent.Department{}, 0, err
	}
	return records, total, nil
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
	employeeSave, err := tx.Employee.Create().
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
		_, err = tx.Doctor.Create().SetEmployee(employeeSave).Save(ctx)
	} else {
		_, err = tx.Nurse.Create().SetEmployee(employeeSave).Save(ctx)
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
