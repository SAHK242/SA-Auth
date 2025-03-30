// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/employee"
	"auth/ent/nurse"
	"auth/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NurseUpdate is the builder for updating Nurse entities.
type NurseUpdate struct {
	config
	hooks    []Hook
	mutation *NurseMutation
}

// Where appends a list predicates to the NurseUpdate builder.
func (nu *NurseUpdate) Where(ps ...predicate.Nurse) *NurseUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetNurseID sets the "nurse_id" field.
func (nu *NurseUpdate) SetNurseID(u uuid.UUID) *NurseUpdate {
	nu.mutation.SetNurseID(u)
	return nu
}

// SetNillableNurseID sets the "nurse_id" field if the given value is not nil.
func (nu *NurseUpdate) SetNillableNurseID(u *uuid.UUID) *NurseUpdate {
	if u != nil {
		nu.SetNurseID(*u)
	}
	return nu
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (nu *NurseUpdate) SetEmployeeID(id uuid.UUID) *NurseUpdate {
	nu.mutation.SetEmployeeID(id)
	return nu
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (nu *NurseUpdate) SetEmployee(e *Employee) *NurseUpdate {
	return nu.SetEmployeeID(e.ID)
}

// Mutation returns the NurseMutation object of the builder.
func (nu *NurseUpdate) Mutation() *NurseMutation {
	return nu.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (nu *NurseUpdate) ClearEmployee() *NurseUpdate {
	nu.mutation.ClearEmployee()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NurseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NurseUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NurseUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NurseUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NurseUpdate) check() error {
	if nu.mutation.EmployeeCleared() && len(nu.mutation.EmployeeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Nurse.employee"`)
	}
	return nil
}

func (nu *NurseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(nurse.Table, nurse.Columns, sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nu.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EmployeeTable,
			Columns: []string{nurse.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EmployeeTable,
			Columns: []string{nurse.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NurseUpdateOne is the builder for updating a single Nurse entity.
type NurseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NurseMutation
}

// SetNurseID sets the "nurse_id" field.
func (nuo *NurseUpdateOne) SetNurseID(u uuid.UUID) *NurseUpdateOne {
	nuo.mutation.SetNurseID(u)
	return nuo
}

// SetNillableNurseID sets the "nurse_id" field if the given value is not nil.
func (nuo *NurseUpdateOne) SetNillableNurseID(u *uuid.UUID) *NurseUpdateOne {
	if u != nil {
		nuo.SetNurseID(*u)
	}
	return nuo
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (nuo *NurseUpdateOne) SetEmployeeID(id uuid.UUID) *NurseUpdateOne {
	nuo.mutation.SetEmployeeID(id)
	return nuo
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (nuo *NurseUpdateOne) SetEmployee(e *Employee) *NurseUpdateOne {
	return nuo.SetEmployeeID(e.ID)
}

// Mutation returns the NurseMutation object of the builder.
func (nuo *NurseUpdateOne) Mutation() *NurseMutation {
	return nuo.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (nuo *NurseUpdateOne) ClearEmployee() *NurseUpdateOne {
	nuo.mutation.ClearEmployee()
	return nuo
}

// Where appends a list predicates to the NurseUpdate builder.
func (nuo *NurseUpdateOne) Where(ps ...predicate.Nurse) *NurseUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NurseUpdateOne) Select(field string, fields ...string) *NurseUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Nurse entity.
func (nuo *NurseUpdateOne) Save(ctx context.Context) (*Nurse, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NurseUpdateOne) SaveX(ctx context.Context) *Nurse {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NurseUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NurseUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NurseUpdateOne) check() error {
	if nuo.mutation.EmployeeCleared() && len(nuo.mutation.EmployeeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Nurse.employee"`)
	}
	return nil
}

func (nuo *NurseUpdateOne) sqlSave(ctx context.Context) (_node *Nurse, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(nurse.Table, nurse.Columns, sqlgraph.NewFieldSpec(nurse.FieldID, field.TypeUUID))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Nurse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nurse.FieldID)
		for _, f := range fields {
			if !nurse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nurse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nuo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EmployeeTable,
			Columns: []string{nurse.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nurse.EmployeeTable,
			Columns: []string{nurse.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Nurse{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
