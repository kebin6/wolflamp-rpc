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
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ExchangeUpdate is the builder for updating Exchange entities.
type ExchangeUpdate struct {
	config
	hooks    []Hook
	mutation *ExchangeMutation
}

// Where appends a list predicates to the ExchangeUpdate builder.
func (eu *ExchangeUpdate) Where(ps ...predicate.Exchange) *ExchangeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *ExchangeUpdate) SetUpdatedAt(t time.Time) *ExchangeUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetStatus sets the "status" field.
func (eu *ExchangeUpdate) SetStatus(u uint8) *ExchangeUpdate {
	eu.mutation.ResetStatus()
	eu.mutation.SetStatus(u)
	return eu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableStatus(u *uint8) *ExchangeUpdate {
	if u != nil {
		eu.SetStatus(*u)
	}
	return eu
}

// AddStatus adds u to the "status" field.
func (eu *ExchangeUpdate) AddStatus(u int8) *ExchangeUpdate {
	eu.mutation.AddStatus(u)
	return eu
}

// ClearStatus clears the value of the "status" field.
func (eu *ExchangeUpdate) ClearStatus() *ExchangeUpdate {
	eu.mutation.ClearStatus()
	return eu
}

// SetPlayerID sets the "player_id" field.
func (eu *ExchangeUpdate) SetPlayerID(u uint64) *ExchangeUpdate {
	eu.mutation.ResetPlayerID()
	eu.mutation.SetPlayerID(u)
	return eu
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillablePlayerID(u *uint64) *ExchangeUpdate {
	if u != nil {
		eu.SetPlayerID(*u)
	}
	return eu
}

// AddPlayerID adds u to the "player_id" field.
func (eu *ExchangeUpdate) AddPlayerID(u int64) *ExchangeUpdate {
	eu.mutation.AddPlayerID(u)
	return eu
}

// SetTransactionID sets the "transaction_id" field.
func (eu *ExchangeUpdate) SetTransactionID(s string) *ExchangeUpdate {
	eu.mutation.SetTransactionID(s)
	return eu
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableTransactionID(s *string) *ExchangeUpdate {
	if s != nil {
		eu.SetTransactionID(*s)
	}
	return eu
}

// SetMode sets the "mode" field.
func (eu *ExchangeUpdate) SetMode(s string) *ExchangeUpdate {
	eu.mutation.SetMode(s)
	return eu
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableMode(s *string) *ExchangeUpdate {
	if s != nil {
		eu.SetMode(*s)
	}
	return eu
}

// SetType sets the "type" field.
func (eu *ExchangeUpdate) SetType(u uint32) *ExchangeUpdate {
	eu.mutation.ResetType()
	eu.mutation.SetType(u)
	return eu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableType(u *uint32) *ExchangeUpdate {
	if u != nil {
		eu.SetType(*u)
	}
	return eu
}

// AddType adds u to the "type" field.
func (eu *ExchangeUpdate) AddType(u int32) *ExchangeUpdate {
	eu.mutation.AddType(u)
	return eu
}

// SetCoinNum sets the "coin_num" field.
func (eu *ExchangeUpdate) SetCoinNum(u uint32) *ExchangeUpdate {
	eu.mutation.ResetCoinNum()
	eu.mutation.SetCoinNum(u)
	return eu
}

// SetNillableCoinNum sets the "coin_num" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableCoinNum(u *uint32) *ExchangeUpdate {
	if u != nil {
		eu.SetCoinNum(*u)
	}
	return eu
}

// AddCoinNum adds u to the "coin_num" field.
func (eu *ExchangeUpdate) AddCoinNum(u int32) *ExchangeUpdate {
	eu.mutation.AddCoinNum(u)
	return eu
}

// SetLampNum sets the "lamp_num" field.
func (eu *ExchangeUpdate) SetLampNum(u uint32) *ExchangeUpdate {
	eu.mutation.ResetLampNum()
	eu.mutation.SetLampNum(u)
	return eu
}

// SetNillableLampNum sets the "lamp_num" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableLampNum(u *uint32) *ExchangeUpdate {
	if u != nil {
		eu.SetLampNum(*u)
	}
	return eu
}

// AddLampNum adds u to the "lamp_num" field.
func (eu *ExchangeUpdate) AddLampNum(u int32) *ExchangeUpdate {
	eu.mutation.AddLampNum(u)
	return eu
}

// SetGcicsOrderID sets the "gcics_order_id" field.
func (eu *ExchangeUpdate) SetGcicsOrderID(s string) *ExchangeUpdate {
	eu.mutation.SetGcicsOrderID(s)
	return eu
}

// SetNillableGcicsOrderID sets the "gcics_order_id" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableGcicsOrderID(s *string) *ExchangeUpdate {
	if s != nil {
		eu.SetGcicsOrderID(*s)
	}
	return eu
}

// SetRemark sets the "remark" field.
func (eu *ExchangeUpdate) SetRemark(s string) *ExchangeUpdate {
	eu.mutation.SetRemark(s)
	return eu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (eu *ExchangeUpdate) SetNillableRemark(s *string) *ExchangeUpdate {
	if s != nil {
		eu.SetRemark(*s)
	}
	return eu
}

// Mutation returns the ExchangeMutation object of the builder.
func (eu *ExchangeUpdate) Mutation() *ExchangeMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ExchangeUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExchangeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExchangeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExchangeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *ExchangeUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := exchange.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

func (eu *ExchangeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(exchange.Table, exchange.Columns, sqlgraph.NewFieldSpec(exchange.FieldID, field.TypeUint64))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(exchange.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.Status(); ok {
		_spec.SetField(exchange.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := eu.mutation.AddedStatus(); ok {
		_spec.AddField(exchange.FieldStatus, field.TypeUint8, value)
	}
	if eu.mutation.StatusCleared() {
		_spec.ClearField(exchange.FieldStatus, field.TypeUint8)
	}
	if value, ok := eu.mutation.PlayerID(); ok {
		_spec.SetField(exchange.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := eu.mutation.AddedPlayerID(); ok {
		_spec.AddField(exchange.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := eu.mutation.TransactionID(); ok {
		_spec.SetField(exchange.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := eu.mutation.Mode(); ok {
		_spec.SetField(exchange.FieldMode, field.TypeString, value)
	}
	if value, ok := eu.mutation.GetType(); ok {
		_spec.SetField(exchange.FieldType, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.AddedType(); ok {
		_spec.AddField(exchange.FieldType, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.CoinNum(); ok {
		_spec.SetField(exchange.FieldCoinNum, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.AddedCoinNum(); ok {
		_spec.AddField(exchange.FieldCoinNum, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.LampNum(); ok {
		_spec.SetField(exchange.FieldLampNum, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.AddedLampNum(); ok {
		_spec.AddField(exchange.FieldLampNum, field.TypeUint32, value)
	}
	if value, ok := eu.mutation.GcicsOrderID(); ok {
		_spec.SetField(exchange.FieldGcicsOrderID, field.TypeString, value)
	}
	if value, ok := eu.mutation.Remark(); ok {
		_spec.SetField(exchange.FieldRemark, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exchange.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// ExchangeUpdateOne is the builder for updating a single Exchange entity.
type ExchangeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExchangeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *ExchangeUpdateOne) SetUpdatedAt(t time.Time) *ExchangeUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetStatus sets the "status" field.
func (euo *ExchangeUpdateOne) SetStatus(u uint8) *ExchangeUpdateOne {
	euo.mutation.ResetStatus()
	euo.mutation.SetStatus(u)
	return euo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableStatus(u *uint8) *ExchangeUpdateOne {
	if u != nil {
		euo.SetStatus(*u)
	}
	return euo
}

// AddStatus adds u to the "status" field.
func (euo *ExchangeUpdateOne) AddStatus(u int8) *ExchangeUpdateOne {
	euo.mutation.AddStatus(u)
	return euo
}

// ClearStatus clears the value of the "status" field.
func (euo *ExchangeUpdateOne) ClearStatus() *ExchangeUpdateOne {
	euo.mutation.ClearStatus()
	return euo
}

// SetPlayerID sets the "player_id" field.
func (euo *ExchangeUpdateOne) SetPlayerID(u uint64) *ExchangeUpdateOne {
	euo.mutation.ResetPlayerID()
	euo.mutation.SetPlayerID(u)
	return euo
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillablePlayerID(u *uint64) *ExchangeUpdateOne {
	if u != nil {
		euo.SetPlayerID(*u)
	}
	return euo
}

// AddPlayerID adds u to the "player_id" field.
func (euo *ExchangeUpdateOne) AddPlayerID(u int64) *ExchangeUpdateOne {
	euo.mutation.AddPlayerID(u)
	return euo
}

// SetTransactionID sets the "transaction_id" field.
func (euo *ExchangeUpdateOne) SetTransactionID(s string) *ExchangeUpdateOne {
	euo.mutation.SetTransactionID(s)
	return euo
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableTransactionID(s *string) *ExchangeUpdateOne {
	if s != nil {
		euo.SetTransactionID(*s)
	}
	return euo
}

// SetMode sets the "mode" field.
func (euo *ExchangeUpdateOne) SetMode(s string) *ExchangeUpdateOne {
	euo.mutation.SetMode(s)
	return euo
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableMode(s *string) *ExchangeUpdateOne {
	if s != nil {
		euo.SetMode(*s)
	}
	return euo
}

// SetType sets the "type" field.
func (euo *ExchangeUpdateOne) SetType(u uint32) *ExchangeUpdateOne {
	euo.mutation.ResetType()
	euo.mutation.SetType(u)
	return euo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableType(u *uint32) *ExchangeUpdateOne {
	if u != nil {
		euo.SetType(*u)
	}
	return euo
}

// AddType adds u to the "type" field.
func (euo *ExchangeUpdateOne) AddType(u int32) *ExchangeUpdateOne {
	euo.mutation.AddType(u)
	return euo
}

// SetCoinNum sets the "coin_num" field.
func (euo *ExchangeUpdateOne) SetCoinNum(u uint32) *ExchangeUpdateOne {
	euo.mutation.ResetCoinNum()
	euo.mutation.SetCoinNum(u)
	return euo
}

// SetNillableCoinNum sets the "coin_num" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableCoinNum(u *uint32) *ExchangeUpdateOne {
	if u != nil {
		euo.SetCoinNum(*u)
	}
	return euo
}

// AddCoinNum adds u to the "coin_num" field.
func (euo *ExchangeUpdateOne) AddCoinNum(u int32) *ExchangeUpdateOne {
	euo.mutation.AddCoinNum(u)
	return euo
}

// SetLampNum sets the "lamp_num" field.
func (euo *ExchangeUpdateOne) SetLampNum(u uint32) *ExchangeUpdateOne {
	euo.mutation.ResetLampNum()
	euo.mutation.SetLampNum(u)
	return euo
}

// SetNillableLampNum sets the "lamp_num" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableLampNum(u *uint32) *ExchangeUpdateOne {
	if u != nil {
		euo.SetLampNum(*u)
	}
	return euo
}

// AddLampNum adds u to the "lamp_num" field.
func (euo *ExchangeUpdateOne) AddLampNum(u int32) *ExchangeUpdateOne {
	euo.mutation.AddLampNum(u)
	return euo
}

// SetGcicsOrderID sets the "gcics_order_id" field.
func (euo *ExchangeUpdateOne) SetGcicsOrderID(s string) *ExchangeUpdateOne {
	euo.mutation.SetGcicsOrderID(s)
	return euo
}

// SetNillableGcicsOrderID sets the "gcics_order_id" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableGcicsOrderID(s *string) *ExchangeUpdateOne {
	if s != nil {
		euo.SetGcicsOrderID(*s)
	}
	return euo
}

// SetRemark sets the "remark" field.
func (euo *ExchangeUpdateOne) SetRemark(s string) *ExchangeUpdateOne {
	euo.mutation.SetRemark(s)
	return euo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (euo *ExchangeUpdateOne) SetNillableRemark(s *string) *ExchangeUpdateOne {
	if s != nil {
		euo.SetRemark(*s)
	}
	return euo
}

// Mutation returns the ExchangeMutation object of the builder.
func (euo *ExchangeUpdateOne) Mutation() *ExchangeMutation {
	return euo.mutation
}

// Where appends a list predicates to the ExchangeUpdate builder.
func (euo *ExchangeUpdateOne) Where(ps ...predicate.Exchange) *ExchangeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ExchangeUpdateOne) Select(field string, fields ...string) *ExchangeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Exchange entity.
func (euo *ExchangeUpdateOne) Save(ctx context.Context) (*Exchange, error) {
	euo.defaults()
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExchangeUpdateOne) SaveX(ctx context.Context) *Exchange {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ExchangeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExchangeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *ExchangeUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := exchange.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

func (euo *ExchangeUpdateOne) sqlSave(ctx context.Context) (_node *Exchange, err error) {
	_spec := sqlgraph.NewUpdateSpec(exchange.Table, exchange.Columns, sqlgraph.NewFieldSpec(exchange.FieldID, field.TypeUint64))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Exchange.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exchange.FieldID)
		for _, f := range fields {
			if !exchange.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != exchange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(exchange.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.Status(); ok {
		_spec.SetField(exchange.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := euo.mutation.AddedStatus(); ok {
		_spec.AddField(exchange.FieldStatus, field.TypeUint8, value)
	}
	if euo.mutation.StatusCleared() {
		_spec.ClearField(exchange.FieldStatus, field.TypeUint8)
	}
	if value, ok := euo.mutation.PlayerID(); ok {
		_spec.SetField(exchange.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := euo.mutation.AddedPlayerID(); ok {
		_spec.AddField(exchange.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := euo.mutation.TransactionID(); ok {
		_spec.SetField(exchange.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := euo.mutation.Mode(); ok {
		_spec.SetField(exchange.FieldMode, field.TypeString, value)
	}
	if value, ok := euo.mutation.GetType(); ok {
		_spec.SetField(exchange.FieldType, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.AddedType(); ok {
		_spec.AddField(exchange.FieldType, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.CoinNum(); ok {
		_spec.SetField(exchange.FieldCoinNum, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.AddedCoinNum(); ok {
		_spec.AddField(exchange.FieldCoinNum, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.LampNum(); ok {
		_spec.SetField(exchange.FieldLampNum, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.AddedLampNum(); ok {
		_spec.AddField(exchange.FieldLampNum, field.TypeUint32, value)
	}
	if value, ok := euo.mutation.GcicsOrderID(); ok {
		_spec.SetField(exchange.FieldGcicsOrderID, field.TypeString, value)
	}
	if value, ok := euo.mutation.Remark(); ok {
		_spec.SetField(exchange.FieldRemark, field.TypeString, value)
	}
	_node = &Exchange{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exchange.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
