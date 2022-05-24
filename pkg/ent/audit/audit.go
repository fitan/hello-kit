// Code generated by entc, DO NOT EDIT.

package audit

import (
	"time"
)

const (
	// Label holds the string label denoting the audit type in the database.
	Label = "audit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldQuery holds the string denoting the query field in the database.
	FieldQuery = "query"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldRequest holds the string denoting the request field in the database.
	FieldRequest = "request"
	// FieldResponse holds the string denoting the response field in the database.
	FieldResponse = "response"
	// FieldHeader holds the string denoting the header field in the database.
	FieldHeader = "header"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldRemoteIP holds the string denoting the remote_ip field in the database.
	FieldRemoteIP = "remote_ip"
	// FieldClientIP holds the string denoting the client_ip field in the database.
	FieldClientIP = "client_ip"
	// FieldCostTime holds the string denoting the cost_time field in the database.
	FieldCostTime = "cost_time"
	// Table holds the table name of the audit in the database.
	Table = "audits"
)

// Columns holds all SQL columns for audit fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldURL,
	FieldQuery,
	FieldMethod,
	FieldRequest,
	FieldResponse,
	FieldHeader,
	FieldStatusCode,
	FieldRemoteIP,
	FieldClientIP,
	FieldCostTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
