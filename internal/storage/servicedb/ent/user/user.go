// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSignup holds the string denoting the signup field in the database.
	FieldSignup = "signup"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// EdgeMeta holds the string denoting the meta edge name in mutations.
	EdgeMeta = "meta"
	// EdgeOauthProviders holds the string denoting the oauth_providers edge name in mutations.
	EdgeOauthProviders = "oauth_providers"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MetaTable is the table that holds the meta relation/edge.
	MetaTable = "user_meta"
	// MetaInverseTable is the table name for the UserMeta entity.
	// It exists in this package in order to avoid circular dependency with the "usermeta" package.
	MetaInverseTable = "user_meta"
	// MetaColumn is the table column denoting the meta relation/edge.
	MetaColumn = "user_id"
	// OauthProvidersTable is the table that holds the oauth_providers relation/edge.
	OauthProvidersTable = "user_oauth_providers"
	// OauthProvidersInverseTable is the table name for the UserOAuthProvider entity.
	// It exists in this package in order to avoid circular dependency with the "useroauthprovider" package.
	OauthProvidersInverseTable = "user_oauth_providers"
	// OauthProvidersColumn is the table column denoting the oauth_providers relation/edge.
	OauthProvidersColumn = "user_id"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "user_devices"
	// DevicesInverseTable is the table name for the UserDevice entity.
	// It exists in this package in order to avoid circular dependency with the "userdevice" package.
	DevicesInverseTable = "user_devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldStatus,
	FieldType,
	FieldSignup,
	FieldUsername,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSignup holds the default value on creation for the "signup" field.
	DefaultSignup bool
)

// Status defines the type for the "status" enum field.
type Status string

// StatusACTIVE is the default value of the Status enum.
const DefaultStatus = StatusACTIVE

// Status values.
const (
	StatusACTIVE   Status = "ACTIVE"
	StatusINACTIVE Status = "INACTIVE"
	StatusDELETED  Status = "DELETED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusACTIVE, StatusINACTIVE, StatusDELETED:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// TypeUSER is the default value of the Type enum.
const DefaultType = TypeUSER

// Type values.
const (
	TypeUSER  Type = "USER"
	TypeADMIN Type = "ADMIN"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUSER, TypeADMIN:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}
