// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/department"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges        DepartmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Employees holds the value of the employees edge.
	Employees []*Employee `json:"employees,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeesOrErr returns the Employees value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) EmployeesOrErr() ([]*Employee, error) {
	if e.loadedTypes[0] {
		return e.Employees, nil
	}
	return nil, &NotLoadedError{edge: "employees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldName:
			values[i] = new(sql.NullString)
		case department.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Department.
// This includes values selected through modifiers, order, etc.
func (d *Department) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryEmployees queries the "employees" edge of the Department entity.
func (d *Department) QueryEmployees() *EmployeeQuery {
	return NewDepartmentClient(d.config).QueryEmployees(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return NewDepartmentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department
