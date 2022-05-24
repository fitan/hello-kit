package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

func (Audit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func Annotations() []schema.Annotation {
	return []schema.Annotation{NodeAnnotation{
		Create:     true,
		Read:       true,
		ReadMany:   true,
		Delete:     true,
		DeleteMany: true,
		Order:      true,
	}}
}

// Fields of the Audit.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.String("query").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.String("method").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.Text("request").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.Text("response").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.String("header").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.Int("status_code").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: NumericQueryOps{
				EQ: true,
			},
		}),
		field.String("remote_ip").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.String("client_ip").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
		field.String("cost_time").Optional().Annotations(FieldAnnotation{
			Create: true,
			Update: false,
			Read:   true,
			QueryOps: StringQueryOps{
				Contains: true,
			},
		}),
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return nil
}
