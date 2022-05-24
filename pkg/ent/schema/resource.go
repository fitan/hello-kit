package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

func (Resource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("key"),
		field.String("path"),
		field.String("action"),
		field.String("comments"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return nil
}
