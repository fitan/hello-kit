// Code generated by entc, DO NOT EDIT.

package role

import (
	"time"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// ResourcesTable is the table that holds the resources relation/edge.
	ResourcesTable = "resources"
	// ResourcesInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourcesInverseTable = "resources"
	// ResourcesColumn is the table column denoting the resources relation/edge.
	ResourcesColumn = "role_resources"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldLevel,
	FieldComments,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)