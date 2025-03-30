package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.Int32("state"),
		field.Time("created_date").Immutable().Default(time.Now()),
		field.Time("last_modified_date").Default(time.Now()),
		field.UUID("created_by", uuid.UUID{}).Optional(),
		field.UUID("last_modified_by", uuid.UUID{}).Optional(),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		hasOne("employee", Employee.Type),
	}
}

func (Auth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("auth"),
	}
}
