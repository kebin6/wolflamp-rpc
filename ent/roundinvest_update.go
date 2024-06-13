// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
)

// RoundInvestUpdate is the builder for updating RoundInvest entities.
type RoundInvestUpdate struct {
	config
	hooks    []Hook
	mutation *RoundInvestMutation
}

// Where appends a list predicates to the RoundInvestUpdate builder.
func (riu *RoundInvestUpdate) Where(ps ...predicate.RoundInvest) *RoundInvestUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// SetUpdatedAt sets the "updated_at" field.
func (riu *RoundInvestUpdate) SetUpdatedAt(t time.Time) *RoundInvestUpdate {
	riu.mutation.SetUpdatedAt(t)
	return riu
}

// SetPlayerID sets the "player_id" field.
func (riu *RoundInvestUpdate) SetPlayerID(u uint64) *RoundInvestUpdate {
	riu.mutation.ResetPlayerID()
	riu.mutation.SetPlayerID(u)
	return riu
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillablePlayerID(u *uint64) *RoundInvestUpdate {
	if u != nil {
		riu.SetPlayerID(*u)
	}
	return riu
}

// AddPlayerID adds u to the "player_id" field.
func (riu *RoundInvestUpdate) AddPlayerID(u int64) *RoundInvestUpdate {
	riu.mutation.AddPlayerID(u)
	return riu
}

// SetPlayerEmail sets the "player_email" field.
func (riu *RoundInvestUpdate) SetPlayerEmail(s string) *RoundInvestUpdate {
	riu.mutation.SetPlayerEmail(s)
	return riu
}

// SetNillablePlayerEmail sets the "player_email" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillablePlayerEmail(s *string) *RoundInvestUpdate {
	if s != nil {
		riu.SetPlayerEmail(*s)
	}
	return riu
}

// SetFoldNo sets the "fold_no" field.
func (riu *RoundInvestUpdate) SetFoldNo(u uint32) *RoundInvestUpdate {
	riu.mutation.ResetFoldNo()
	riu.mutation.SetFoldNo(u)
	return riu
}

// SetNillableFoldNo sets the "fold_no" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillableFoldNo(u *uint32) *RoundInvestUpdate {
	if u != nil {
		riu.SetFoldNo(*u)
	}
	return riu
}

// AddFoldNo adds u to the "fold_no" field.
func (riu *RoundInvestUpdate) AddFoldNo(u int32) *RoundInvestUpdate {
	riu.mutation.AddFoldNo(u)
	return riu
}

// SetLambNum sets the "lamb_num" field.
func (riu *RoundInvestUpdate) SetLambNum(u uint32) *RoundInvestUpdate {
	riu.mutation.ResetLambNum()
	riu.mutation.SetLambNum(u)
	return riu
}

// SetNillableLambNum sets the "lamb_num" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillableLambNum(u *uint32) *RoundInvestUpdate {
	if u != nil {
		riu.SetLambNum(*u)
	}
	return riu
}

// AddLambNum adds u to the "lamb_num" field.
func (riu *RoundInvestUpdate) AddLambNum(u int32) *RoundInvestUpdate {
	riu.mutation.AddLambNum(u)
	return riu
}

// SetProfitAndLoss sets the "profit_and_loss" field.
func (riu *RoundInvestUpdate) SetProfitAndLoss(f float32) *RoundInvestUpdate {
	riu.mutation.ResetProfitAndLoss()
	riu.mutation.SetProfitAndLoss(f)
	return riu
}

// SetNillableProfitAndLoss sets the "profit_and_loss" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillableProfitAndLoss(f *float32) *RoundInvestUpdate {
	if f != nil {
		riu.SetProfitAndLoss(*f)
	}
	return riu
}

// AddProfitAndLoss adds f to the "profit_and_loss" field.
func (riu *RoundInvestUpdate) AddProfitAndLoss(f float32) *RoundInvestUpdate {
	riu.mutation.AddProfitAndLoss(f)
	return riu
}

// SetRoundID sets the "round_id" field.
func (riu *RoundInvestUpdate) SetRoundID(u uint64) *RoundInvestUpdate {
	riu.mutation.SetRoundID(u)
	return riu
}

// SetNillableRoundID sets the "round_id" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillableRoundID(u *uint64) *RoundInvestUpdate {
	if u != nil {
		riu.SetRoundID(*u)
	}
	return riu
}

// ClearRoundID clears the value of the "round_id" field.
func (riu *RoundInvestUpdate) ClearRoundID() *RoundInvestUpdate {
	riu.mutation.ClearRoundID()
	return riu
}

// SetRoundCount sets the "round_count" field.
func (riu *RoundInvestUpdate) SetRoundCount(u uint32) *RoundInvestUpdate {
	riu.mutation.ResetRoundCount()
	riu.mutation.SetRoundCount(u)
	return riu
}

// SetNillableRoundCount sets the "round_count" field if the given value is not nil.
func (riu *RoundInvestUpdate) SetNillableRoundCount(u *uint32) *RoundInvestUpdate {
	if u != nil {
		riu.SetRoundCount(*u)
	}
	return riu
}

// AddRoundCount adds u to the "round_count" field.
func (riu *RoundInvestUpdate) AddRoundCount(u int32) *RoundInvestUpdate {
	riu.mutation.AddRoundCount(u)
	return riu
}

// ClearRoundCount clears the value of the "round_count" field.
func (riu *RoundInvestUpdate) ClearRoundCount() *RoundInvestUpdate {
	riu.mutation.ClearRoundCount()
	return riu
}

// SetRound sets the "round" edge to the Round entity.
func (riu *RoundInvestUpdate) SetRound(r *Round) *RoundInvestUpdate {
	return riu.SetRoundID(r.ID)
}

// Mutation returns the RoundInvestMutation object of the builder.
func (riu *RoundInvestUpdate) Mutation() *RoundInvestMutation {
	return riu.mutation
}

// ClearRound clears the "round" edge to the Round entity.
func (riu *RoundInvestUpdate) ClearRound() *RoundInvestUpdate {
	riu.mutation.ClearRound()
	return riu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *RoundInvestUpdate) Save(ctx context.Context) (int, error) {
	riu.defaults()
	return withHooks(ctx, riu.sqlSave, riu.mutation, riu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riu *RoundInvestUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *RoundInvestUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *RoundInvestUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (riu *RoundInvestUpdate) defaults() {
	if _, ok := riu.mutation.UpdatedAt(); !ok {
		v := roundinvest.UpdateDefaultUpdatedAt()
		riu.mutation.SetUpdatedAt(v)
	}
}

func (riu *RoundInvestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(roundinvest.Table, roundinvest.Columns, sqlgraph.NewFieldSpec(roundinvest.FieldID, field.TypeUint64))
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riu.mutation.UpdatedAt(); ok {
		_spec.SetField(roundinvest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := riu.mutation.PlayerID(); ok {
		_spec.SetField(roundinvest.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := riu.mutation.AddedPlayerID(); ok {
		_spec.AddField(roundinvest.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := riu.mutation.PlayerEmail(); ok {
		_spec.SetField(roundinvest.FieldPlayerEmail, field.TypeString, value)
	}
	if value, ok := riu.mutation.FoldNo(); ok {
		_spec.SetField(roundinvest.FieldFoldNo, field.TypeUint32, value)
	}
	if value, ok := riu.mutation.AddedFoldNo(); ok {
		_spec.AddField(roundinvest.FieldFoldNo, field.TypeUint32, value)
	}
	if value, ok := riu.mutation.LambNum(); ok {
		_spec.SetField(roundinvest.FieldLambNum, field.TypeUint32, value)
	}
	if value, ok := riu.mutation.AddedLambNum(); ok {
		_spec.AddField(roundinvest.FieldLambNum, field.TypeUint32, value)
	}
	if value, ok := riu.mutation.ProfitAndLoss(); ok {
		_spec.SetField(roundinvest.FieldProfitAndLoss, field.TypeFloat32, value)
	}
	if value, ok := riu.mutation.AddedProfitAndLoss(); ok {
		_spec.AddField(roundinvest.FieldProfitAndLoss, field.TypeFloat32, value)
	}
	if value, ok := riu.mutation.RoundCount(); ok {
		_spec.SetField(roundinvest.FieldRoundCount, field.TypeUint32, value)
	}
	if value, ok := riu.mutation.AddedRoundCount(); ok {
		_spec.AddField(roundinvest.FieldRoundCount, field.TypeUint32, value)
	}
	if riu.mutation.RoundCountCleared() {
		_spec.ClearField(roundinvest.FieldRoundCount, field.TypeUint32)
	}
	if riu.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roundinvest.RoundTable,
			Columns: []string{roundinvest.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := riu.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roundinvest.RoundTable,
			Columns: []string{roundinvest.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roundinvest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	riu.mutation.done = true
	return n, nil
}

// RoundInvestUpdateOne is the builder for updating a single RoundInvest entity.
type RoundInvestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoundInvestMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (riuo *RoundInvestUpdateOne) SetUpdatedAt(t time.Time) *RoundInvestUpdateOne {
	riuo.mutation.SetUpdatedAt(t)
	return riuo
}

// SetPlayerID sets the "player_id" field.
func (riuo *RoundInvestUpdateOne) SetPlayerID(u uint64) *RoundInvestUpdateOne {
	riuo.mutation.ResetPlayerID()
	riuo.mutation.SetPlayerID(u)
	return riuo
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillablePlayerID(u *uint64) *RoundInvestUpdateOne {
	if u != nil {
		riuo.SetPlayerID(*u)
	}
	return riuo
}

// AddPlayerID adds u to the "player_id" field.
func (riuo *RoundInvestUpdateOne) AddPlayerID(u int64) *RoundInvestUpdateOne {
	riuo.mutation.AddPlayerID(u)
	return riuo
}

// SetPlayerEmail sets the "player_email" field.
func (riuo *RoundInvestUpdateOne) SetPlayerEmail(s string) *RoundInvestUpdateOne {
	riuo.mutation.SetPlayerEmail(s)
	return riuo
}

// SetNillablePlayerEmail sets the "player_email" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillablePlayerEmail(s *string) *RoundInvestUpdateOne {
	if s != nil {
		riuo.SetPlayerEmail(*s)
	}
	return riuo
}

// SetFoldNo sets the "fold_no" field.
func (riuo *RoundInvestUpdateOne) SetFoldNo(u uint32) *RoundInvestUpdateOne {
	riuo.mutation.ResetFoldNo()
	riuo.mutation.SetFoldNo(u)
	return riuo
}

// SetNillableFoldNo sets the "fold_no" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillableFoldNo(u *uint32) *RoundInvestUpdateOne {
	if u != nil {
		riuo.SetFoldNo(*u)
	}
	return riuo
}

// AddFoldNo adds u to the "fold_no" field.
func (riuo *RoundInvestUpdateOne) AddFoldNo(u int32) *RoundInvestUpdateOne {
	riuo.mutation.AddFoldNo(u)
	return riuo
}

// SetLambNum sets the "lamb_num" field.
func (riuo *RoundInvestUpdateOne) SetLambNum(u uint32) *RoundInvestUpdateOne {
	riuo.mutation.ResetLambNum()
	riuo.mutation.SetLambNum(u)
	return riuo
}

// SetNillableLambNum sets the "lamb_num" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillableLambNum(u *uint32) *RoundInvestUpdateOne {
	if u != nil {
		riuo.SetLambNum(*u)
	}
	return riuo
}

// AddLambNum adds u to the "lamb_num" field.
func (riuo *RoundInvestUpdateOne) AddLambNum(u int32) *RoundInvestUpdateOne {
	riuo.mutation.AddLambNum(u)
	return riuo
}

// SetProfitAndLoss sets the "profit_and_loss" field.
func (riuo *RoundInvestUpdateOne) SetProfitAndLoss(f float32) *RoundInvestUpdateOne {
	riuo.mutation.ResetProfitAndLoss()
	riuo.mutation.SetProfitAndLoss(f)
	return riuo
}

// SetNillableProfitAndLoss sets the "profit_and_loss" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillableProfitAndLoss(f *float32) *RoundInvestUpdateOne {
	if f != nil {
		riuo.SetProfitAndLoss(*f)
	}
	return riuo
}

// AddProfitAndLoss adds f to the "profit_and_loss" field.
func (riuo *RoundInvestUpdateOne) AddProfitAndLoss(f float32) *RoundInvestUpdateOne {
	riuo.mutation.AddProfitAndLoss(f)
	return riuo
}

// SetRoundID sets the "round_id" field.
func (riuo *RoundInvestUpdateOne) SetRoundID(u uint64) *RoundInvestUpdateOne {
	riuo.mutation.SetRoundID(u)
	return riuo
}

// SetNillableRoundID sets the "round_id" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillableRoundID(u *uint64) *RoundInvestUpdateOne {
	if u != nil {
		riuo.SetRoundID(*u)
	}
	return riuo
}

// ClearRoundID clears the value of the "round_id" field.
func (riuo *RoundInvestUpdateOne) ClearRoundID() *RoundInvestUpdateOne {
	riuo.mutation.ClearRoundID()
	return riuo
}

// SetRoundCount sets the "round_count" field.
func (riuo *RoundInvestUpdateOne) SetRoundCount(u uint32) *RoundInvestUpdateOne {
	riuo.mutation.ResetRoundCount()
	riuo.mutation.SetRoundCount(u)
	return riuo
}

// SetNillableRoundCount sets the "round_count" field if the given value is not nil.
func (riuo *RoundInvestUpdateOne) SetNillableRoundCount(u *uint32) *RoundInvestUpdateOne {
	if u != nil {
		riuo.SetRoundCount(*u)
	}
	return riuo
}

// AddRoundCount adds u to the "round_count" field.
func (riuo *RoundInvestUpdateOne) AddRoundCount(u int32) *RoundInvestUpdateOne {
	riuo.mutation.AddRoundCount(u)
	return riuo
}

// ClearRoundCount clears the value of the "round_count" field.
func (riuo *RoundInvestUpdateOne) ClearRoundCount() *RoundInvestUpdateOne {
	riuo.mutation.ClearRoundCount()
	return riuo
}

// SetRound sets the "round" edge to the Round entity.
func (riuo *RoundInvestUpdateOne) SetRound(r *Round) *RoundInvestUpdateOne {
	return riuo.SetRoundID(r.ID)
}

// Mutation returns the RoundInvestMutation object of the builder.
func (riuo *RoundInvestUpdateOne) Mutation() *RoundInvestMutation {
	return riuo.mutation
}

// ClearRound clears the "round" edge to the Round entity.
func (riuo *RoundInvestUpdateOne) ClearRound() *RoundInvestUpdateOne {
	riuo.mutation.ClearRound()
	return riuo
}

// Where appends a list predicates to the RoundInvestUpdate builder.
func (riuo *RoundInvestUpdateOne) Where(ps ...predicate.RoundInvest) *RoundInvestUpdateOne {
	riuo.mutation.Where(ps...)
	return riuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *RoundInvestUpdateOne) Select(field string, fields ...string) *RoundInvestUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated RoundInvest entity.
func (riuo *RoundInvestUpdateOne) Save(ctx context.Context) (*RoundInvest, error) {
	riuo.defaults()
	return withHooks(ctx, riuo.sqlSave, riuo.mutation, riuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *RoundInvestUpdateOne) SaveX(ctx context.Context) *RoundInvest {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *RoundInvestUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *RoundInvestUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (riuo *RoundInvestUpdateOne) defaults() {
	if _, ok := riuo.mutation.UpdatedAt(); !ok {
		v := roundinvest.UpdateDefaultUpdatedAt()
		riuo.mutation.SetUpdatedAt(v)
	}
}

func (riuo *RoundInvestUpdateOne) sqlSave(ctx context.Context) (_node *RoundInvest, err error) {
	_spec := sqlgraph.NewUpdateSpec(roundinvest.Table, roundinvest.Columns, sqlgraph.NewFieldSpec(roundinvest.FieldID, field.TypeUint64))
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoundInvest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roundinvest.FieldID)
		for _, f := range fields {
			if !roundinvest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roundinvest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riuo.mutation.UpdatedAt(); ok {
		_spec.SetField(roundinvest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := riuo.mutation.PlayerID(); ok {
		_spec.SetField(roundinvest.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := riuo.mutation.AddedPlayerID(); ok {
		_spec.AddField(roundinvest.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := riuo.mutation.PlayerEmail(); ok {
		_spec.SetField(roundinvest.FieldPlayerEmail, field.TypeString, value)
	}
	if value, ok := riuo.mutation.FoldNo(); ok {
		_spec.SetField(roundinvest.FieldFoldNo, field.TypeUint32, value)
	}
	if value, ok := riuo.mutation.AddedFoldNo(); ok {
		_spec.AddField(roundinvest.FieldFoldNo, field.TypeUint32, value)
	}
	if value, ok := riuo.mutation.LambNum(); ok {
		_spec.SetField(roundinvest.FieldLambNum, field.TypeUint32, value)
	}
	if value, ok := riuo.mutation.AddedLambNum(); ok {
		_spec.AddField(roundinvest.FieldLambNum, field.TypeUint32, value)
	}
	if value, ok := riuo.mutation.ProfitAndLoss(); ok {
		_spec.SetField(roundinvest.FieldProfitAndLoss, field.TypeFloat32, value)
	}
	if value, ok := riuo.mutation.AddedProfitAndLoss(); ok {
		_spec.AddField(roundinvest.FieldProfitAndLoss, field.TypeFloat32, value)
	}
	if value, ok := riuo.mutation.RoundCount(); ok {
		_spec.SetField(roundinvest.FieldRoundCount, field.TypeUint32, value)
	}
	if value, ok := riuo.mutation.AddedRoundCount(); ok {
		_spec.AddField(roundinvest.FieldRoundCount, field.TypeUint32, value)
	}
	if riuo.mutation.RoundCountCleared() {
		_spec.ClearField(roundinvest.FieldRoundCount, field.TypeUint32)
	}
	if riuo.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roundinvest.RoundTable,
			Columns: []string{roundinvest.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := riuo.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roundinvest.RoundTable,
			Columns: []string{roundinvest.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RoundInvest{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roundinvest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	riuo.mutation.done = true
	return _node, nil
}
