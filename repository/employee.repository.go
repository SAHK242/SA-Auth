package repository

import (
	"auth/ent"
	"auth/ent/auth"
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

type employeeRepository struct {
	client *ent.Client
	db     *db.DB
}

func (a employeeRepository) FindAllEmployee(ctx context.Context, req *auth2.ListEmployeeRequest) ([]*ent.Employee, int, error) {
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
		employeeIds := strings.Split(req.EmployeeId, ",")
		employeeUUIDs := make([]uuid.UUID, len(employeeIds))
		for i, id := range employeeIds {
			employeeUUIDs[i] = uuid.MustParse(id)
		}
		baseQuery.Where(employee.IDIn(employeeUUIDs...))
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

func (a employeeRepository) CreateEmployee(ctx context.Context, tx *ent.Tx, request *auth2.CreateEmployeeRequest) error {
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

func NewEmployeeRepository(client *ent.Client, db *db.DB) EmployeeRepository {
	return &employeeRepository{client: client, db: db}
}
