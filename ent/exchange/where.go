// Code generated by ent, DO NOT EDIT.

package exchange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldStatus, v))
}

// PlayerID applies equality check predicate on the "player_id" field. It's identical to PlayerIDEQ.
func PlayerID(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldPlayerID, v))
}

// TransactionID applies equality check predicate on the "transaction_id" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldTransactionID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldType, v))
}

// CoinNum applies equality check predicate on the "coin_num" field. It's identical to CoinNumEQ.
func CoinNum(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldCoinNum, v))
}

// LampNum applies equality check predicate on the "lamp_num" field. It's identical to LampNumEQ.
func LampNum(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldLampNum, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Exchange {
	return predicate.Exchange(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Exchange {
	return predicate.Exchange(sql.FieldNotNull(FieldStatus))
}

// PlayerIDEQ applies the EQ predicate on the "player_id" field.
func PlayerIDEQ(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldPlayerID, v))
}

// PlayerIDNEQ applies the NEQ predicate on the "player_id" field.
func PlayerIDNEQ(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldPlayerID, v))
}

// PlayerIDIn applies the In predicate on the "player_id" field.
func PlayerIDIn(vs ...uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldPlayerID, vs...))
}

// PlayerIDNotIn applies the NotIn predicate on the "player_id" field.
func PlayerIDNotIn(vs ...uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldPlayerID, vs...))
}

// PlayerIDGT applies the GT predicate on the "player_id" field.
func PlayerIDGT(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldPlayerID, v))
}

// PlayerIDGTE applies the GTE predicate on the "player_id" field.
func PlayerIDGTE(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldPlayerID, v))
}

// PlayerIDLT applies the LT predicate on the "player_id" field.
func PlayerIDLT(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldPlayerID, v))
}

// PlayerIDLTE applies the LTE predicate on the "player_id" field.
func PlayerIDLTE(v uint64) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldPlayerID, v))
}

// TransactionIDEQ applies the EQ predicate on the "transaction_id" field.
func TransactionIDEQ(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldTransactionID, v))
}

// TransactionIDNEQ applies the NEQ predicate on the "transaction_id" field.
func TransactionIDNEQ(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldTransactionID, v))
}

// TransactionIDIn applies the In predicate on the "transaction_id" field.
func TransactionIDIn(vs ...string) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldTransactionID, vs...))
}

// TransactionIDNotIn applies the NotIn predicate on the "transaction_id" field.
func TransactionIDNotIn(vs ...string) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldTransactionID, vs...))
}

// TransactionIDGT applies the GT predicate on the "transaction_id" field.
func TransactionIDGT(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldTransactionID, v))
}

// TransactionIDGTE applies the GTE predicate on the "transaction_id" field.
func TransactionIDGTE(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldTransactionID, v))
}

// TransactionIDLT applies the LT predicate on the "transaction_id" field.
func TransactionIDLT(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldTransactionID, v))
}

// TransactionIDLTE applies the LTE predicate on the "transaction_id" field.
func TransactionIDLTE(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldTransactionID, v))
}

// TransactionIDContains applies the Contains predicate on the "transaction_id" field.
func TransactionIDContains(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldContains(FieldTransactionID, v))
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "transaction_id" field.
func TransactionIDHasPrefix(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldHasPrefix(FieldTransactionID, v))
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "transaction_id" field.
func TransactionIDHasSuffix(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldHasSuffix(FieldTransactionID, v))
}

// TransactionIDEqualFold applies the EqualFold predicate on the "transaction_id" field.
func TransactionIDEqualFold(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldEqualFold(FieldTransactionID, v))
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "transaction_id" field.
func TransactionIDContainsFold(v string) predicate.Exchange {
	return predicate.Exchange(sql.FieldContainsFold(FieldTransactionID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldType, v))
}

// CoinNumEQ applies the EQ predicate on the "coin_num" field.
func CoinNumEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldCoinNum, v))
}

// CoinNumNEQ applies the NEQ predicate on the "coin_num" field.
func CoinNumNEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldCoinNum, v))
}

// CoinNumIn applies the In predicate on the "coin_num" field.
func CoinNumIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldCoinNum, vs...))
}

// CoinNumNotIn applies the NotIn predicate on the "coin_num" field.
func CoinNumNotIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldCoinNum, vs...))
}

// CoinNumGT applies the GT predicate on the "coin_num" field.
func CoinNumGT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldCoinNum, v))
}

// CoinNumGTE applies the GTE predicate on the "coin_num" field.
func CoinNumGTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldCoinNum, v))
}

// CoinNumLT applies the LT predicate on the "coin_num" field.
func CoinNumLT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldCoinNum, v))
}

// CoinNumLTE applies the LTE predicate on the "coin_num" field.
func CoinNumLTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldCoinNum, v))
}

// LampNumEQ applies the EQ predicate on the "lamp_num" field.
func LampNumEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldEQ(FieldLampNum, v))
}

// LampNumNEQ applies the NEQ predicate on the "lamp_num" field.
func LampNumNEQ(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNEQ(FieldLampNum, v))
}

// LampNumIn applies the In predicate on the "lamp_num" field.
func LampNumIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldIn(FieldLampNum, vs...))
}

// LampNumNotIn applies the NotIn predicate on the "lamp_num" field.
func LampNumNotIn(vs ...uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldNotIn(FieldLampNum, vs...))
}

// LampNumGT applies the GT predicate on the "lamp_num" field.
func LampNumGT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGT(FieldLampNum, v))
}

// LampNumGTE applies the GTE predicate on the "lamp_num" field.
func LampNumGTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldGTE(FieldLampNum, v))
}

// LampNumLT applies the LT predicate on the "lamp_num" field.
func LampNumLT(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLT(FieldLampNum, v))
}

// LampNumLTE applies the LTE predicate on the "lamp_num" field.
func LampNumLTE(v uint32) predicate.Exchange {
	return predicate.Exchange(sql.FieldLTE(FieldLampNum, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(sql.NotPredicates(p))
}
