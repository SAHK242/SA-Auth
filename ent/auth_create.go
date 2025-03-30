// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/auth"
	"auth/ent/employee"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthCreate is the builder for creating a Auth entity.
type AuthCreate struct {
	config
	mutation *AuthMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (ac *AuthCreate) SetEmail(s string) *AuthCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetUsername sets the "username" field.
func (ac *AuthCreate) SetUsername(s string) *AuthCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AuthCreate) SetPassword(s string) *AuthCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetState sets the "state" field.
func (ac *AuthCreate) SetState(i int32) *AuthCreate {
	ac.mutation.SetState(i)
	return ac
}

// SetCreatedDate sets the "created_date" field.
func (ac *AuthCreate) SetCreatedDate(t time.Time) *AuthCreate {
	ac.mutation.SetCreatedDate(t)
	return ac
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (ac *AuthCreate) SetNillableCreatedDate(t *time.Time) *AuthCreate {
	if t != nil {
		ac.SetCreatedDate(*t)
	}
	return ac
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (ac *AuthCreate) SetLastModifiedDate(t time.Time) *AuthCreate {
	ac.mutation.SetLastModifiedDate(t)
	return ac
}

// SetNillableLastModifiedDate sets the "last_modified_date" field if the given value is not nil.
func (ac *AuthCreate) SetNillableLastModifiedDate(t *time.Time) *AuthCreate {
	if t != nil {
		ac.SetLastModifiedDate(*t)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *AuthCreate) SetCreatedBy(u uuid.UUID) *AuthCreate {
	ac.mutation.SetCreatedBy(u)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *AuthCreate) SetNillableCreatedBy(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetCreatedBy(*u)
	}
	return ac
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (ac *AuthCreate) SetLastModifiedBy(u uuid.UUID) *AuthCreate {
	ac.mutation.SetLastModifiedBy(u)
	return ac
}

// SetNillableLastModifiedBy sets the "last_modified_by" field if the given value is not nil.
func (ac *AuthCreate) SetNillableLastModifiedBy(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetLastModifiedBy(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AuthCreate) SetID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (ac *AuthCreate) SetEmployeeID(id uuid.UUID) *AuthCreate {
	ac.mutation.SetEmployeeID(id)
	return ac
}

// SetNillableEmployeeID sets the "employee" edge to the Employee entity by ID if the given value is not nil.
func (ac *AuthCreate) SetNillableEmployeeID(id *uuid.UUID) *AuthCreate {
	if id != nil {
		ac = ac.SetEmployeeID(*id)
	}
	return ac
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (ac *AuthCreate) SetEmployee(e *Employee) *AuthCreate {
	return ac.SetEmployeeID(e.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (ac *AuthCreate) Mutation() *AuthMutation {
	return ac.mutation
}

// Save creates the Auth in the database.
func (ac *AuthCreate) Save(ctx context.Context) (*Auth, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthCreate) SaveX(ctx context.Context) *Auth {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuthCreate) defaults() {
	if _, ok := ac.mutation.CreatedDate(); !ok {
		v := auth.DefaultCreatedDate
		ac.mutation.SetCreatedDate(v)
	}
	if _, ok := ac.mutation.LastModifiedDate(); !ok {
		v := auth.DefaultLastModifiedDate
		ac.mutation.SetLastModifiedDate(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := auth.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthCreate) check() error {
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Auth.email"`)}
	}
	if _, ok := ac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Auth.username"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Auth.password"`)}
	}
	if _, ok := ac.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Auth.state"`)}
	}
	if _, ok := ac.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "created_date", err: errors.New(`ent: missing required field "Auth.created_date"`)}
	}
	if _, ok := ac.mutation.LastModifiedDate(); !ok {
		return &ValidationError{Name: "last_modified_date", err: errors.New(`ent: missing required field "Auth.last_modified_date"`)}
	}
	return nil
}

func (ac *AuthCreate) sqlSave(ctx context.Context) (*Auth, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuthCreate) createSpec() (*Auth, *sqlgraph.CreateSpec) {
	var (
		_node = &Auth{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(auth.Table, sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(auth.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.SetField(auth.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(auth.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.State(); ok {
		_spec.SetField(auth.FieldState, field.TypeInt32, value)
		_node.State = value
	}
	if value, ok := ac.mutation.CreatedDate(); ok {
		_spec.SetField(auth.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if value, ok := ac.mutation.LastModifiedDate(); ok {
		_spec.SetField(auth.FieldLastModifiedDate, field.TypeTime, value)
		_node.LastModifiedDate = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(auth.FieldCreatedBy, field.TypeUUID, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.LastModifiedBy(); ok {
		_spec.SetField(auth.FieldLastModifiedBy, field.TypeUUID, value)
		_node.LastModifiedBy = value
	}
	if nodes := ac.mutation.EmployeeIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (ac *AuthCreate) OnConflict(opts ...sql.ConflictOption) *AuthUpsertOne {
	ac.conflict = opts
	return &AuthUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AuthCreate) OnConflictColumns(columns ...string) *AuthUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertOne{
		create: ac,
	}
}

type (
	// AuthUpsertOne is the builder for "upsert"-ing
	//  one Auth node.
	AuthUpsertOne struct {
		create *AuthCreate
	}

	// AuthUpsert is the "OnConflict" setter.
	AuthUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *AuthUpsert) SetEmail(v string) *AuthUpsert {
	u.Set(auth.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsert) UpdateEmail() *AuthUpsert {
	u.SetExcluded(auth.FieldEmail)
	return u
}

// SetUsername sets the "username" field.
func (u *AuthUpsert) SetUsername(v string) *AuthUpsert {
	u.Set(auth.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AuthUpsert) UpdateUsername() *AuthUpsert {
	u.SetExcluded(auth.FieldUsername)
	return u
}

// SetPassword sets the "password" field.
func (u *AuthUpsert) SetPassword(v string) *AuthUpsert {
	u.Set(auth.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsert) UpdatePassword() *AuthUpsert {
	u.SetExcluded(auth.FieldPassword)
	return u
}

// SetState sets the "state" field.
func (u *AuthUpsert) SetState(v int32) *AuthUpsert {
	u.Set(auth.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *AuthUpsert) UpdateState() *AuthUpsert {
	u.SetExcluded(auth.FieldState)
	return u
}

// AddState adds v to the "state" field.
func (u *AuthUpsert) AddState(v int32) *AuthUpsert {
	u.Add(auth.FieldState, v)
	return u
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (u *AuthUpsert) SetLastModifiedDate(v time.Time) *AuthUpsert {
	u.Set(auth.FieldLastModifiedDate, v)
	return u
}

// UpdateLastModifiedDate sets the "last_modified_date" field to the value that was provided on create.
func (u *AuthUpsert) UpdateLastModifiedDate() *AuthUpsert {
	u.SetExcluded(auth.FieldLastModifiedDate)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *AuthUpsert) SetCreatedBy(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AuthUpsert) UpdateCreatedBy() *AuthUpsert {
	u.SetExcluded(auth.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AuthUpsert) ClearCreatedBy() *AuthUpsert {
	u.SetNull(auth.FieldCreatedBy)
	return u
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (u *AuthUpsert) SetLastModifiedBy(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldLastModifiedBy, v)
	return u
}

// UpdateLastModifiedBy sets the "last_modified_by" field to the value that was provided on create.
func (u *AuthUpsert) UpdateLastModifiedBy() *AuthUpsert {
	u.SetExcluded(auth.FieldLastModifiedBy)
	return u
}

// ClearLastModifiedBy clears the value of the "last_modified_by" field.
func (u *AuthUpsert) ClearLastModifiedBy() *AuthUpsert {
	u.SetNull(auth.FieldLastModifiedBy)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthUpsertOne) UpdateNewValues() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(auth.FieldID)
		}
		if _, exists := u.create.mutation.CreatedDate(); exists {
			s.SetIgnore(auth.FieldCreatedDate)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthUpsertOne) Ignore() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertOne) DoNothing() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreate.OnConflict
// documentation for more info.
func (u *AuthUpsertOne) Update(set func(*AuthUpsert)) *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *AuthUpsertOne) SetEmail(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateEmail() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEmail()
	})
}

// SetUsername sets the "username" field.
func (u *AuthUpsertOne) SetUsername(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateUsername() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUsername()
	})
}

// SetPassword sets the "password" field.
func (u *AuthUpsertOne) SetPassword(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdatePassword() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdatePassword()
	})
}

// SetState sets the "state" field.
func (u *AuthUpsertOne) SetState(v int32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *AuthUpsertOne) AddState(v int32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateState() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateState()
	})
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (u *AuthUpsertOne) SetLastModifiedDate(v time.Time) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetLastModifiedDate(v)
	})
}

// UpdateLastModifiedDate sets the "last_modified_date" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateLastModifiedDate() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLastModifiedDate()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AuthUpsertOne) SetCreatedBy(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateCreatedBy() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AuthUpsertOne) ClearCreatedBy() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearCreatedBy()
	})
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (u *AuthUpsertOne) SetLastModifiedBy(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetLastModifiedBy(v)
	})
}

// UpdateLastModifiedBy sets the "last_modified_by" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateLastModifiedBy() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLastModifiedBy()
	})
}

// ClearLastModifiedBy clears the value of the "last_modified_by" field.
func (u *AuthUpsertOne) ClearLastModifiedBy() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearLastModifiedBy()
	})
}

// Exec executes the query.
func (u *AuthUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuthUpsertOne.ID is not supported by MySQL driver. Use AuthUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthCreateBulk is the builder for creating many Auth entities in bulk.
type AuthCreateBulk struct {
	config
	err      error
	builders []*AuthCreate
	conflict []sql.ConflictOption
}

// Save creates the Auth entities in the database.
func (acb *AuthCreateBulk) Save(ctx context.Context) ([]*Auth, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Auth, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthCreateBulk) SaveX(ctx context.Context) []*Auth {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (acb *AuthCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthUpsertBulk {
	acb.conflict = opts
	return &AuthUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AuthCreateBulk) OnConflictColumns(columns ...string) *AuthUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertBulk{
		create: acb,
	}
}

// AuthUpsertBulk is the builder for "upsert"-ing
// a bulk of Auth nodes.
type AuthUpsertBulk struct {
	create *AuthCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthUpsertBulk) UpdateNewValues() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(auth.FieldID)
			}
			if _, exists := b.mutation.CreatedDate(); exists {
				s.SetIgnore(auth.FieldCreatedDate)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthUpsertBulk) Ignore() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertBulk) DoNothing() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreateBulk.OnConflict
// documentation for more info.
func (u *AuthUpsertBulk) Update(set func(*AuthUpsert)) *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *AuthUpsertBulk) SetEmail(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateEmail() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEmail()
	})
}

// SetUsername sets the "username" field.
func (u *AuthUpsertBulk) SetUsername(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateUsername() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUsername()
	})
}

// SetPassword sets the "password" field.
func (u *AuthUpsertBulk) SetPassword(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdatePassword() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdatePassword()
	})
}

// SetState sets the "state" field.
func (u *AuthUpsertBulk) SetState(v int32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *AuthUpsertBulk) AddState(v int32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateState() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateState()
	})
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (u *AuthUpsertBulk) SetLastModifiedDate(v time.Time) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetLastModifiedDate(v)
	})
}

// UpdateLastModifiedDate sets the "last_modified_date" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateLastModifiedDate() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLastModifiedDate()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AuthUpsertBulk) SetCreatedBy(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateCreatedBy() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AuthUpsertBulk) ClearCreatedBy() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearCreatedBy()
	})
}

// SetLastModifiedBy sets the "last_modified_by" field.
func (u *AuthUpsertBulk) SetLastModifiedBy(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetLastModifiedBy(v)
	})
}

// UpdateLastModifiedBy sets the "last_modified_by" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateLastModifiedBy() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLastModifiedBy()
	})
}

// ClearLastModifiedBy clears the value of the "last_modified_by" field.
func (u *AuthUpsertBulk) ClearLastModifiedBy() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearLastModifiedBy()
	})
}

// Exec executes the query.
func (u *AuthUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
