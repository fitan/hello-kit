package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{NodeAnnotation{
		Create:         true,
		CreateMany:     true,
		Update:         true,
		UpdateMany:     true,
		Read:           true,
		ReadMany:       true,
		Delete:         true,
		DeleteMany:     true,
		PagingMust:     false,
		PagingLimitMax: 200,
		Order:          false,
	}}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("pass_word").StructTag(`json:"pass_word"`).Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.String("token").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.Bool("enable").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: nil,
		}),
		field.Int("age").Optional().StructTag(`fake:"{number:0,150}"`).Annotations(FieldAnnotation{
			Create: true,
			Update: true,
			Read:   true,
			QueryOps: NumericQueryOps{
				EQ:    true,
				NEQ:   false,
				In:    false,
				NotIn: false,
				GT:    false,
				GTE:   false,
				LT:    false,
				LTE:   false,
			},
		}),
		field.String("name").Default("unknown").Annotations(FieldAnnotation{
			Create: true,
			Update: true,
			Read:   true,
			QueryOps: StringQueryOps{
				EQ:        true,
				NEQ:       false,
				In:        false,
				NotIn:     true,
				GT:        false,
				GTE:       false,
				LT:        false,
				LTE:       false,
				Contains:  false,
				HasPrefix: false,
				HasSuffix: false,
			},
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
