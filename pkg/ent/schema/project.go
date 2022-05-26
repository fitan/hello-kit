package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type Project struct {
	ent.Schema
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.String("aname").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.String("comments").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
	}
}

// Edges of the User.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
	}
}
