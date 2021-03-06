// Code generated by entc, DO NOT EDIT.

package userdevice

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the userdevice type in the database.
	Label = "user_device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
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
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userdevice in the database.
	Table = "user_devices"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_devices"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userdevice fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldStatus,
	FieldType,
	FieldOs,
	FieldPlatform,
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
		return fmt.Errorf("userdevice: invalid enum value for status field: %q", s)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeDESKTOP Type = "DESKTOP"
	TypeMOBILE  Type = "MOBILE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDESKTOP, TypeMOBILE:
		return nil
	default:
		return fmt.Errorf("userdevice: invalid enum value for type field: %q", _type)
	}
}

// Os defines the type for the "os" enum field.
type Os string

// Os values.
const (
	OsMAC     Os = "MAC"
	OsANDROID Os = "ANDROID"
	OsWINDOWS Os = "WINDOWS"
)

func (o Os) String() string {
	return string(o)
}

// OsValidator is a validator for the "os" field enum values. It is called by the builders before save.
func OsValidator(o Os) error {
	switch o {
	case OsMAC, OsANDROID, OsWINDOWS:
		return nil
	default:
		return fmt.Errorf("userdevice: invalid enum value for os field: %q", o)
	}
}

// Platform defines the type for the "platform" enum field.
type Platform string

// Platform values.
const (
	PlatformNative   Platform = "native"
	PlatformChrome   Platform = "chrome"
	PlatformSafari   Platform = "safari"
	PlatformExplorer Platform = "explorer"
)

func (pl Platform) String() string {
	return string(pl)
}

// PlatformValidator is a validator for the "platform" field enum values. It is called by the builders before save.
func PlatformValidator(pl Platform) error {
	switch pl {
	case PlatformNative, PlatformChrome, PlatformSafari, PlatformExplorer:
		return nil
	default:
		return fmt.Errorf("userdevice: invalid enum value for platform field: %q", pl)
	}
}
