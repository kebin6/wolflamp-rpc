// Code generated by ent, DO NOT EDIT.

package roundlambfold

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldUpdatedAt, v))
}

// FoldNo applies equality check predicate on the "fold_no" field. It's identical to FoldNoEQ.
func FoldNo(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldFoldNo, v))
}

// LambNum applies equality check predicate on the "lamb_num" field. It's identical to LambNumEQ.
func LambNum(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldLambNum, v))
}

// RoundID applies equality check predicate on the "round_id" field. It's identical to RoundIDEQ.
func RoundID(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldRoundID, v))
}

// ProfitAndLoss applies equality check predicate on the "profit_and_loss" field. It's identical to ProfitAndLossEQ.
func ProfitAndLoss(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldProfitAndLoss, v))
}

// RoundCount applies equality check predicate on the "round_count" field. It's identical to RoundCountEQ.
func RoundCount(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldRoundCount, v))
}

// TotalRoundCount applies equality check predicate on the "total_round_count" field. It's identical to TotalRoundCountEQ.
func TotalRoundCount(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldTotalRoundCount, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldMode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldUpdatedAt, v))
}

// FoldNoEQ applies the EQ predicate on the "fold_no" field.
func FoldNoEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldFoldNo, v))
}

// FoldNoNEQ applies the NEQ predicate on the "fold_no" field.
func FoldNoNEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldFoldNo, v))
}

// FoldNoIn applies the In predicate on the "fold_no" field.
func FoldNoIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldFoldNo, vs...))
}

// FoldNoNotIn applies the NotIn predicate on the "fold_no" field.
func FoldNoNotIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldFoldNo, vs...))
}

// FoldNoGT applies the GT predicate on the "fold_no" field.
func FoldNoGT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldFoldNo, v))
}

// FoldNoGTE applies the GTE predicate on the "fold_no" field.
func FoldNoGTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldFoldNo, v))
}

// FoldNoLT applies the LT predicate on the "fold_no" field.
func FoldNoLT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldFoldNo, v))
}

// FoldNoLTE applies the LTE predicate on the "fold_no" field.
func FoldNoLTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldFoldNo, v))
}

// LambNumEQ applies the EQ predicate on the "lamb_num" field.
func LambNumEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldLambNum, v))
}

// LambNumNEQ applies the NEQ predicate on the "lamb_num" field.
func LambNumNEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldLambNum, v))
}

// LambNumIn applies the In predicate on the "lamb_num" field.
func LambNumIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldLambNum, vs...))
}

// LambNumNotIn applies the NotIn predicate on the "lamb_num" field.
func LambNumNotIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldLambNum, vs...))
}

// LambNumGT applies the GT predicate on the "lamb_num" field.
func LambNumGT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldLambNum, v))
}

// LambNumGTE applies the GTE predicate on the "lamb_num" field.
func LambNumGTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldLambNum, v))
}

// LambNumLT applies the LT predicate on the "lamb_num" field.
func LambNumLT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldLambNum, v))
}

// LambNumLTE applies the LTE predicate on the "lamb_num" field.
func LambNumLTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldLambNum, v))
}

// RoundIDEQ applies the EQ predicate on the "round_id" field.
func RoundIDEQ(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldRoundID, v))
}

// RoundIDNEQ applies the NEQ predicate on the "round_id" field.
func RoundIDNEQ(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldRoundID, v))
}

// RoundIDIn applies the In predicate on the "round_id" field.
func RoundIDIn(vs ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldRoundID, vs...))
}

// RoundIDNotIn applies the NotIn predicate on the "round_id" field.
func RoundIDNotIn(vs ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldRoundID, vs...))
}

// RoundIDIsNil applies the IsNil predicate on the "round_id" field.
func RoundIDIsNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIsNull(FieldRoundID))
}

// RoundIDNotNil applies the NotNil predicate on the "round_id" field.
func RoundIDNotNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotNull(FieldRoundID))
}

// ProfitAndLossEQ applies the EQ predicate on the "profit_and_loss" field.
func ProfitAndLossEQ(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldProfitAndLoss, v))
}

// ProfitAndLossNEQ applies the NEQ predicate on the "profit_and_loss" field.
func ProfitAndLossNEQ(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldProfitAndLoss, v))
}

// ProfitAndLossIn applies the In predicate on the "profit_and_loss" field.
func ProfitAndLossIn(vs ...float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldProfitAndLoss, vs...))
}

// ProfitAndLossNotIn applies the NotIn predicate on the "profit_and_loss" field.
func ProfitAndLossNotIn(vs ...float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldProfitAndLoss, vs...))
}

// ProfitAndLossGT applies the GT predicate on the "profit_and_loss" field.
func ProfitAndLossGT(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldProfitAndLoss, v))
}

// ProfitAndLossGTE applies the GTE predicate on the "profit_and_loss" field.
func ProfitAndLossGTE(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldProfitAndLoss, v))
}

// ProfitAndLossLT applies the LT predicate on the "profit_and_loss" field.
func ProfitAndLossLT(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldProfitAndLoss, v))
}

// ProfitAndLossLTE applies the LTE predicate on the "profit_and_loss" field.
func ProfitAndLossLTE(v float32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldProfitAndLoss, v))
}

// RoundCountEQ applies the EQ predicate on the "round_count" field.
func RoundCountEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldRoundCount, v))
}

// RoundCountNEQ applies the NEQ predicate on the "round_count" field.
func RoundCountNEQ(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldRoundCount, v))
}

// RoundCountIn applies the In predicate on the "round_count" field.
func RoundCountIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldRoundCount, vs...))
}

// RoundCountNotIn applies the NotIn predicate on the "round_count" field.
func RoundCountNotIn(vs ...uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldRoundCount, vs...))
}

// RoundCountGT applies the GT predicate on the "round_count" field.
func RoundCountGT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldRoundCount, v))
}

// RoundCountGTE applies the GTE predicate on the "round_count" field.
func RoundCountGTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldRoundCount, v))
}

// RoundCountLT applies the LT predicate on the "round_count" field.
func RoundCountLT(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldRoundCount, v))
}

// RoundCountLTE applies the LTE predicate on the "round_count" field.
func RoundCountLTE(v uint32) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldRoundCount, v))
}

// RoundCountIsNil applies the IsNil predicate on the "round_count" field.
func RoundCountIsNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIsNull(FieldRoundCount))
}

// RoundCountNotNil applies the NotNil predicate on the "round_count" field.
func RoundCountNotNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotNull(FieldRoundCount))
}

// TotalRoundCountEQ applies the EQ predicate on the "total_round_count" field.
func TotalRoundCountEQ(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldTotalRoundCount, v))
}

// TotalRoundCountNEQ applies the NEQ predicate on the "total_round_count" field.
func TotalRoundCountNEQ(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldTotalRoundCount, v))
}

// TotalRoundCountIn applies the In predicate on the "total_round_count" field.
func TotalRoundCountIn(vs ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldTotalRoundCount, vs...))
}

// TotalRoundCountNotIn applies the NotIn predicate on the "total_round_count" field.
func TotalRoundCountNotIn(vs ...uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldTotalRoundCount, vs...))
}

// TotalRoundCountGT applies the GT predicate on the "total_round_count" field.
func TotalRoundCountGT(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldTotalRoundCount, v))
}

// TotalRoundCountGTE applies the GTE predicate on the "total_round_count" field.
func TotalRoundCountGTE(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldTotalRoundCount, v))
}

// TotalRoundCountLT applies the LT predicate on the "total_round_count" field.
func TotalRoundCountLT(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldTotalRoundCount, v))
}

// TotalRoundCountLTE applies the LTE predicate on the "total_round_count" field.
func TotalRoundCountLTE(v uint64) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldTotalRoundCount, v))
}

// TotalRoundCountIsNil applies the IsNil predicate on the "total_round_count" field.
func TotalRoundCountIsNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIsNull(FieldTotalRoundCount))
}

// TotalRoundCountNotNil applies the NotNil predicate on the "total_round_count" field.
func TotalRoundCountNotNil() predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotNull(FieldTotalRoundCount))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldLTE(FieldMode, v))
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldContains(FieldMode, v))
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldHasPrefix(FieldMode, v))
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldHasSuffix(FieldMode, v))
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldEqualFold(FieldMode, v))
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.FieldContainsFold(FieldMode, v))
}

// HasRound applies the HasEdge predicate on the "round" edge.
func HasRound() predicate.RoundLambFold {
	return predicate.RoundLambFold(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoundTable, RoundColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoundWith applies the HasEdge predicate on the "round" edge with a given conditions (other predicates).
func HasRoundWith(preds ...predicate.Round) predicate.RoundLambFold {
	return predicate.RoundLambFold(func(s *sql.Selector) {
		step := newRoundStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoundLambFold) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoundLambFold) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoundLambFold) predicate.RoundLambFold {
	return predicate.RoundLambFold(sql.NotPredicates(p))
}
