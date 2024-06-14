// Code generated by ent, DO NOT EDIT.

package exchange

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the exchange type in the database.
	Label = "exchange"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPlayerID holds the string denoting the player_id field in the database.
	FieldPlayerID = "player_id"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCoinNum holds the string denoting the coin_num field in the database.
	FieldCoinNum = "coin_num"
	// FieldLampNum holds the string denoting the lamp_num field in the database.
	FieldLampNum = "lamp_num"
	// Table holds the table name of the exchange in the database.
	Table = "wl_exchange"
)

// Columns holds all SQL columns for exchange fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldPlayerID,
	FieldTransactionID,
	FieldType,
	FieldCoinNum,
	FieldLampNum,
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
	// DefaultPlayerID holds the default value on creation for the "player_id" field.
	DefaultPlayerID uint64
	// DefaultTransactionID holds the default value on creation for the "transaction_id" field.
	DefaultTransactionID string
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType uint32
	// DefaultCoinNum holds the default value on creation for the "coin_num" field.
	DefaultCoinNum uint32
	// DefaultLampNum holds the default value on creation for the "lamp_num" field.
	DefaultLampNum uint32
)

// OrderOption defines the ordering options for the Exchange queries.
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

// ByPlayerID orders the results by the player_id field.
func ByPlayerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayerID, opts...).ToFunc()
}

// ByTransactionID orders the results by the transaction_id field.
func ByTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByCoinNum orders the results by the coin_num field.
func ByCoinNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinNum, opts...).ToFunc()
}

// ByLampNum orders the results by the lamp_num field.
func ByLampNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLampNum, opts...).ToFunc()
}