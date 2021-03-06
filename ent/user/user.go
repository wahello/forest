// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/vicanso/forest/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRoles holds the string denoting the roles field in the database.
	FieldRoles = "roles"
	// FieldGroups holds the string denoting the groups field in the database.
	FieldGroups = "groups"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldAccount,
	FieldPassword,
	FieldName,
	FieldRoles,
	FieldGroups,
	FieldEmail,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus schema.Status
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int8) error
	// AccountValidator is a validator for the "account" field. It is called by the builders before save.
	AccountValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
