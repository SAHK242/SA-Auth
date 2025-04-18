// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/auth"
	"auth/ent/employee"
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

// AuthUpdate is the builder for updating Auth entities.
type AuthUpdate struct {
	config
	hooks    []Hook
	mutation *AuthMutation
}

// Where appends a list predicates to the AuthUpdate builder.
func (au *AuthUpdate) Where(ps ...predicate.Auth) *AuthUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetEmail sets the "email" field.
func (au *AuthUpdate) SetEmail(s string) *AuthUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (au *AuthUpdate) SetNillableEmail(s *string) *AuthUpdate {
	if s != nil {
		au.SetEmail(*s)
	}
	return au
}

// SetUsername sets the "username" field.
func (au *AuthUpdate) SetUsername(s string) *AuthUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AuthUpdate) SetNillableUsername(s *string) *AuthUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// SetPassword sets the "password" field.
func (au *AuthUpdate) SetPassword(s string) *AuthUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AuthUpdate) SetNillablePassword(s *string) *AuthUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetState sets the "state" field.
func (au *AuthUpdate) SetState(i int32) *AuthUpdate {
	au.mutation.ResetState()
	au.mutation.SetState(i)
	return au
}

// SetNillableState sets the "state" field if the given value is not nil.
func (au *AuthUpdate) SetNillableState(i *int32) *AuthUpdate {
	if i != nil {
		au.SetState(*i)
	}
	return au
}

// AddState adds i to the "state" field.
func (au *AuthUpdate) AddState(i int32) *AuthUpdate {
	au.mutation.AddState(i)
	return au
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (au *AuthUpdate) SetLastModifiedDate(t time.Time) *AuthUpdate {
	au.mutation.SetLastModifiedDate(t)
	return au
}

// SetNillableLastModifiedDate sets the "last_modified_date" field if the given value is not nil.
func (au *AuthUpdate) SetNillableLastModifiedDate(t *time.Time) *AuthUpdate {
	if t != nil {
		au.SetLastModifiedDate(*t)
	}
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *AuthUpdate) SetCreatedBy(u uuid.UUID) *AuthUpdate {
	au.mutation.SetCreatedBy(u)
	return au
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (au *AuthUpdate) SetNillableCreatedBy(u *uuid.UUID) *AuthUpdate {
	if u != nil {
		au.SetCreatedBy(*u)
	}
	return au
}

// ClearCreatedBy clears the value of the "created_by" field.
func (au *AuthUpdate) ClearCreatedBy() *AuthUpdate {
	au.mutation.ClearCreatedBy()
	return au
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (au *AuthUpdate) SetLastModifiedBy(u uuid.UUID) *AuthUpdate {
	au.mutation.SetLastModifiedBy(u)
	return au
}

// SetNillableLastModifiedBy sets the "last_modified_by" field if the given value is not nil.
func (au *AuthUpdate) SetNillableLastModifiedBy(u *uuid.UUID) *AuthUpdate {
	if u != nil {
		au.SetLastModifiedBy(*u)
	}
	return au
}

// ClearLastModifiedBy clears the value of the "last_modified_by" field.
func (au *AuthUpdate) ClearLastModifiedBy() *AuthUpdate {
	au.mutation.ClearLastModifiedBy()
	return au
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (au *AuthUpdate) SetEmployeeID(id uuid.UUID) *AuthUpdate {
	au.mutation.SetEmployeeID(id)
	return au
}

// SetNillableEmployeeID sets the "employee" edge to the Employee entity by ID if the given value is not nil.
func (au *AuthUpdate) SetNillableEmployeeID(id *uuid.UUID) *AuthUpdate {
	if id != nil {
		au = au.SetEmployeeID(*id)
	}
	return au
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (au *AuthUpdate) SetEmployee(e *Employee) *AuthUpdate {
	return au.SetEmployeeID(e.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (au *AuthUpdate) Mutation() *AuthMutation {
	return au.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (au *AuthUpdate) ClearEmployee() *AuthUpdate {
	au.mutation.ClearEmployee()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(auth.Table, auth.Columns, sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(auth.FieldEmail, field.TypeString, value)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(auth.FieldUsername, field.TypeString, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(auth.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.State(); ok {
		_spec.SetField(auth.FieldState, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedState(); ok {
		_spec.AddField(auth.FieldState, field.TypeInt32, value)
	}
	if value, ok := au.mutation.LastModifiedDate(); ok {
		_spec.SetField(auth.FieldLastModifiedDate, field.TypeTime, value)
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.SetField(auth.FieldCreatedBy, field.TypeUUID, value)
	}
	if au.mutation.CreatedByCleared() {
		_spec.ClearField(auth.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := au.mutation.LastModifiedBy(); ok {
		_spec.SetField(auth.FieldLastModifiedBy, field.TypeUUID, value)
	}
	if au.mutation.LastModifiedByCleared() {
		_spec.ClearField(auth.FieldLastModifiedBy, field.TypeUUID)
	}
	if au.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   auth.EmployeeTable,
			Columns: []string{auth.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   auth.EmployeeTable,
			Columns: []string{auth.EmployeeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{auth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuthUpdateOne is the builder for updating a single Auth entity.
type AuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthMutation
}

// SetEmail sets the "email" field.
func (auo *AuthUpdateOne) SetEmail(s string) *AuthUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableEmail(s *string) *AuthUpdateOne {
	if s != nil {
		auo.SetEmail(*s)
	}
	return auo
}

// SetUsername sets the "username" field.
func (auo *AuthUpdateOne) SetUsername(s string) *AuthUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableUsername(s *string) *AuthUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// SetPassword sets the "password" field.
func (auo *AuthUpdateOne) SetPassword(s string) *AuthUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillablePassword(s *string) *AuthUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetState sets the "state" field.
func (auo *AuthUpdateOne) SetState(i int32) *AuthUpdateOne {
	auo.mutation.ResetState()
	auo.mutation.SetState(i)
	return auo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableState(i *int32) *AuthUpdateOne {
	if i != nil {
		auo.SetState(*i)
	}
	return auo
}

// AddState adds i to the "state" field.
func (auo *AuthUpdateOne) AddState(i int32) *AuthUpdateOne {
	auo.mutation.AddState(i)
	return auo
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (auo *AuthUpdateOne) SetLastModifiedDate(t time.Time) *AuthUpdateOne {
	auo.mutation.SetLastModifiedDate(t)
	return auo
}

// SetNillableLastModifiedDate sets the "last_modified_date" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableLastModifiedDate(t *time.Time) *AuthUpdateOne {
	if t != nil {
		auo.SetLastModifiedDate(*t)
	}
	return auo
}

// SetCreatedBy sets the "created_by" field.
func (auo *AuthUpdateOne) SetCreatedBy(u uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetCreatedBy(u)
	return auo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *AuthUpdateOne {
	if u != nil {
		auo.SetCreatedBy(*u)
	}
	return auo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (auo *AuthUpdateOne) ClearCreatedBy() *AuthUpdateOne {
	auo.mutation.ClearCreatedBy()
	return auo
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (auo *AuthUpdateOne) SetLastModifiedBy(u uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetLastModifiedBy(u)
	return auo
}

// SetNillableLastModifiedBy sets the "last_modified_by" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableLastModifiedBy(u *uuid.UUID) *AuthUpdateOne {
	if u != nil {
		auo.SetLastModifiedBy(*u)
	}
	return auo
}

// ClearLastModifiedBy clears the value of the "last_modified_by" field.
func (auo *AuthUpdateOne) ClearLastModifiedBy() *AuthUpdateOne {
	auo.mutation.ClearLastModifiedBy()
	return auo
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (auo *AuthUpdateOne) SetEmployeeID(id uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetEmployeeID(id)
	return auo
}

// SetNillableEmployeeID sets the "employee" edge to the Employee entity by ID if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableEmployeeID(id *uuid.UUID) *AuthUpdateOne {
	if id != nil {
		auo = auo.SetEmployeeID(*id)
	}
	return auo
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (auo *AuthUpdateOne) SetEmployee(e *Employee) *AuthUpdateOne {
	return auo.SetEmployeeID(e.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (auo *AuthUpdateOne) Mutation() *AuthMutation {
	return auo.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (auo *AuthUpdateOne) ClearEmployee() *AuthUpdateOne {
	auo.mutation.ClearEmployee()
	return auo
}

// Where appends a list predicates to the AuthUpdate builder.
func (auo *AuthUpdateOne) Where(ps ...predicate.Auth) *AuthUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthUpdateOne) Select(field string, fields ...string) *AuthUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Auth entity.
func (auo *AuthUpdateOne) Save(ctx context.Context) (*Auth, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthUpdateOne) SaveX(ctx context.Context) *Auth {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AuthUpdateOne) sqlSave(ctx context.Context) (_node *Auth, err error) {
	_spec := sqlgraph.NewUpdateSpec(auth.Table, auth.Columns, sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Auth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, auth.FieldID)
		for _, f := range fields {
			if !auth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != auth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(auth.FieldEmail, field.TypeString, value)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(auth.FieldUsername, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(auth.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.State(); ok {
		_spec.SetField(auth.FieldState, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedState(); ok {
		_spec.AddField(auth.FieldState, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.LastModifiedDate(); ok {
		_spec.SetField(auth.FieldLastModifiedDate, field.TypeTime, value)
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.SetField(auth.FieldCreatedBy, field.TypeUUID, value)
	}
	if auo.mutation.CreatedByCleared() {
		_spec.ClearField(auth.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := auo.mutation.LastModifiedBy(); ok {
		_spec.SetField(auth.FieldLastModifiedBy, field.TypeUUID, value)
	}
	if auo.mutation.LastModifiedByCleared() {
		_spec.ClearField(auth.FieldLastModifiedBy, field.TypeUUID)
	}
	if auo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   auth.EmployeeTable,
			Columns: []string{auth.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   auth.EmployeeTable,
			Columns: []string{auth.EmployeeColumn},
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
	_node = &Auth{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{auth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
