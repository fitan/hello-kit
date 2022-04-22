package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pod holds the schema definition for the Pod entity.
type Pod struct {
	ent.Schema
}

func (Pod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:       "pods",
		},
		NodeAnnotation{
			Create:         true,
			CreateMany:     true,
			Update:         true,
			UpdateMany:     true,
			Read:           true,
			ReadMany:       true,
			Delete:         true,
			DeleteMany:     true,
			PagingMust:     false,
			PagingLimitMax: 100,
			Order:          false,
		},
	}
}

// Fields of the Pod.
func (Pod) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Optional().Unique().Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: NumericQueryOps{
				EQ:    true,
			},
		}),

		field.String("cluster_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("集群标识").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("namespace").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("所属项目空间").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("service_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("应用名称").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("pod_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Optional().Comment("pod名称").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("resource_version").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("资源版本号").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("pod_ip").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("PodIP").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("host_ip").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("主机IP").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.Time("start_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("启动时间").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("phase").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("阶段").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("reason").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Optional().Comment("原因").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.String("message").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Optional().Comment("消息").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.Text("detail").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Optional().Comment("详细内容").Annotations(FieldAnnotation{
			Create:   true,
			Update:   true,
			Read:     true,
			QueryOps: StringQueryOps{EQ: true},
		}),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional(),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional(),
	}
}

// Edges of the Pod.
func (Pod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servicetree", SpiderDevTblServicetree.Type).StorageKey(edge.Symbol("service_name"),edge.Column("aname")).Unique().Annotations(NodeAnnotation{
			Create: true,
			Update: false,
			Read:   false,
			Delete: false,
		}),
	}
}
