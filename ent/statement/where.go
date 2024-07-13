// Code generated by ent, DO NOT EDIT.

package statement

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldStatus, v))
}

// PlayerID applies equality check predicate on the "player_id" field. It's identical to PlayerIDEQ.
func PlayerID(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldPlayerID, v))
}

// Module applies equality check predicate on the "module" field. It's identical to ModuleEQ.
func Module(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldModule, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldCode, v))
}

// InoutType applies equality check predicate on the "inout_type" field. It's identical to InoutTypeEQ.
func InoutType(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldInoutType, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldAmount, v))
}

// ReferID applies equality check predicate on the "refer_id" field. It's identical to ReferIDEQ.
func ReferID(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldReferID, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldMode, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Statement {
	return predicate.Statement(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Statement {
	return predicate.Statement(sql.FieldNotNull(FieldStatus))
}

// PlayerIDEQ applies the EQ predicate on the "player_id" field.
func PlayerIDEQ(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldPlayerID, v))
}

// PlayerIDNEQ applies the NEQ predicate on the "player_id" field.
func PlayerIDNEQ(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldPlayerID, v))
}

// PlayerIDIn applies the In predicate on the "player_id" field.
func PlayerIDIn(vs ...uint64) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldPlayerID, vs...))
}

// PlayerIDNotIn applies the NotIn predicate on the "player_id" field.
func PlayerIDNotIn(vs ...uint64) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldPlayerID, vs...))
}

// PlayerIDGT applies the GT predicate on the "player_id" field.
func PlayerIDGT(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldPlayerID, v))
}

// PlayerIDGTE applies the GTE predicate on the "player_id" field.
func PlayerIDGTE(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldPlayerID, v))
}

// PlayerIDLT applies the LT predicate on the "player_id" field.
func PlayerIDLT(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldPlayerID, v))
}

// PlayerIDLTE applies the LTE predicate on the "player_id" field.
func PlayerIDLTE(v uint64) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldPlayerID, v))
}

// ModuleEQ applies the EQ predicate on the "module" field.
func ModuleEQ(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldModule, v))
}

// ModuleNEQ applies the NEQ predicate on the "module" field.
func ModuleNEQ(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldModule, v))
}

// ModuleIn applies the In predicate on the "module" field.
func ModuleIn(vs ...uint32) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldModule, vs...))
}

// ModuleNotIn applies the NotIn predicate on the "module" field.
func ModuleNotIn(vs ...uint32) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldModule, vs...))
}

// ModuleGT applies the GT predicate on the "module" field.
func ModuleGT(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldModule, v))
}

// ModuleGTE applies the GTE predicate on the "module" field.
func ModuleGTE(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldModule, v))
}

// ModuleLT applies the LT predicate on the "module" field.
func ModuleLT(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldModule, v))
}

// ModuleLTE applies the LTE predicate on the "module" field.
func ModuleLTE(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldModule, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContainsFold(FieldCode, v))
}

// InoutTypeEQ applies the EQ predicate on the "inout_type" field.
func InoutTypeEQ(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldInoutType, v))
}

// InoutTypeNEQ applies the NEQ predicate on the "inout_type" field.
func InoutTypeNEQ(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldInoutType, v))
}

// InoutTypeIn applies the In predicate on the "inout_type" field.
func InoutTypeIn(vs ...uint32) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldInoutType, vs...))
}

// InoutTypeNotIn applies the NotIn predicate on the "inout_type" field.
func InoutTypeNotIn(vs ...uint32) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldInoutType, vs...))
}

// InoutTypeGT applies the GT predicate on the "inout_type" field.
func InoutTypeGT(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldInoutType, v))
}

// InoutTypeGTE applies the GTE predicate on the "inout_type" field.
func InoutTypeGTE(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldInoutType, v))
}

// InoutTypeLT applies the LT predicate on the "inout_type" field.
func InoutTypeLT(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldInoutType, v))
}

// InoutTypeLTE applies the LTE predicate on the "inout_type" field.
func InoutTypeLTE(v uint32) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldInoutType, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldAmount, v))
}

// ReferIDEQ applies the EQ predicate on the "refer_id" field.
func ReferIDEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldReferID, v))
}

// ReferIDNEQ applies the NEQ predicate on the "refer_id" field.
func ReferIDNEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldReferID, v))
}

// ReferIDIn applies the In predicate on the "refer_id" field.
func ReferIDIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldReferID, vs...))
}

// ReferIDNotIn applies the NotIn predicate on the "refer_id" field.
func ReferIDNotIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldReferID, vs...))
}

// ReferIDGT applies the GT predicate on the "refer_id" field.
func ReferIDGT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldReferID, v))
}

// ReferIDGTE applies the GTE predicate on the "refer_id" field.
func ReferIDGTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldReferID, v))
}

// ReferIDLT applies the LT predicate on the "refer_id" field.
func ReferIDLT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldReferID, v))
}

// ReferIDLTE applies the LTE predicate on the "refer_id" field.
func ReferIDLTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldReferID, v))
}

// ReferIDContains applies the Contains predicate on the "refer_id" field.
func ReferIDContains(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContains(FieldReferID, v))
}

// ReferIDHasPrefix applies the HasPrefix predicate on the "refer_id" field.
func ReferIDHasPrefix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasPrefix(FieldReferID, v))
}

// ReferIDHasSuffix applies the HasSuffix predicate on the "refer_id" field.
func ReferIDHasSuffix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasSuffix(FieldReferID, v))
}

// ReferIDEqualFold applies the EqualFold predicate on the "refer_id" field.
func ReferIDEqualFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEqualFold(FieldReferID, v))
}

// ReferIDContainsFold applies the ContainsFold predicate on the "refer_id" field.
func ReferIDContainsFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContainsFold(FieldReferID, v))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldMode, v))
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContains(FieldMode, v))
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasPrefix(FieldMode, v))
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasSuffix(FieldMode, v))
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEqualFold(FieldMode, v))
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContainsFold(FieldMode, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Statement {
	return predicate.Statement(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Statement {
	return predicate.Statement(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Statement {
	return predicate.Statement(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Statement {
	return predicate.Statement(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Statement {
	return predicate.Statement(sql.FieldContainsFold(FieldRemark, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Statement) predicate.Statement {
	return predicate.Statement(sql.NotPredicates(p))
}
