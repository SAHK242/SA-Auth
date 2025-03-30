package schema

import "entgo.io/ent"
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("nurse_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("employee", Employee.Type, "nurse", "nurse_id"),
	}
}

func (Nurse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("nurse"),
	}
}
