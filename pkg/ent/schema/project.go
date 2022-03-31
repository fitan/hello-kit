package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Project struct {
	ent.Schema
}

// Fields of the User.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("alias"),
		field.String("name").Default("unknown"),
	}
}

// Edges of the User.
func (Project) Edges() []ent.Edge {
	return nil
}
