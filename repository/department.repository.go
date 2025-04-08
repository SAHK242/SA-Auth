package repository

import (
	"auth/ent"
	"auth/ent/department"
	"auth/ent/employee"
	auth2 "auth/proto/auth"
	"context"
)
import db "database/sql"

type departmentRepository struct {
	client *ent.Client
	db     *db.DB
}

func (a departmentRepository) FindAllDepartment(ctx context.Context, req *auth2.ListDepartmentRequest) ([]*ent.Department, int, error) {
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

func (a departmentRepository) CreateDepartment(ctx context.Context, tx *ent.Tx, request *auth2.CreateDepartmentRequest) error {
	_, err := tx.Department.Create().
		SetName(request.GetName()).
		Save(ctx)
	return err
}

func NewDepartmentRepository(client *ent.Client, db *db.DB) DepartmentRepository {
	return &departmentRepository{client: client, db: db}
}
