package schema

import "entgo.io/ent"
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Unique(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		hasMany("employees", Employee.Type),
	}
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("department"),
	}
}
