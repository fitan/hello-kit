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
		field.Enum("classes").Values("app", "middleware").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.String("lang").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.String("git").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("services").Unique(),
	}
}
