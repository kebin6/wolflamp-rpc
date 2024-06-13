// Code generated by ent, DO NOT EDIT.

package roundinvest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldUpdatedAt, v))
}

// PlayerID applies equality check predicate on the "player_id" field. It's identical to PlayerIDEQ.
func PlayerID(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldPlayerID, v))
}

// PlayerEmail applies equality check predicate on the "player_email" field. It's identical to PlayerEmailEQ.
func PlayerEmail(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldPlayerEmail, v))
}

// FoldNo applies equality check predicate on the "fold_no" field. It's identical to FoldNoEQ.
func FoldNo(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldFoldNo, v))
}

// LambNum applies equality check predicate on the "lamb_num" field. It's identical to LambNumEQ.
func LambNum(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldLambNum, v))
}

// ProfitAndLoss applies equality check predicate on the "profit_and_loss" field. It's identical to ProfitAndLossEQ.
func ProfitAndLoss(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldProfitAndLoss, v))
}

// RoundID applies equality check predicate on the "round_id" field. It's identical to RoundIDEQ.
func RoundID(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldRoundID, v))
}

// RoundCount applies equality check predicate on the "round_count" field. It's identical to RoundCountEQ.
func RoundCount(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldRoundCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldUpdatedAt, v))
}

// PlayerIDEQ applies the EQ predicate on the "player_id" field.
func PlayerIDEQ(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldPlayerID, v))
}

// PlayerIDNEQ applies the NEQ predicate on the "player_id" field.
func PlayerIDNEQ(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldPlayerID, v))
}

// PlayerIDIn applies the In predicate on the "player_id" field.
func PlayerIDIn(vs ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldPlayerID, vs...))
}

// PlayerIDNotIn applies the NotIn predicate on the "player_id" field.
func PlayerIDNotIn(vs ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldPlayerID, vs...))
}

// PlayerIDGT applies the GT predicate on the "player_id" field.
func PlayerIDGT(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldPlayerID, v))
}

// PlayerIDGTE applies the GTE predicate on the "player_id" field.
func PlayerIDGTE(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldPlayerID, v))
}

// PlayerIDLT applies the LT predicate on the "player_id" field.
func PlayerIDLT(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldPlayerID, v))
}

// PlayerIDLTE applies the LTE predicate on the "player_id" field.
func PlayerIDLTE(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldPlayerID, v))
}

// PlayerEmailEQ applies the EQ predicate on the "player_email" field.
func PlayerEmailEQ(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldPlayerEmail, v))
}

// PlayerEmailNEQ applies the NEQ predicate on the "player_email" field.
func PlayerEmailNEQ(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldPlayerEmail, v))
}

// PlayerEmailIn applies the In predicate on the "player_email" field.
func PlayerEmailIn(vs ...string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldPlayerEmail, vs...))
}

// PlayerEmailNotIn applies the NotIn predicate on the "player_email" field.
func PlayerEmailNotIn(vs ...string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldPlayerEmail, vs...))
}

// PlayerEmailGT applies the GT predicate on the "player_email" field.
func PlayerEmailGT(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldPlayerEmail, v))
}

// PlayerEmailGTE applies the GTE predicate on the "player_email" field.
func PlayerEmailGTE(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldPlayerEmail, v))
}

// PlayerEmailLT applies the LT predicate on the "player_email" field.
func PlayerEmailLT(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldPlayerEmail, v))
}

// PlayerEmailLTE applies the LTE predicate on the "player_email" field.
func PlayerEmailLTE(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldPlayerEmail, v))
}

// PlayerEmailContains applies the Contains predicate on the "player_email" field.
func PlayerEmailContains(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldContains(FieldPlayerEmail, v))
}

// PlayerEmailHasPrefix applies the HasPrefix predicate on the "player_email" field.
func PlayerEmailHasPrefix(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldHasPrefix(FieldPlayerEmail, v))
}

// PlayerEmailHasSuffix applies the HasSuffix predicate on the "player_email" field.
func PlayerEmailHasSuffix(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldHasSuffix(FieldPlayerEmail, v))
}

// PlayerEmailEqualFold applies the EqualFold predicate on the "player_email" field.
func PlayerEmailEqualFold(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEqualFold(FieldPlayerEmail, v))
}

// PlayerEmailContainsFold applies the ContainsFold predicate on the "player_email" field.
func PlayerEmailContainsFold(v string) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldContainsFold(FieldPlayerEmail, v))
}

// FoldNoEQ applies the EQ predicate on the "fold_no" field.
func FoldNoEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldFoldNo, v))
}

// FoldNoNEQ applies the NEQ predicate on the "fold_no" field.
func FoldNoNEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldFoldNo, v))
}

// FoldNoIn applies the In predicate on the "fold_no" field.
func FoldNoIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldFoldNo, vs...))
}

// FoldNoNotIn applies the NotIn predicate on the "fold_no" field.
func FoldNoNotIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldFoldNo, vs...))
}

// FoldNoGT applies the GT predicate on the "fold_no" field.
func FoldNoGT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldFoldNo, v))
}

// FoldNoGTE applies the GTE predicate on the "fold_no" field.
func FoldNoGTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldFoldNo, v))
}

// FoldNoLT applies the LT predicate on the "fold_no" field.
func FoldNoLT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldFoldNo, v))
}

// FoldNoLTE applies the LTE predicate on the "fold_no" field.
func FoldNoLTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldFoldNo, v))
}

// LambNumEQ applies the EQ predicate on the "lamb_num" field.
func LambNumEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldLambNum, v))
}

// LambNumNEQ applies the NEQ predicate on the "lamb_num" field.
func LambNumNEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldLambNum, v))
}

// LambNumIn applies the In predicate on the "lamb_num" field.
func LambNumIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldLambNum, vs...))
}

// LambNumNotIn applies the NotIn predicate on the "lamb_num" field.
func LambNumNotIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldLambNum, vs...))
}

// LambNumGT applies the GT predicate on the "lamb_num" field.
func LambNumGT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldLambNum, v))
}

// LambNumGTE applies the GTE predicate on the "lamb_num" field.
func LambNumGTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldLambNum, v))
}

// LambNumLT applies the LT predicate on the "lamb_num" field.
func LambNumLT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldLambNum, v))
}

// LambNumLTE applies the LTE predicate on the "lamb_num" field.
func LambNumLTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldLambNum, v))
}

// ProfitAndLossEQ applies the EQ predicate on the "profit_and_loss" field.
func ProfitAndLossEQ(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldProfitAndLoss, v))
}

// ProfitAndLossNEQ applies the NEQ predicate on the "profit_and_loss" field.
func ProfitAndLossNEQ(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldProfitAndLoss, v))
}

// ProfitAndLossIn applies the In predicate on the "profit_and_loss" field.
func ProfitAndLossIn(vs ...float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldProfitAndLoss, vs...))
}

// ProfitAndLossNotIn applies the NotIn predicate on the "profit_and_loss" field.
func ProfitAndLossNotIn(vs ...float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldProfitAndLoss, vs...))
}

// ProfitAndLossGT applies the GT predicate on the "profit_and_loss" field.
func ProfitAndLossGT(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldProfitAndLoss, v))
}

// ProfitAndLossGTE applies the GTE predicate on the "profit_and_loss" field.
func ProfitAndLossGTE(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldProfitAndLoss, v))
}

// ProfitAndLossLT applies the LT predicate on the "profit_and_loss" field.
func ProfitAndLossLT(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldProfitAndLoss, v))
}

// ProfitAndLossLTE applies the LTE predicate on the "profit_and_loss" field.
func ProfitAndLossLTE(v float32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldProfitAndLoss, v))
}

// RoundIDEQ applies the EQ predicate on the "round_id" field.
func RoundIDEQ(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldRoundID, v))
}

// RoundIDNEQ applies the NEQ predicate on the "round_id" field.
func RoundIDNEQ(v uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldRoundID, v))
}

// RoundIDIn applies the In predicate on the "round_id" field.
func RoundIDIn(vs ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldRoundID, vs...))
}

// RoundIDNotIn applies the NotIn predicate on the "round_id" field.
func RoundIDNotIn(vs ...uint64) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldRoundID, vs...))
}

// RoundIDIsNil applies the IsNil predicate on the "round_id" field.
func RoundIDIsNil() predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIsNull(FieldRoundID))
}

// RoundIDNotNil applies the NotNil predicate on the "round_id" field.
func RoundIDNotNil() predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotNull(FieldRoundID))
}

// RoundCountEQ applies the EQ predicate on the "round_count" field.
func RoundCountEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldEQ(FieldRoundCount, v))
}

// RoundCountNEQ applies the NEQ predicate on the "round_count" field.
func RoundCountNEQ(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNEQ(FieldRoundCount, v))
}

// RoundCountIn applies the In predicate on the "round_count" field.
func RoundCountIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIn(FieldRoundCount, vs...))
}

// RoundCountNotIn applies the NotIn predicate on the "round_count" field.
func RoundCountNotIn(vs ...uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotIn(FieldRoundCount, vs...))
}

// RoundCountGT applies the GT predicate on the "round_count" field.
func RoundCountGT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGT(FieldRoundCount, v))
}

// RoundCountGTE applies the GTE predicate on the "round_count" field.
func RoundCountGTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldGTE(FieldRoundCount, v))
}

// RoundCountLT applies the LT predicate on the "round_count" field.
func RoundCountLT(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLT(FieldRoundCount, v))
}

// RoundCountLTE applies the LTE predicate on the "round_count" field.
func RoundCountLTE(v uint32) predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldLTE(FieldRoundCount, v))
}

// RoundCountIsNil applies the IsNil predicate on the "round_count" field.
func RoundCountIsNil() predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldIsNull(FieldRoundCount))
}

// RoundCountNotNil applies the NotNil predicate on the "round_count" field.
func RoundCountNotNil() predicate.RoundInvest {
	return predicate.RoundInvest(sql.FieldNotNull(FieldRoundCount))
}

// HasRound applies the HasEdge predicate on the "round" edge.
func HasRound() predicate.RoundInvest {
	return predicate.RoundInvest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoundTable, RoundColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoundWith applies the HasEdge predicate on the "round" edge with a given conditions (other predicates).
func HasRoundWith(preds ...predicate.Round) predicate.RoundInvest {
	return predicate.RoundInvest(func(s *sql.Selector) {
		step := newRoundStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoundInvest) predicate.RoundInvest {
	return predicate.RoundInvest(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoundInvest) predicate.RoundInvest {
	return predicate.RoundInvest(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoundInvest) predicate.RoundInvest {
	return predicate.RoundInvest(sql.NotPredicates(p))
}
