package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SpiderDevTblServicetree holds the schema definition for the SpiderDevTblServicetree entity.
type SpiderDevTblServicetree struct {
	ent.Schema
}


func (SpiderDevTblServicetree) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:       "tbl_servicetree",
		},
	}
}

// Fields of the SpiderDevTblServicetree.
func (SpiderDevTblServicetree) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional(),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("aname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.Int32("pnode").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional(),

		field.Int32("type").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional(),

		field.String("key").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("origin").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),
	}

}

// Edges of the SpiderDevTblServicetree.
func (SpiderDevTblServicetree) Edges() []ent.Edge {
	return nil
}
