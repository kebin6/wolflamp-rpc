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
	"github.com/kebin6/wolflamp-rpc/ent/order"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(u uint8) *OrderUpdate {
	ou.mutation.ResetStatus()
	ou.mutation.SetStatus(u)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatus(u *uint8) *OrderUpdate {
	if u != nil {
		ou.SetStatus(*u)
	}
	return ou
}

// AddStatus adds u to the "status" field.
func (ou *OrderUpdate) AddStatus(u int8) *OrderUpdate {
	ou.mutation.AddStatus(u)
	return ou
}

// ClearStatus clears the value of the "status" field.
func (ou *OrderUpdate) ClearStatus() *OrderUpdate {
	ou.mutation.ClearStatus()
	return ou
}

// SetPlayerID sets the "player_id" field.
func (ou *OrderUpdate) SetPlayerID(u uint64) *OrderUpdate {
	ou.mutation.ResetPlayerID()
	ou.mutation.SetPlayerID(u)
	return ou
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePlayerID(u *uint64) *OrderUpdate {
	if u != nil {
		ou.SetPlayerID(*u)
	}
	return ou
}

// AddPlayerID adds u to the "player_id" field.
func (ou *OrderUpdate) AddPlayerID(u int64) *OrderUpdate {
	ou.mutation.AddPlayerID(u)
	return ou
}

// SetType sets the "type" field.
func (ou *OrderUpdate) SetType(s string) *OrderUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetType(*s)
	}
	return ou
}

// SetCode sets the "code" field.
func (ou *OrderUpdate) SetCode(s string) *OrderUpdate {
	ou.mutation.SetCode(s)
	return ou
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCode(s *string) *OrderUpdate {
	if s != nil {
		ou.SetCode(*s)
	}
	return ou
}

// SetTransactionID sets the "transaction_id" field.
func (ou *OrderUpdate) SetTransactionID(s string) *OrderUpdate {
	ou.mutation.SetTransactionID(s)
	return ou
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableTransactionID(s *string) *OrderUpdate {
	if s != nil {
		ou.SetTransactionID(*s)
	}
	return ou
}

// SetFromAddress sets the "from_address" field.
func (ou *OrderUpdate) SetFromAddress(s string) *OrderUpdate {
	ou.mutation.SetFromAddress(s)
	return ou
}

// SetNillableFromAddress sets the "from_address" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableFromAddress(s *string) *OrderUpdate {
	if s != nil {
		ou.SetFromAddress(*s)
	}
	return ou
}

// SetToAddress sets the "to_address" field.
func (ou *OrderUpdate) SetToAddress(s string) *OrderUpdate {
	ou.mutation.SetToAddress(s)
	return ou
}

// SetNillableToAddress sets the "to_address" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableToAddress(s *string) *OrderUpdate {
	if s != nil {
		ou.SetToAddress(*s)
	}
	return ou
}

// SetNum sets the "num" field.
func (ou *OrderUpdate) SetNum(f float64) *OrderUpdate {
	ou.mutation.ResetNum()
	ou.mutation.SetNum(f)
	return ou
}

// SetNillableNum sets the "num" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableNum(f *float64) *OrderUpdate {
	if f != nil {
		ou.SetNum(*f)
	}
	return ou
}

// AddNum adds f to the "num" field.
func (ou *OrderUpdate) AddNum(f float64) *OrderUpdate {
	ou.mutation.AddNum(f)
	return ou
}

// SetHandlingFee sets the "handling_fee" field.
func (ou *OrderUpdate) SetHandlingFee(f float64) *OrderUpdate {
	ou.mutation.ResetHandlingFee()
	ou.mutation.SetHandlingFee(f)
	return ou
}

// SetNillableHandlingFee sets the "handling_fee" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableHandlingFee(f *float64) *OrderUpdate {
	if f != nil {
		ou.SetHandlingFee(*f)
	}
	return ou
}

// AddHandlingFee adds f to the "handling_fee" field.
func (ou *OrderUpdate) AddHandlingFee(f float64) *OrderUpdate {
	ou.mutation.AddHandlingFee(f)
	return ou
}

// SetNetwork sets the "network" field.
func (ou *OrderUpdate) SetNetwork(s string) *OrderUpdate {
	ou.mutation.SetNetwork(s)
	return ou
}

// SetNillableNetwork sets the "network" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableNetwork(s *string) *OrderUpdate {
	if s != nil {
		ou.SetNetwork(*s)
	}
	return ou
}

// SetRemark sets the "remark" field.
func (ou *OrderUpdate) SetRemark(s string) *OrderUpdate {
	ou.mutation.SetRemark(s)
	return ou
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableRemark(s *string) *OrderUpdate {
	if s != nil {
		ou.SetRemark(*s)
	}
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUint64))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ou.mutation.AddedStatus(); ok {
		_spec.AddField(order.FieldStatus, field.TypeUint8, value)
	}
	if ou.mutation.StatusCleared() {
		_spec.ClearField(order.FieldStatus, field.TypeUint8)
	}
	if value, ok := ou.mutation.PlayerID(); ok {
		_spec.SetField(order.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := ou.mutation.AddedPlayerID(); ok {
		_spec.AddField(order.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.SetField(order.FieldType, field.TypeString, value)
	}
	if value, ok := ou.mutation.Code(); ok {
		_spec.SetField(order.FieldCode, field.TypeString, value)
	}
	if value, ok := ou.mutation.TransactionID(); ok {
		_spec.SetField(order.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := ou.mutation.FromAddress(); ok {
		_spec.SetField(order.FieldFromAddress, field.TypeString, value)
	}
	if value, ok := ou.mutation.ToAddress(); ok {
		_spec.SetField(order.FieldToAddress, field.TypeString, value)
	}
	if value, ok := ou.mutation.Num(); ok {
		_spec.SetField(order.FieldNum, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.AddedNum(); ok {
		_spec.AddField(order.FieldNum, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.HandlingFee(); ok {
		_spec.SetField(order.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.AddedHandlingFee(); ok {
		_spec.AddField(order.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.Network(); ok {
		_spec.SetField(order.FieldNetwork, field.TypeString, value)
	}
	if value, ok := ou.mutation.Remark(); ok {
		_spec.SetField(order.FieldRemark, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(u uint8) *OrderUpdateOne {
	ouo.mutation.ResetStatus()
	ouo.mutation.SetStatus(u)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatus(u *uint8) *OrderUpdateOne {
	if u != nil {
		ouo.SetStatus(*u)
	}
	return ouo
}

// AddStatus adds u to the "status" field.
func (ouo *OrderUpdateOne) AddStatus(u int8) *OrderUpdateOne {
	ouo.mutation.AddStatus(u)
	return ouo
}

// ClearStatus clears the value of the "status" field.
func (ouo *OrderUpdateOne) ClearStatus() *OrderUpdateOne {
	ouo.mutation.ClearStatus()
	return ouo
}

// SetPlayerID sets the "player_id" field.
func (ouo *OrderUpdateOne) SetPlayerID(u uint64) *OrderUpdateOne {
	ouo.mutation.ResetPlayerID()
	ouo.mutation.SetPlayerID(u)
	return ouo
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePlayerID(u *uint64) *OrderUpdateOne {
	if u != nil {
		ouo.SetPlayerID(*u)
	}
	return ouo
}

// AddPlayerID adds u to the "player_id" field.
func (ouo *OrderUpdateOne) AddPlayerID(u int64) *OrderUpdateOne {
	ouo.mutation.AddPlayerID(u)
	return ouo
}

// SetType sets the "type" field.
func (ouo *OrderUpdateOne) SetType(s string) *OrderUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetType(*s)
	}
	return ouo
}

// SetCode sets the "code" field.
func (ouo *OrderUpdateOne) SetCode(s string) *OrderUpdateOne {
	ouo.mutation.SetCode(s)
	return ouo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCode(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetCode(*s)
	}
	return ouo
}

// SetTransactionID sets the "transaction_id" field.
func (ouo *OrderUpdateOne) SetTransactionID(s string) *OrderUpdateOne {
	ouo.mutation.SetTransactionID(s)
	return ouo
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableTransactionID(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetTransactionID(*s)
	}
	return ouo
}

// SetFromAddress sets the "from_address" field.
func (ouo *OrderUpdateOne) SetFromAddress(s string) *OrderUpdateOne {
	ouo.mutation.SetFromAddress(s)
	return ouo
}

// SetNillableFromAddress sets the "from_address" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableFromAddress(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetFromAddress(*s)
	}
	return ouo
}

// SetToAddress sets the "to_address" field.
func (ouo *OrderUpdateOne) SetToAddress(s string) *OrderUpdateOne {
	ouo.mutation.SetToAddress(s)
	return ouo
}

// SetNillableToAddress sets the "to_address" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableToAddress(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetToAddress(*s)
	}
	return ouo
}

// SetNum sets the "num" field.
func (ouo *OrderUpdateOne) SetNum(f float64) *OrderUpdateOne {
	ouo.mutation.ResetNum()
	ouo.mutation.SetNum(f)
	return ouo
}

// SetNillableNum sets the "num" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableNum(f *float64) *OrderUpdateOne {
	if f != nil {
		ouo.SetNum(*f)
	}
	return ouo
}

// AddNum adds f to the "num" field.
func (ouo *OrderUpdateOne) AddNum(f float64) *OrderUpdateOne {
	ouo.mutation.AddNum(f)
	return ouo
}

// SetHandlingFee sets the "handling_fee" field.
func (ouo *OrderUpdateOne) SetHandlingFee(f float64) *OrderUpdateOne {
	ouo.mutation.ResetHandlingFee()
	ouo.mutation.SetHandlingFee(f)
	return ouo
}

// SetNillableHandlingFee sets the "handling_fee" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableHandlingFee(f *float64) *OrderUpdateOne {
	if f != nil {
		ouo.SetHandlingFee(*f)
	}
	return ouo
}

// AddHandlingFee adds f to the "handling_fee" field.
func (ouo *OrderUpdateOne) AddHandlingFee(f float64) *OrderUpdateOne {
	ouo.mutation.AddHandlingFee(f)
	return ouo
}

// SetNetwork sets the "network" field.
func (ouo *OrderUpdateOne) SetNetwork(s string) *OrderUpdateOne {
	ouo.mutation.SetNetwork(s)
	return ouo
}

// SetNillableNetwork sets the "network" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableNetwork(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetNetwork(*s)
	}
	return ouo
}

// SetRemark sets the "remark" field.
func (ouo *OrderUpdateOne) SetRemark(s string) *OrderUpdateOne {
	ouo.mutation.SetRemark(s)
	return ouo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableRemark(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetRemark(*s)
	}
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	ouo.defaults()
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUint64))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ouo.mutation.AddedStatus(); ok {
		_spec.AddField(order.FieldStatus, field.TypeUint8, value)
	}
	if ouo.mutation.StatusCleared() {
		_spec.ClearField(order.FieldStatus, field.TypeUint8)
	}
	if value, ok := ouo.mutation.PlayerID(); ok {
		_spec.SetField(order.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := ouo.mutation.AddedPlayerID(); ok {
		_spec.AddField(order.FieldPlayerID, field.TypeUint64, value)
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.SetField(order.FieldType, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Code(); ok {
		_spec.SetField(order.FieldCode, field.TypeString, value)
	}
	if value, ok := ouo.mutation.TransactionID(); ok {
		_spec.SetField(order.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := ouo.mutation.FromAddress(); ok {
		_spec.SetField(order.FieldFromAddress, field.TypeString, value)
	}
	if value, ok := ouo.mutation.ToAddress(); ok {
		_spec.SetField(order.FieldToAddress, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Num(); ok {
		_spec.SetField(order.FieldNum, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.AddedNum(); ok {
		_spec.AddField(order.FieldNum, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.HandlingFee(); ok {
		_spec.SetField(order.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.AddedHandlingFee(); ok {
		_spec.AddField(order.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.Network(); ok {
		_spec.SetField(order.FieldNetwork, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Remark(); ok {
		_spec.SetField(order.FieldRemark, field.TypeString, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
