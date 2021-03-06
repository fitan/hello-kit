// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hello/pkg/ent/spiderdevtblservicetree"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// SpiderDevTblServicetree is the model entity for the SpiderDevTblServicetree schema.
type SpiderDevTblServicetree struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Aname holds the value of the "aname" field.
	Aname string `json:"aname,omitempty"`
	// Pnode holds the value of the "pnode" field.
	Pnode int32 `json:"pnode,omitempty"`
	// Type holds the value of the "type" field.
	Type int32 `json:"type,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SpiderDevTblServicetree) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case spiderdevtblservicetree.FieldID, spiderdevtblservicetree.FieldPnode, spiderdevtblservicetree.FieldType:
			values[i] = new(sql.NullInt64)
		case spiderdevtblservicetree.FieldName, spiderdevtblservicetree.FieldAname, spiderdevtblservicetree.FieldKey, spiderdevtblservicetree.FieldOrigin:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SpiderDevTblServicetree", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SpiderDevTblServicetree fields.
func (sdts *SpiderDevTblServicetree) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spiderdevtblservicetree.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sdts.ID = int32(value.Int64)
		case spiderdevtblservicetree.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sdts.Name = value.String
			}
		case spiderdevtblservicetree.FieldAname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field aname", values[i])
			} else if value.Valid {
				sdts.Aname = value.String
			}
		case spiderdevtblservicetree.FieldPnode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pnode", values[i])
			} else if value.Valid {
				sdts.Pnode = int32(value.Int64)
			}
		case spiderdevtblservicetree.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				sdts.Type = int32(value.Int64)
			}
		case spiderdevtblservicetree.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				sdts.Key = value.String
			}
		case spiderdevtblservicetree.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				sdts.Origin = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SpiderDevTblServicetree.
// Note that you need to call SpiderDevTblServicetree.Unwrap() before calling this method if this SpiderDevTblServicetree
// was returned from a transaction, and the transaction was committed or rolled back.
func (sdts *SpiderDevTblServicetree) Update() *SpiderDevTblServicetreeUpdateOne {
	return (&SpiderDevTblServicetreeClient{config: sdts.config}).UpdateOne(sdts)
}

// Unwrap unwraps the SpiderDevTblServicetree entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sdts *SpiderDevTblServicetree) Unwrap() *SpiderDevTblServicetree {
	tx, ok := sdts.config.driver.(*txDriver)
	if !ok {
		panic("ent: SpiderDevTblServicetree is not a transactional entity")
	}
	sdts.config.driver = tx.drv
	return sdts
}

// String implements the fmt.Stringer.
func (sdts *SpiderDevTblServicetree) String() string {
	var builder strings.Builder
	builder.WriteString("SpiderDevTblServicetree(")
	builder.WriteString(fmt.Sprintf("id=%v", sdts.ID))
	builder.WriteString(", name=")
	builder.WriteString(sdts.Name)
	builder.WriteString(", aname=")
	builder.WriteString(sdts.Aname)
	builder.WriteString(", pnode=")
	builder.WriteString(fmt.Sprintf("%v", sdts.Pnode))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", sdts.Type))
	builder.WriteString(", key=")
	builder.WriteString(sdts.Key)
	builder.WriteString(", origin=")
	builder.WriteString(sdts.Origin)
	builder.WriteByte(')')
	return builder.String()
}

// SpiderDevTblServicetrees is a parsable slice of SpiderDevTblServicetree.
type SpiderDevTblServicetrees []*SpiderDevTblServicetree

func (sdts SpiderDevTblServicetrees) config(cfg config) {
	for _i := range sdts {
		sdts[_i].config = cfg
	}
}
