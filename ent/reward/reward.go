// Code generated by ent, DO NOT EDIT.

package reward

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the reward type in the database.
	Label = "reward"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldToID holds the string denoting the to_id field in the database.
	FieldToID = "to_id"
	// FieldContributorID holds the string denoting the contributor_id field in the database.
	FieldContributorID = "contributor_id"
	// FieldContributorEmail holds the string denoting the contributor_email field in the database.
	FieldContributorEmail = "contributor_email"
	// FieldContributorLevel holds the string denoting the contributor_level field in the database.
	FieldContributorLevel = "contributor_level"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// FieldFormula holds the string denoting the formula field in the database.
	FieldFormula = "formula"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the reward in the database.
	Table = "wl_reward"
)

// Columns holds all SQL columns for reward fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldToID,
	FieldContributorID,
	FieldContributorEmail,
	FieldContributorLevel,
	FieldNum,
	FieldFormula,
	FieldMode,
	FieldRemark,
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
	DefaultStatus uint8
	// DefaultToID holds the default value on creation for the "to_id" field.
	DefaultToID uint64
	// DefaultContributorID holds the default value on creation for the "contributor_id" field.
	DefaultContributorID uint64
	// DefaultContributorEmail holds the default value on creation for the "contributor_email" field.
	DefaultContributorEmail string
	// DefaultContributorLevel holds the default value on creation for the "contributor_level" field.
	DefaultContributorLevel uint32
	// DefaultNum holds the default value on creation for the "num" field.
	DefaultNum float32
	// DefaultFormula holds the default value on creation for the "formula" field.
	DefaultFormula string
	// DefaultMode holds the default value on creation for the "mode" field.
	DefaultMode string
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
)

// OrderOption defines the ordering options for the Reward queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByToID orders the results by the to_id field.
func ByToID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToID, opts...).ToFunc()
}

// ByContributorID orders the results by the contributor_id field.
func ByContributorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContributorID, opts...).ToFunc()
}

// ByContributorEmail orders the results by the contributor_email field.
func ByContributorEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContributorEmail, opts...).ToFunc()
}

// ByContributorLevel orders the results by the contributor_level field.
func ByContributorLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContributorLevel, opts...).ToFunc()
}

// ByNum orders the results by the num field.
func ByNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNum, opts...).ToFunc()
}

// ByFormula orders the results by the formula field.
func ByFormula(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormula, opts...).ToFunc()
}

// ByMode orders the results by the mode field.
func ByMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMode, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}
