package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("aname"),
		field.String("comments"),
		field.Enum("classes").Values("app", "middleware"),
		field.String("lang"),
		field.String("git"),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("services").Unique(),
	}
}
