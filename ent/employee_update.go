// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/auth"
	"auth/ent/department"
	"auth/ent/doctor"
	"auth/ent/employee"
	"auth/ent/nurse"
	"auth/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetEmployeeID sets the "employee_id" field.
func (eu *EmployeeUpdate) SetEmployeeID(u uuid.UUID) *EmployeeUpdate {
	eu.mutation.SetEmployeeID(u)
	return eu
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableEmployeeID(u *uuid.UUID) *EmployeeUpdate {
	if u != nil {
		eu.SetEmployeeID(*u)
	}
	return eu
}

// SetFirstName sets the "first_name" field.
func (eu *EmployeeUpdate) SetFirstName(s string) *EmployeeUpdate {
	eu.mutation.SetFirstName(s)
	return eu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableFirstName(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetFirstName(*s)
	}
	return eu
}

// SetLastName sets the "last_name" field.
func (eu *EmployeeUpdate) SetLastName(s string) *EmployeeUpdate {
	eu.mutation.SetLastName(s)
	return eu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableLastName(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetLastName(*s)
	}
	return eu
}

// SetGender sets the "gender" field.
func (eu *EmployeeUpdate) SetGender(i int32) *EmployeeUpdate {
	eu.mutation.ResetGender()
	eu.mutation.SetGender(i)
	return eu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableGender(i *int32) *EmployeeUpdate {
	if i != nil {
		eu.SetGender(*i)
	}
	return eu
}

// AddGender adds i to the "gender" field.
func (eu *EmployeeUpdate) AddGender(i int32) *EmployeeUpdate {
	eu.mutation.AddGender(i)
	return eu
}

// SetDateOfBirth sets the "date_of_birth" field.
func (eu *EmployeeUpdate) SetDateOfBirth(t time.Time) *EmployeeUpdate {
	eu.mutation.SetDateOfBirth(t)
	return eu
}

// SetNillableDateOfBirth sets the "date_of_birth" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDateOfBirth(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetDateOfBirth(*t)
	}
	return eu
}

// SetCode sets the "code" field.
func (eu *EmployeeUpdate) SetCode(s string) *EmployeeUpdate {
	eu.mutation.SetCode(s)
	return eu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableCode(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetCode(*s)
	}
	return eu
}

// SetAddress sets the "address" field.
func (eu *EmployeeUpdate) SetAddress(s string) *EmployeeUpdate {
	eu.mutation.SetAddress(s)
	return eu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableAddress(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetAddress(*s)
	}
	return eu
}

// SetStartDate sets the "start_date" field.
func (eu *EmployeeUpdate) SetStartDate(t time.Time) *EmployeeUpdate {
	eu.mutation.SetStartDate(t)
	return eu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableStartDate(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetStartDate(*t)
	}
	return eu
}

// SetEndDate sets the "end_date" field.
func (eu *EmployeeUpdate) SetEndDate(t time.Time) *EmployeeUpdate {
	eu.mutation.SetEndDate(t)
	return eu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableEndDate(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetEndDate(*t)
	}
	return eu
}

// ClearEndDate clears the value of the "end_date" field.
func (eu *EmployeeUpdate) ClearEndDate() *EmployeeUpdate {
	eu.mutation.ClearEndDate()
	return eu
}

// SetPhoneNumber sets the "phone_number" field.
func (eu *EmployeeUpdate) SetPhoneNumber(s string) *EmployeeUpdate {
	eu.mutation.SetPhoneNumber(s)
	return eu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillablePhoneNumber(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetPhoneNumber(*s)
	}
	return eu
}

// SetDegreeName sets the "degree_name" field.
func (eu *EmployeeUpdate) SetDegreeName(s string) *EmployeeUpdate {
	eu.mutation.SetDegreeName(s)
	return eu
}

// SetNillableDegreeName sets the "degree_name" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDegreeName(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetDegreeName(*s)
	}
	return eu
}

// SetDegreeYear sets the "degree_year" field.
func (eu *EmployeeUpdate) SetDegreeYear(i int32) *EmployeeUpdate {
	eu.mutation.ResetDegreeYear()
	eu.mutation.SetDegreeYear(i)
	return eu
}

// SetNillableDegreeYear sets the "degree_year" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDegreeYear(i *int32) *EmployeeUpdate {
	if i != nil {
		eu.SetDegreeYear(*i)
	}
	return eu
}

// AddDegreeYear adds i to the "degree_year" field.
func (eu *EmployeeUpdate) AddDegreeYear(i int32) *EmployeeUpdate {
	eu.mutation.AddDegreeYear(i)
	return eu
}

// SetDepartmentID sets the "department_id" field.
func (eu *EmployeeUpdate) SetDepartmentID(u uuid.UUID) *EmployeeUpdate {
	eu.mutation.SetDepartmentID(u)
	return eu
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDepartmentID(u *uuid.UUID) *EmployeeUpdate {
	if u != nil {
		eu.SetDepartmentID(*u)
	}
	return eu
}

// SetEmployeeType sets the "employee_type" field.
func (eu *EmployeeUpdate) SetEmployeeType(i int32) *EmployeeUpdate {
	eu.mutation.ResetEmployeeType()
	eu.mutation.SetEmployeeType(i)
	return eu
}

// SetNillableEmployeeType sets the "employee_type" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableEmployeeType(i *int32) *EmployeeUpdate {
	if i != nil {
		eu.SetEmployeeType(*i)
	}
	return eu
}

// AddEmployeeType adds i to the "employee_type" field.
func (eu *EmployeeUpdate) AddEmployeeType(i int32) *EmployeeUpdate {
	eu.mutation.AddEmployeeType(i)
	return eu
}

// SetAuthID sets the "auth" edge to the Auth entity by ID.
func (eu *EmployeeUpdate) SetAuthID(id uuid.UUID) *EmployeeUpdate {
	eu.mutation.SetAuthID(id)
	return eu
}

// SetAuth sets the "auth" edge to the Auth entity.
func (eu *EmployeeUpdate) SetAuth(a *Auth) *EmployeeUpdate {
	return eu.SetAuthID(a.ID)
}

// SetDepartment sets the "department" edge to the Department entity.
func (eu *EmployeeUpdate) SetDepartment(d *Department) *EmployeeUpdate {
	return eu.SetDepartmentID(d.ID)
}

// SetDoctorID sets the "doctor" edge to the Doctor entity by ID.
func (eu *EmployeeUpdate) SetDoctorID(id uuid.UUID) *EmployeeUpdate {
	eu.mutation.SetDoctorID(id)
	return eu
}

// SetNillableDoctorID sets the "doctor" edge to the Doctor entity by ID if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDoctorID(id *uuid.UUID) *EmployeeUpdate {
	if id != nil {
		eu = eu.SetDoctorID(*id)
	}
	return eu
}

// SetDoctor sets the "doctor" edge to the Doctor entity.
func (eu *EmployeeUpdate) SetDoctor(d *Doctor) *EmployeeUpdate {
	return eu.SetDoctorID(d.ID)
}

// SetNurseID sets the "nurse" edge to the Nurse entity by ID.
func (eu *EmployeeUpdate) SetNurseID(id uuid.UUID) *EmployeeUpdate {
	eu.mutation.SetNurseID(id)
	return eu
}

// SetNillableNurseID sets the "nurse" edge to the Nurse entity by ID if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableNurseID(id *uuid.UUID) *EmployeeUpdate {
	if id != nil {
		eu = eu.SetNurseID(*id)
	}
	return eu
}

// SetNurse sets the "nurse" edge to the Nurse entity.
func (eu *EmployeeUpdate) SetNurse(n *Nurse) *EmployeeUpdate {
	return eu.SetNurseID(n.ID)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// ClearAuth clears the "auth" edge to the Auth entity.
func (eu *EmployeeUpdate) ClearAuth() *EmployeeUpdate {
	eu.mutation.ClearAuth()
	return eu
}

// ClearDepartment clears the "department" edge to the Department entity.
func (eu *EmployeeUpdate) ClearDepartment() *EmployeeUpdate {
	eu.mutation.ClearDepartment()
	return eu
}

// ClearDoctor clears the "doctor" edge to the Doctor entity.
func (eu *EmployeeUpdate) ClearDoctor() *EmployeeUpdate {
	eu.mutation.ClearDoctor()
	return eu
}

// ClearNurse clears the "nurse" edge to the Nurse entity.
func (eu *EmployeeUpdate) ClearNurse() *EmployeeUpdate {
	eu.mutation.ClearNurse()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmployeeUpdate) check() error {
	if v, ok := eu.mutation.Code(); ok {
		if err := employee.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Employee.code": %w`, err)}
		}
	}
	if eu.mutation.AuthCleared() && len(eu.mutation.AuthIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Employee.auth"`)
	}
	if eu.mutation.DepartmentCleared() && len(eu.mutation.DepartmentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Employee.department"`)
	}
	return nil
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := eu.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeInt32, value)
	}
	if value, ok := eu.mutation.AddedGender(); ok {
		_spec.AddField(employee.FieldGender, field.TypeInt32, value)
	}
	if value, ok := eu.mutation.DateOfBirth(); ok {
		_spec.SetField(employee.FieldDateOfBirth, field.TypeTime, value)
	}
	if value, ok := eu.mutation.Code(); ok {
		_spec.SetField(employee.FieldCode, field.TypeString, value)
	}
	if value, ok := eu.mutation.Address(); ok {
		_spec.SetField(employee.FieldAddress, field.TypeString, value)
	}
	if value, ok := eu.mutation.StartDate(); ok {
		_spec.SetField(employee.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := eu.mutation.EndDate(); ok {
		_spec.SetField(employee.FieldEndDate, field.TypeTime, value)
	}
	if eu.mutation.EndDateCleared() {
		_spec.ClearField(employee.FieldEndDate, field.TypeTime)
	}
	if value, ok := eu.mutation.PhoneNumber(); ok {
		_spec.SetField(employee.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := eu.mutation.DegreeName(); ok {
		_spec.SetField(employee.FieldDegreeName, field.TypeString, value)
	}
	if value, ok := eu.mutation.DegreeYear(); ok {
		_spec.SetField(employee.FieldDegreeYear, field.TypeInt32, value)
	}
	if value, ok := eu.mutation.AddedDegreeYear(); ok {
		_spec.AddField(employee.FieldDegreeYear, field.TypeInt32, value)
	}
	if value, ok := eu.mutation.EmployeeType(); ok {
		_spec.SetField(employee.FieldEmployeeType, field.TypeInt32, value)
	}
	if value, ok := eu.mutation.AddedEmployeeType(); ok {
		_spec.AddField(employee.FieldEmployeeType, field.TypeInt32, value)
	}
	if eu.mutation.AuthCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.AuthTable,
			Columns: []string{employee.AuthColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AuthIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.AuthTable,
			Columns: []string{employee.AuthColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.DepartmentTable,
			Columns: []string{employee.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.DepartmentTable,
			Columns: []string{employee.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.DoctorTable,
			Columns: []string{employee.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.DoctorTable,
			Columns: []string{employee.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.NurseTable,
			Columns: []string{employee.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.NurseTable,
			Columns: []string{employee.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetEmployeeID sets the "employee_id" field.
func (euo *EmployeeUpdateOne) SetEmployeeID(u uuid.UUID) *EmployeeUpdateOne {
	euo.mutation.SetEmployeeID(u)
	return euo
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableEmployeeID(u *uuid.UUID) *EmployeeUpdateOne {
	if u != nil {
		euo.SetEmployeeID(*u)
	}
	return euo
}

// SetFirstName sets the "first_name" field.
func (euo *EmployeeUpdateOne) SetFirstName(s string) *EmployeeUpdateOne {
	euo.mutation.SetFirstName(s)
	return euo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableFirstName(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetFirstName(*s)
	}
	return euo
}

// SetLastName sets the "last_name" field.
func (euo *EmployeeUpdateOne) SetLastName(s string) *EmployeeUpdateOne {
	euo.mutation.SetLastName(s)
	return euo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableLastName(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetLastName(*s)
	}
	return euo
}

// SetGender sets the "gender" field.
func (euo *EmployeeUpdateOne) SetGender(i int32) *EmployeeUpdateOne {
	euo.mutation.ResetGender()
	euo.mutation.SetGender(i)
	return euo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableGender(i *int32) *EmployeeUpdateOne {
	if i != nil {
		euo.SetGender(*i)
	}
	return euo
}

// AddGender adds i to the "gender" field.
func (euo *EmployeeUpdateOne) AddGender(i int32) *EmployeeUpdateOne {
	euo.mutation.AddGender(i)
	return euo
}

// SetDateOfBirth sets the "date_of_birth" field.
func (euo *EmployeeUpdateOne) SetDateOfBirth(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetDateOfBirth(t)
	return euo
}

// SetNillableDateOfBirth sets the "date_of_birth" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDateOfBirth(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetDateOfBirth(*t)
	}
	return euo
}

// SetCode sets the "code" field.
func (euo *EmployeeUpdateOne) SetCode(s string) *EmployeeUpdateOne {
	euo.mutation.SetCode(s)
	return euo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableCode(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetCode(*s)
	}
	return euo
}

// SetAddress sets the "address" field.
func (euo *EmployeeUpdateOne) SetAddress(s string) *EmployeeUpdateOne {
	euo.mutation.SetAddress(s)
	return euo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableAddress(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetAddress(*s)
	}
	return euo
}

// SetStartDate sets the "start_date" field.
func (euo *EmployeeUpdateOne) SetStartDate(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetStartDate(t)
	return euo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableStartDate(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetStartDate(*t)
	}
	return euo
}

// SetEndDate sets the "end_date" field.
func (euo *EmployeeUpdateOne) SetEndDate(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetEndDate(t)
	return euo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableEndDate(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetEndDate(*t)
	}
	return euo
}

// ClearEndDate clears the value of the "end_date" field.
func (euo *EmployeeUpdateOne) ClearEndDate() *EmployeeUpdateOne {
	euo.mutation.ClearEndDate()
	return euo
}

// SetPhoneNumber sets the "phone_number" field.
func (euo *EmployeeUpdateOne) SetPhoneNumber(s string) *EmployeeUpdateOne {
	euo.mutation.SetPhoneNumber(s)
	return euo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillablePhoneNumber(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetPhoneNumber(*s)
	}
	return euo
}

// SetDegreeName sets the "degree_name" field.
func (euo *EmployeeUpdateOne) SetDegreeName(s string) *EmployeeUpdateOne {
	euo.mutation.SetDegreeName(s)
	return euo
}

// SetNillableDegreeName sets the "degree_name" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDegreeName(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetDegreeName(*s)
	}
	return euo
}

// SetDegreeYear sets the "degree_year" field.
func (euo *EmployeeUpdateOne) SetDegreeYear(i int32) *EmployeeUpdateOne {
	euo.mutation.ResetDegreeYear()
	euo.mutation.SetDegreeYear(i)
	return euo
}

// SetNillableDegreeYear sets the "degree_year" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDegreeYear(i *int32) *EmployeeUpdateOne {
	if i != nil {
		euo.SetDegreeYear(*i)
	}
	return euo
}

// AddDegreeYear adds i to the "degree_year" field.
func (euo *EmployeeUpdateOne) AddDegreeYear(i int32) *EmployeeUpdateOne {
	euo.mutation.AddDegreeYear(i)
	return euo
}

// SetDepartmentID sets the "department_id" field.
func (euo *EmployeeUpdateOne) SetDepartmentID(u uuid.UUID) *EmployeeUpdateOne {
	euo.mutation.SetDepartmentID(u)
	return euo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDepartmentID(u *uuid.UUID) *EmployeeUpdateOne {
	if u != nil {
		euo.SetDepartmentID(*u)
	}
	return euo
}

// SetEmployeeType sets the "employee_type" field.
func (euo *EmployeeUpdateOne) SetEmployeeType(i int32) *EmployeeUpdateOne {
	euo.mutation.ResetEmployeeType()
	euo.mutation.SetEmployeeType(i)
	return euo
}

// SetNillableEmployeeType sets the "employee_type" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableEmployeeType(i *int32) *EmployeeUpdateOne {
	if i != nil {
		euo.SetEmployeeType(*i)
	}
	return euo
}

// AddEmployeeType adds i to the "employee_type" field.
func (euo *EmployeeUpdateOne) AddEmployeeType(i int32) *EmployeeUpdateOne {
	euo.mutation.AddEmployeeType(i)
	return euo
}

// SetAuthID sets the "auth" edge to the Auth entity by ID.
func (euo *EmployeeUpdateOne) SetAuthID(id uuid.UUID) *EmployeeUpdateOne {
	euo.mutation.SetAuthID(id)
	return euo
}

// SetAuth sets the "auth" edge to the Auth entity.
func (euo *EmployeeUpdateOne) SetAuth(a *Auth) *EmployeeUpdateOne {
	return euo.SetAuthID(a.ID)
}

// SetDepartment sets the "department" edge to the Department entity.
func (euo *EmployeeUpdateOne) SetDepartment(d *Department) *EmployeeUpdateOne {
	return euo.SetDepartmentID(d.ID)
}

// SetDoctorID sets the "doctor" edge to the Doctor entity by ID.
func (euo *EmployeeUpdateOne) SetDoctorID(id uuid.UUID) *EmployeeUpdateOne {
	euo.mutation.SetDoctorID(id)
	return euo
}

// SetNillableDoctorID sets the "doctor" edge to the Doctor entity by ID if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDoctorID(id *uuid.UUID) *EmployeeUpdateOne {
	if id != nil {
		euo = euo.SetDoctorID(*id)
	}
	return euo
}

// SetDoctor sets the "doctor" edge to the Doctor entity.
func (euo *EmployeeUpdateOne) SetDoctor(d *Doctor) *EmployeeUpdateOne {
	return euo.SetDoctorID(d.ID)
}

// SetNurseID sets the "nurse" edge to the Nurse entity by ID.
func (euo *EmployeeUpdateOne) SetNurseID(id uuid.UUID) *EmployeeUpdateOne {
	euo.mutation.SetNurseID(id)
	return euo
}

// SetNillableNurseID sets the "nurse" edge to the Nurse entity by ID if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableNurseID(id *uuid.UUID) *EmployeeUpdateOne {
	if id != nil {
		euo = euo.SetNurseID(*id)
	}
	return euo
}

// SetNurse sets the "nurse" edge to the Nurse entity.
func (euo *EmployeeUpdateOne) SetNurse(n *Nurse) *EmployeeUpdateOne {
	return euo.SetNurseID(n.ID)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// ClearAuth clears the "auth" edge to the Auth entity.
func (euo *EmployeeUpdateOne) ClearAuth() *EmployeeUpdateOne {
	euo.mutation.ClearAuth()
	return euo
}

// ClearDepartment clears the "department" edge to the Department entity.
func (euo *EmployeeUpdateOne) ClearDepartment() *EmployeeUpdateOne {
	euo.mutation.ClearDepartment()
	return euo
}

// ClearDoctor clears the "doctor" edge to the Doctor entity.
func (euo *EmployeeUpdateOne) ClearDoctor() *EmployeeUpdateOne {
	euo.mutation.ClearDoctor()
	return euo
}

// ClearNurse clears the "nurse" edge to the Nurse entity.
func (euo *EmployeeUpdateOne) ClearNurse() *EmployeeUpdateOne {
	euo.mutation.ClearNurse()
	return euo
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmployeeUpdateOne) check() error {
	if v, ok := euo.mutation.Code(); ok {
		if err := employee.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Employee.code": %w`, err)}
		}
	}
	if euo.mutation.AuthCleared() && len(euo.mutation.AuthIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Employee.auth"`)
	}
	if euo.mutation.DepartmentCleared() && len(euo.mutation.DepartmentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Employee.department"`)
	}
	return nil
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := euo.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeInt32, value)
	}
	if value, ok := euo.mutation.AddedGender(); ok {
		_spec.AddField(employee.FieldGender, field.TypeInt32, value)
	}
	if value, ok := euo.mutation.DateOfBirth(); ok {
		_spec.SetField(employee.FieldDateOfBirth, field.TypeTime, value)
	}
	if value, ok := euo.mutation.Code(); ok {
		_spec.SetField(employee.FieldCode, field.TypeString, value)
	}
	if value, ok := euo.mutation.Address(); ok {
		_spec.SetField(employee.FieldAddress, field.TypeString, value)
	}
	if value, ok := euo.mutation.StartDate(); ok {
		_spec.SetField(employee.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := euo.mutation.EndDate(); ok {
		_spec.SetField(employee.FieldEndDate, field.TypeTime, value)
	}
	if euo.mutation.EndDateCleared() {
		_spec.ClearField(employee.FieldEndDate, field.TypeTime)
	}
	if value, ok := euo.mutation.PhoneNumber(); ok {
		_spec.SetField(employee.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := euo.mutation.DegreeName(); ok {
		_spec.SetField(employee.FieldDegreeName, field.TypeString, value)
	}
	if value, ok := euo.mutation.DegreeYear(); ok {
		_spec.SetField(employee.FieldDegreeYear, field.TypeInt32, value)
	}
	if value, ok := euo.mutation.AddedDegreeYear(); ok {
		_spec.AddField(employee.FieldDegreeYear, field.TypeInt32, value)
	}
	if value, ok := euo.mutation.EmployeeType(); ok {
		_spec.SetField(employee.FieldEmployeeType, field.TypeInt32, value)
	}
	if value, ok := euo.mutation.AddedEmployeeType(); ok {
		_spec.AddField(employee.FieldEmployeeType, field.TypeInt32, value)
	}
	if euo.mutation.AuthCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.AuthTable,
			Columns: []string{employee.AuthColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AuthIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.AuthTable,
			Columns: []string{employee.AuthColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.DepartmentTable,
			Columns: []string{employee.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.DepartmentTable,
			Columns: []string{employee.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.DoctorTable,
			Columns: []string{employee.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.DoctorTable,
			Columns: []string{employee.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.NurseTable,
			Columns: []string{employee.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   employee.NurseTable,
			Columns: []string{employee.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
