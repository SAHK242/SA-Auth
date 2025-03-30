package schema

import "entgo.io/ent"
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("employee_id", uuid.UUID{}).Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.Int32("gender"),
		field.Time("date_of_birth"),
		field.String("code").NotEmpty().Unique(),
		field.String("address"),
		field.Time("start_date"),
		field.Time("end_date").Nillable().Optional(),
		field.String("phone_number"),
		field.String("degree_name"),
		field.Int32("degree_year"),
		field.UUID("department_id", uuid.UUID{}),
		field.Int32("employee_type"),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("auth", Auth.Type, "employee", "employee_id"),
		mustBelongToOne("department", Department.Type, "employees", "department_id"),
		hasOne("doctor", Doctor.Type),
		hasOne("nurse", Nurse.Type),
	}
}

func (Employee) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("employee"),
	}
}
