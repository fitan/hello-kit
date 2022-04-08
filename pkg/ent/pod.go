// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hello/pkg/ent/pod"
	"hello/pkg/ent/spiderdevtblservicetree"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Pod is the model entity for the Pod schema.
type Pod struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ClusterName holds the value of the "cluster_name" field.
	// 集群标识
	ClusterName string `json:"cluster_name,omitempty"`
	// Namespace holds the value of the "namespace" field.
	// 所属项目空间
	Namespace string `json:"namespace,omitempty"`
	// ServiceName holds the value of the "service_name" field.
	// 应用名称
	ServiceName string `json:"service_name,omitempty"`
	// PodName holds the value of the "pod_name" field.
	// pod名称
	PodName string `json:"pod_name,omitempty"`
	// ResourceVersion holds the value of the "resource_version" field.
	// 资源版本号
	ResourceVersion string `json:"resource_version,omitempty"`
	// PodIP holds the value of the "pod_ip" field.
	// PodIP
	PodIP string `json:"pod_ip,omitempty"`
	// HostIP holds the value of the "host_ip" field.
	// 主机IP
	HostIP string `json:"host_ip,omitempty"`
	// StartTime holds the value of the "start_time" field.
	// 启动时间
	StartTime time.Time `json:"start_time,omitempty"`
	// Phase holds the value of the "phase" field.
	// 阶段
	Phase string `json:"phase,omitempty"`
	// Reason holds the value of the "reason" field.
	// 原因
	Reason string `json:"reason,omitempty"`
	// Message holds the value of the "message" field.
	// 消息
	Message string `json:"message,omitempty"`
	// Detail holds the value of the "detail" field.
	// 详细内容
	Detail string `json:"detail,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PodQuery when eager-loading is set.
	Edges     PodEdges `json:"edges"`
	aname     *int32
	user_pods *int
}

// PodEdges holds the relations/edges for other nodes in the graph.
type PodEdges struct {
	// Servicetree holds the value of the servicetree edge.
	Servicetree *SpiderDevTblServicetree `json:"servicetree,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ServicetreeOrErr returns the Servicetree value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PodEdges) ServicetreeOrErr() (*SpiderDevTblServicetree, error) {
	if e.loadedTypes[0] {
		if e.Servicetree == nil {
			// The edge servicetree was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: spiderdevtblservicetree.Label}
		}
		return e.Servicetree, nil
	}
	return nil, &NotLoadedError{edge: "servicetree"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pod) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pod.FieldID:
			values[i] = new(sql.NullInt64)
		case pod.FieldClusterName, pod.FieldNamespace, pod.FieldServiceName, pod.FieldPodName, pod.FieldResourceVersion, pod.FieldPodIP, pod.FieldHostIP, pod.FieldPhase, pod.FieldReason, pod.FieldMessage, pod.FieldDetail:
			values[i] = new(sql.NullString)
		case pod.FieldStartTime, pod.FieldCreatedAt, pod.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case pod.ForeignKeys[0]: // aname
			values[i] = new(sql.NullInt64)
		case pod.ForeignKeys[1]: // user_pods
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pod", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pod fields.
func (po *Pod) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pod.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int64(value.Int64)
		case pod.FieldClusterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cluster_name", values[i])
			} else if value.Valid {
				po.ClusterName = value.String
			}
		case pod.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				po.Namespace = value.String
			}
		case pod.FieldServiceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_name", values[i])
			} else if value.Valid {
				po.ServiceName = value.String
			}
		case pod.FieldPodName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pod_name", values[i])
			} else if value.Valid {
				po.PodName = value.String
			}
		case pod.FieldResourceVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_version", values[i])
			} else if value.Valid {
				po.ResourceVersion = value.String
			}
		case pod.FieldPodIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pod_ip", values[i])
			} else if value.Valid {
				po.PodIP = value.String
			}
		case pod.FieldHostIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_ip", values[i])
			} else if value.Valid {
				po.HostIP = value.String
			}
		case pod.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				po.StartTime = value.Time
			}
		case pod.FieldPhase:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phase", values[i])
			} else if value.Valid {
				po.Phase = value.String
			}
		case pod.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				po.Reason = value.String
			}
		case pod.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				po.Message = value.String
			}
		case pod.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				po.Detail = value.String
			}
		case pod.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case pod.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case pod.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field aname", value)
			} else if value.Valid {
				po.aname = new(int32)
				*po.aname = int32(value.Int64)
			}
		case pod.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_pods", value)
			} else if value.Valid {
				po.user_pods = new(int)
				*po.user_pods = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryServicetree queries the "servicetree" edge of the Pod entity.
func (po *Pod) QueryServicetree() *SpiderDevTblServicetreeQuery {
	return (&PodClient{config: po.config}).QueryServicetree(po)
}

// Update returns a builder for updating this Pod.
// Note that you need to call Pod.Unwrap() before calling this method if this Pod
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pod) Update() *PodUpdateOne {
	return (&PodClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pod entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pod) Unwrap() *Pod {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pod is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pod) String() string {
	var builder strings.Builder
	builder.WriteString("Pod(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", cluster_name=")
	builder.WriteString(po.ClusterName)
	builder.WriteString(", namespace=")
	builder.WriteString(po.Namespace)
	builder.WriteString(", service_name=")
	builder.WriteString(po.ServiceName)
	builder.WriteString(", pod_name=")
	builder.WriteString(po.PodName)
	builder.WriteString(", resource_version=")
	builder.WriteString(po.ResourceVersion)
	builder.WriteString(", pod_ip=")
	builder.WriteString(po.PodIP)
	builder.WriteString(", host_ip=")
	builder.WriteString(po.HostIP)
	builder.WriteString(", start_time=")
	builder.WriteString(po.StartTime.Format(time.ANSIC))
	builder.WriteString(", phase=")
	builder.WriteString(po.Phase)
	builder.WriteString(", reason=")
	builder.WriteString(po.Reason)
	builder.WriteString(", message=")
	builder.WriteString(po.Message)
	builder.WriteString(", detail=")
	builder.WriteString(po.Detail)
	builder.WriteString(", created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pods is a parsable slice of Pod.
type Pods []*Pod

func (po Pods) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
