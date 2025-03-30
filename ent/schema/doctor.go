package schema

import "entgo.io/ent"
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("doctor_id", uuid.UUID{}),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("employee", Employee.Type, "doctor", "doctor_id"),
	}
}

func (Doctor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("doctor"),
	}
}
