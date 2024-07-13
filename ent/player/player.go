// Code generated by ent, DO NOT EDIT.

package player

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldTransactionPassword holds the string denoting the transaction_password field in the database.
	FieldTransactionPassword = "transaction_password"
	// FieldCoinLamb holds the string denoting the coin_lamb field in the database.
	FieldCoinLamb = "coin_lamb"
	// FieldTokenLamb holds the string denoting the token_lamb field in the database.
	FieldTokenLamb = "token_lamb"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldDepositAddress holds the string denoting the deposit_address field in the database.
	FieldDepositAddress = "deposit_address"
	// FieldInvitedNum holds the string denoting the invited_num field in the database.
	FieldInvitedNum = "invited_num"
	// FieldTotalIncome holds the string denoting the total_income field in the database.
	FieldTotalIncome = "total_income"
	// FieldProfitAndLoss holds the string denoting the profit_and_loss field in the database.
	FieldProfitAndLoss = "profit_and_loss"
	// FieldInviteCode holds the string denoting the invite_code field in the database.
	FieldInviteCode = "invite_code"
	// FieldInviterID holds the string denoting the inviter_id field in the database.
	FieldInviterID = "inviter_id"
	// FieldInvitedCode holds the string denoting the invited_code field in the database.
	FieldInvitedCode = "invited_code"
	// FieldSystemCommission holds the string denoting the system_commission field in the database.
	FieldSystemCommission = "system_commission"
	// FieldGcicsUserID holds the string denoting the gcics_user_id field in the database.
	FieldGcicsUserID = "gcics_user_id"
	// FieldGcicsToken holds the string denoting the gcics_token field in the database.
	FieldGcicsToken = "gcics_token"
	// FieldGcicsReturnURL holds the string denoting the gcics_return_url field in the database.
	FieldGcicsReturnURL = "gcics_return_url"
	// EdgeInviter holds the string denoting the inviter edge name in mutations.
	EdgeInviter = "inviter"
	// EdgeInvitees holds the string denoting the invitees edge name in mutations.
	EdgeInvitees = "invitees"
	// Table holds the table name of the player in the database.
	Table = "wl_player"
	// InviterTable is the table that holds the inviter relation/edge.
	InviterTable = "wl_player"
	// InviterColumn is the table column denoting the inviter relation/edge.
	InviterColumn = "inviter_id"
	// InviteesTable is the table that holds the invitees relation/edge.
	InviteesTable = "wl_player"
	// InviteesColumn is the table column denoting the invitees relation/edge.
	InviteesColumn = "inviter_id"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldTransactionPassword,
	FieldCoinLamb,
	FieldTokenLamb,
	FieldRank,
	FieldAmount,
	FieldDepositAddress,
	FieldInvitedNum,
	FieldTotalIncome,
	FieldProfitAndLoss,
	FieldInviteCode,
	FieldInviterID,
	FieldInvitedCode,
	FieldSystemCommission,
	FieldGcicsUserID,
	FieldGcicsToken,
	FieldGcicsReturnURL,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultTransactionPassword holds the default value on creation for the "transaction_password" field.
	DefaultTransactionPassword string
	// DefaultCoinLamb holds the default value on creation for the "coin_lamb" field.
	DefaultCoinLamb float32
	// DefaultTokenLamb holds the default value on creation for the "token_lamb" field.
	DefaultTokenLamb float32
	// DefaultRank holds the default value on creation for the "rank" field.
	DefaultRank uint32
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount float64
	// DefaultDepositAddress holds the default value on creation for the "deposit_address" field.
	DefaultDepositAddress string
	// DefaultInvitedNum holds the default value on creation for the "invited_num" field.
	DefaultInvitedNum uint32
	// DefaultTotalIncome holds the default value on creation for the "total_income" field.
	DefaultTotalIncome float64
	// DefaultProfitAndLoss holds the default value on creation for the "profit_and_loss" field.
	DefaultProfitAndLoss float32
	// DefaultInviteCode holds the default value on creation for the "invite_code" field.
	DefaultInviteCode string
	// DefaultInviterID holds the default value on creation for the "inviter_id" field.
	DefaultInviterID uint64
	// DefaultInvitedCode holds the default value on creation for the "invited_code" field.
	DefaultInvitedCode string
	// DefaultSystemCommission holds the default value on creation for the "system_commission" field.
	DefaultSystemCommission float32
	// DefaultGcicsUserID holds the default value on creation for the "gcics_user_id" field.
	DefaultGcicsUserID uint64
	// DefaultGcicsToken holds the default value on creation for the "gcics_token" field.
	DefaultGcicsToken string
	// DefaultGcicsReturnURL holds the default value on creation for the "gcics_return_url" field.
	DefaultGcicsReturnURL string
)

// OrderOption defines the ordering options for the Player queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByTransactionPassword orders the results by the transaction_password field.
func ByTransactionPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionPassword, opts...).ToFunc()
}

// ByCoinLamb orders the results by the coin_lamb field.
func ByCoinLamb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinLamb, opts...).ToFunc()
}

// ByTokenLamb orders the results by the token_lamb field.
func ByTokenLamb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenLamb, opts...).ToFunc()
}

// ByRank orders the results by the rank field.
func ByRank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRank, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByDepositAddress orders the results by the deposit_address field.
func ByDepositAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepositAddress, opts...).ToFunc()
}

// ByInvitedNum orders the results by the invited_num field.
func ByInvitedNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvitedNum, opts...).ToFunc()
}

// ByTotalIncome orders the results by the total_income field.
func ByTotalIncome(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalIncome, opts...).ToFunc()
}

// ByProfitAndLoss orders the results by the profit_and_loss field.
func ByProfitAndLoss(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfitAndLoss, opts...).ToFunc()
}

// ByInviteCode orders the results by the invite_code field.
func ByInviteCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInviteCode, opts...).ToFunc()
}

// ByInviterID orders the results by the inviter_id field.
func ByInviterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInviterID, opts...).ToFunc()
}

// ByInvitedCode orders the results by the invited_code field.
func ByInvitedCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvitedCode, opts...).ToFunc()
}

// BySystemCommission orders the results by the system_commission field.
func BySystemCommission(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemCommission, opts...).ToFunc()
}

// ByGcicsUserID orders the results by the gcics_user_id field.
func ByGcicsUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGcicsUserID, opts...).ToFunc()
}

// ByGcicsToken orders the results by the gcics_token field.
func ByGcicsToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGcicsToken, opts...).ToFunc()
}

// ByGcicsReturnURL orders the results by the gcics_return_url field.
func ByGcicsReturnURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGcicsReturnURL, opts...).ToFunc()
}

// ByInviterField orders the results by inviter field.
func ByInviterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInviterStep(), sql.OrderByField(field, opts...))
	}
}

// ByInviteesCount orders the results by invitees count.
func ByInviteesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInviteesStep(), opts...)
	}
}

// ByInvitees orders the results by invitees terms.
func ByInvitees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInviteesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInviterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InviterTable, InviterColumn),
	)
}
func newInviteesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InviteesTable, InviteesColumn),
	)
}
