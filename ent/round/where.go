// Code generated by ent, DO NOT EDIT.

package round

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldStatus, v))
}

// RoundCount applies equality check predicate on the "round_count" field. It's identical to RoundCountEQ.
func RoundCount(v uint32) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldRoundCount, v))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldStartAt, v))
}

// OpenAt applies equality check predicate on the "open_at" field. It's identical to OpenAtEQ.
func OpenAt(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldOpenAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldEndAt, v))
}

// SelectedFold applies equality check predicate on the "selected_fold" field. It's identical to SelectedFoldEQ.
func SelectedFold(v uint32) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldSelectedFold, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Round {
	return predicate.Round(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Round {
	return predicate.Round(sql.FieldNotNull(FieldStatus))
}

// RoundCountEQ applies the EQ predicate on the "round_count" field.
func RoundCountEQ(v uint32) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldRoundCount, v))
}

// RoundCountNEQ applies the NEQ predicate on the "round_count" field.
func RoundCountNEQ(v uint32) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldRoundCount, v))
}

// RoundCountIn applies the In predicate on the "round_count" field.
func RoundCountIn(vs ...uint32) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldRoundCount, vs...))
}

// RoundCountNotIn applies the NotIn predicate on the "round_count" field.
func RoundCountNotIn(vs ...uint32) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldRoundCount, vs...))
}

// RoundCountGT applies the GT predicate on the "round_count" field.
func RoundCountGT(v uint32) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldRoundCount, v))
}

// RoundCountGTE applies the GTE predicate on the "round_count" field.
func RoundCountGTE(v uint32) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldRoundCount, v))
}

// RoundCountLT applies the LT predicate on the "round_count" field.
func RoundCountLT(v uint32) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldRoundCount, v))
}

// RoundCountLTE applies the LTE predicate on the "round_count" field.
func RoundCountLTE(v uint32) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldRoundCount, v))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldStartAt, v))
}

// OpenAtEQ applies the EQ predicate on the "open_at" field.
func OpenAtEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldOpenAt, v))
}

// OpenAtNEQ applies the NEQ predicate on the "open_at" field.
func OpenAtNEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldOpenAt, v))
}

// OpenAtIn applies the In predicate on the "open_at" field.
func OpenAtIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldOpenAt, vs...))
}

// OpenAtNotIn applies the NotIn predicate on the "open_at" field.
func OpenAtNotIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldOpenAt, vs...))
}

// OpenAtGT applies the GT predicate on the "open_at" field.
func OpenAtGT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldOpenAt, v))
}

// OpenAtGTE applies the GTE predicate on the "open_at" field.
func OpenAtGTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldOpenAt, v))
}

// OpenAtLT applies the LT predicate on the "open_at" field.
func OpenAtLT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldOpenAt, v))
}

// OpenAtLTE applies the LTE predicate on the "open_at" field.
func OpenAtLTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldOpenAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldEndAt, v))
}

// SelectedFoldEQ applies the EQ predicate on the "selected_fold" field.
func SelectedFoldEQ(v uint32) predicate.Round {
	return predicate.Round(sql.FieldEQ(FieldSelectedFold, v))
}

// SelectedFoldNEQ applies the NEQ predicate on the "selected_fold" field.
func SelectedFoldNEQ(v uint32) predicate.Round {
	return predicate.Round(sql.FieldNEQ(FieldSelectedFold, v))
}

// SelectedFoldIn applies the In predicate on the "selected_fold" field.
func SelectedFoldIn(vs ...uint32) predicate.Round {
	return predicate.Round(sql.FieldIn(FieldSelectedFold, vs...))
}

// SelectedFoldNotIn applies the NotIn predicate on the "selected_fold" field.
func SelectedFoldNotIn(vs ...uint32) predicate.Round {
	return predicate.Round(sql.FieldNotIn(FieldSelectedFold, vs...))
}

// SelectedFoldGT applies the GT predicate on the "selected_fold" field.
func SelectedFoldGT(v uint32) predicate.Round {
	return predicate.Round(sql.FieldGT(FieldSelectedFold, v))
}

// SelectedFoldGTE applies the GTE predicate on the "selected_fold" field.
func SelectedFoldGTE(v uint32) predicate.Round {
	return predicate.Round(sql.FieldGTE(FieldSelectedFold, v))
}

// SelectedFoldLT applies the LT predicate on the "selected_fold" field.
func SelectedFoldLT(v uint32) predicate.Round {
	return predicate.Round(sql.FieldLT(FieldSelectedFold, v))
}

// SelectedFoldLTE applies the LTE predicate on the "selected_fold" field.
func SelectedFoldLTE(v uint32) predicate.Round {
	return predicate.Round(sql.FieldLTE(FieldSelectedFold, v))
}

// HasFold applies the HasEdge predicate on the "fold" edge.
func HasFold() predicate.Round {
	return predicate.Round(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FoldTable, FoldColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFoldWith applies the HasEdge predicate on the "fold" edge with a given conditions (other predicates).
func HasFoldWith(preds ...predicate.RoundLambFold) predicate.Round {
	return predicate.Round(func(s *sql.Selector) {
		step := newFoldStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInvest applies the HasEdge predicate on the "invest" edge.
func HasInvest() predicate.Round {
	return predicate.Round(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InvestTable, InvestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvestWith applies the HasEdge predicate on the "invest" edge with a given conditions (other predicates).
func HasInvestWith(preds ...predicate.RoundInvest) predicate.Round {
	return predicate.Round(func(s *sql.Selector) {
		step := newInvestStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Round) predicate.Round {
	return predicate.Round(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Round) predicate.Round {
	return predicate.Round(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Round) predicate.Round {
	return predicate.Round(sql.NotPredicates(p))
}
