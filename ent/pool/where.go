// Code generated by ent, DO NOT EDIT.

package pool

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldStatus, v))
}

// RoundID applies equality check predicate on the "round_id" field. It's identical to RoundIDEQ.
func RoundID(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldRoundID, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldMode, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldType, v))
}

// LambNum applies equality check predicate on the "lamb_num" field. It's identical to LambNumEQ.
func LambNum(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldLambNum, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Pool {
	return predicate.Pool(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Pool {
	return predicate.Pool(sql.FieldNotNull(FieldStatus))
}

// RoundIDEQ applies the EQ predicate on the "round_id" field.
func RoundIDEQ(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldRoundID, v))
}

// RoundIDNEQ applies the NEQ predicate on the "round_id" field.
func RoundIDNEQ(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldRoundID, v))
}

// RoundIDIn applies the In predicate on the "round_id" field.
func RoundIDIn(vs ...uint64) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldRoundID, vs...))
}

// RoundIDNotIn applies the NotIn predicate on the "round_id" field.
func RoundIDNotIn(vs ...uint64) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldRoundID, vs...))
}

// RoundIDGT applies the GT predicate on the "round_id" field.
func RoundIDGT(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldRoundID, v))
}

// RoundIDGTE applies the GTE predicate on the "round_id" field.
func RoundIDGTE(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldRoundID, v))
}

// RoundIDLT applies the LT predicate on the "round_id" field.
func RoundIDLT(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldRoundID, v))
}

// RoundIDLTE applies the LTE predicate on the "round_id" field.
func RoundIDLTE(v uint64) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldRoundID, v))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldMode, v))
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.Pool {
	return predicate.Pool(sql.FieldContains(FieldMode, v))
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.Pool {
	return predicate.Pool(sql.FieldHasPrefix(FieldMode, v))
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.Pool {
	return predicate.Pool(sql.FieldHasSuffix(FieldMode, v))
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEqualFold(FieldMode, v))
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.Pool {
	return predicate.Pool(sql.FieldContainsFold(FieldMode, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint32) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint32) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint32) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldType, v))
}

// LambNumEQ applies the EQ predicate on the "lamb_num" field.
func LambNumEQ(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldLambNum, v))
}

// LambNumNEQ applies the NEQ predicate on the "lamb_num" field.
func LambNumNEQ(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldLambNum, v))
}

// LambNumIn applies the In predicate on the "lamb_num" field.
func LambNumIn(vs ...float64) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldLambNum, vs...))
}

// LambNumNotIn applies the NotIn predicate on the "lamb_num" field.
func LambNumNotIn(vs ...float64) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldLambNum, vs...))
}

// LambNumGT applies the GT predicate on the "lamb_num" field.
func LambNumGT(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldLambNum, v))
}

// LambNumGTE applies the GTE predicate on the "lamb_num" field.
func LambNumGTE(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldLambNum, v))
}

// LambNumLT applies the LT predicate on the "lamb_num" field.
func LambNumLT(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldLambNum, v))
}

// LambNumLTE applies the LTE predicate on the "lamb_num" field.
func LambNumLTE(v float64) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldLambNum, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Pool {
	return predicate.Pool(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Pool {
	return predicate.Pool(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Pool {
	return predicate.Pool(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Pool {
	return predicate.Pool(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Pool {
	return predicate.Pool(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Pool {
	return predicate.Pool(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Pool {
	return predicate.Pool(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Pool {
	return predicate.Pool(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Pool {
	return predicate.Pool(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Pool {
	return predicate.Pool(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Pool {
	return predicate.Pool(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Pool {
	return predicate.Pool(sql.FieldContainsFold(FieldRemark, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pool) predicate.Pool {
	return predicate.Pool(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pool) predicate.Pool {
	return predicate.Pool(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Pool) predicate.Pool {
	return predicate.Pool(sql.NotPredicates(p))
}
