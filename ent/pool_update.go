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
	"github.com/kebin6/wolflamp-rpc/ent/pool"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// PoolUpdate is the builder for updating Pool entities.
type PoolUpdate struct {
	config
	hooks    []Hook
	mutation *PoolMutation
}

// Where appends a list predicates to the PoolUpdate builder.
func (pu *PoolUpdate) Where(ps ...predicate.Pool) *PoolUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PoolUpdate) SetUpdatedAt(t time.Time) *PoolUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PoolUpdate) SetStatus(u uint8) *PoolUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(u)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableStatus(u *uint8) *PoolUpdate {
	if u != nil {
		pu.SetStatus(*u)
	}
	return pu
}

// AddStatus adds u to the "status" field.
func (pu *PoolUpdate) AddStatus(u int8) *PoolUpdate {
	pu.mutation.AddStatus(u)
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *PoolUpdate) ClearStatus() *PoolUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetRoundID sets the "round_id" field.
func (pu *PoolUpdate) SetRoundID(u uint64) *PoolUpdate {
	pu.mutation.ResetRoundID()
	pu.mutation.SetRoundID(u)
	return pu
}

// SetNillableRoundID sets the "round_id" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableRoundID(u *uint64) *PoolUpdate {
	if u != nil {
		pu.SetRoundID(*u)
	}
	return pu
}

// AddRoundID adds u to the "round_id" field.
func (pu *PoolUpdate) AddRoundID(u int64) *PoolUpdate {
	pu.mutation.AddRoundID(u)
	return pu
}

// SetMode sets the "mode" field.
func (pu *PoolUpdate) SetMode(s string) *PoolUpdate {
	pu.mutation.SetMode(s)
	return pu
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableMode(s *string) *PoolUpdate {
	if s != nil {
		pu.SetMode(*s)
	}
	return pu
}

// SetType sets the "type" field.
func (pu *PoolUpdate) SetType(u uint32) *PoolUpdate {
	pu.mutation.ResetType()
	pu.mutation.SetType(u)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableType(u *uint32) *PoolUpdate {
	if u != nil {
		pu.SetType(*u)
	}
	return pu
}

// AddType adds u to the "type" field.
func (pu *PoolUpdate) AddType(u int32) *PoolUpdate {
	pu.mutation.AddType(u)
	return pu
}

// SetLambNum sets the "lamb_num" field.
func (pu *PoolUpdate) SetLambNum(f float64) *PoolUpdate {
	pu.mutation.ResetLambNum()
	pu.mutation.SetLambNum(f)
	return pu
}

// SetNillableLambNum sets the "lamb_num" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableLambNum(f *float64) *PoolUpdate {
	if f != nil {
		pu.SetLambNum(*f)
	}
	return pu
}

// AddLambNum adds f to the "lamb_num" field.
func (pu *PoolUpdate) AddLambNum(f float64) *PoolUpdate {
	pu.mutation.AddLambNum(f)
	return pu
}

// SetRemark sets the "remark" field.
func (pu *PoolUpdate) SetRemark(s string) *PoolUpdate {
	pu.mutation.SetRemark(s)
	return pu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableRemark(s *string) *PoolUpdate {
	if s != nil {
		pu.SetRemark(*s)
	}
	return pu
}

// Mutation returns the PoolMutation object of the builder.
func (pu *PoolUpdate) Mutation() *PoolMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PoolUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PoolUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PoolUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PoolUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PoolUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := pool.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pool.Table, pool.Columns, sqlgraph.NewFieldSpec(pool.FieldID, field.TypeUint64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(pool.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(pool.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(pool.FieldStatus, field.TypeUint8, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(pool.FieldStatus, field.TypeUint8)
	}
	if value, ok := pu.mutation.RoundID(); ok {
		_spec.SetField(pool.FieldRoundID, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.AddedRoundID(); ok {
		_spec.AddField(pool.FieldRoundID, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.Mode(); ok {
		_spec.SetField(pool.FieldMode, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(pool.FieldType, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.AddedType(); ok {
		_spec.AddField(pool.FieldType, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.LambNum(); ok {
		_spec.SetField(pool.FieldLambNum, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedLambNum(); ok {
		_spec.AddField(pool.FieldLambNum, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Remark(); ok {
		_spec.SetField(pool.FieldRemark, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PoolUpdateOne is the builder for updating a single Pool entity.
type PoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PoolMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PoolUpdateOne) SetUpdatedAt(t time.Time) *PoolUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PoolUpdateOne) SetStatus(u uint8) *PoolUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(u)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableStatus(u *uint8) *PoolUpdateOne {
	if u != nil {
		puo.SetStatus(*u)
	}
	return puo
}

// AddStatus adds u to the "status" field.
func (puo *PoolUpdateOne) AddStatus(u int8) *PoolUpdateOne {
	puo.mutation.AddStatus(u)
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *PoolUpdateOne) ClearStatus() *PoolUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetRoundID sets the "round_id" field.
func (puo *PoolUpdateOne) SetRoundID(u uint64) *PoolUpdateOne {
	puo.mutation.ResetRoundID()
	puo.mutation.SetRoundID(u)
	return puo
}

// SetNillableRoundID sets the "round_id" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableRoundID(u *uint64) *PoolUpdateOne {
	if u != nil {
		puo.SetRoundID(*u)
	}
	return puo
}

// AddRoundID adds u to the "round_id" field.
func (puo *PoolUpdateOne) AddRoundID(u int64) *PoolUpdateOne {
	puo.mutation.AddRoundID(u)
	return puo
}

// SetMode sets the "mode" field.
func (puo *PoolUpdateOne) SetMode(s string) *PoolUpdateOne {
	puo.mutation.SetMode(s)
	return puo
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableMode(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetMode(*s)
	}
	return puo
}

// SetType sets the "type" field.
func (puo *PoolUpdateOne) SetType(u uint32) *PoolUpdateOne {
	puo.mutation.ResetType()
	puo.mutation.SetType(u)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableType(u *uint32) *PoolUpdateOne {
	if u != nil {
		puo.SetType(*u)
	}
	return puo
}

// AddType adds u to the "type" field.
func (puo *PoolUpdateOne) AddType(u int32) *PoolUpdateOne {
	puo.mutation.AddType(u)
	return puo
}

// SetLambNum sets the "lamb_num" field.
func (puo *PoolUpdateOne) SetLambNum(f float64) *PoolUpdateOne {
	puo.mutation.ResetLambNum()
	puo.mutation.SetLambNum(f)
	return puo
}

// SetNillableLambNum sets the "lamb_num" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableLambNum(f *float64) *PoolUpdateOne {
	if f != nil {
		puo.SetLambNum(*f)
	}
	return puo
}

// AddLambNum adds f to the "lamb_num" field.
func (puo *PoolUpdateOne) AddLambNum(f float64) *PoolUpdateOne {
	puo.mutation.AddLambNum(f)
	return puo
}

// SetRemark sets the "remark" field.
func (puo *PoolUpdateOne) SetRemark(s string) *PoolUpdateOne {
	puo.mutation.SetRemark(s)
	return puo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableRemark(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetRemark(*s)
	}
	return puo
}

// Mutation returns the PoolMutation object of the builder.
func (puo *PoolUpdateOne) Mutation() *PoolMutation {
	return puo.mutation
}

// Where appends a list predicates to the PoolUpdate builder.
func (puo *PoolUpdateOne) Where(ps ...predicate.Pool) *PoolUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PoolUpdateOne) Select(field string, fields ...string) *PoolUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pool entity.
func (puo *PoolUpdateOne) Save(ctx context.Context) (*Pool, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PoolUpdateOne) SaveX(ctx context.Context) *Pool {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PoolUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PoolUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PoolUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := pool.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PoolUpdateOne) sqlSave(ctx context.Context) (_node *Pool, err error) {
	_spec := sqlgraph.NewUpdateSpec(pool.Table, pool.Columns, sqlgraph.NewFieldSpec(pool.FieldID, field.TypeUint64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for _, f := range fields {
			if !pool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(pool.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(pool.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(pool.FieldStatus, field.TypeUint8, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(pool.FieldStatus, field.TypeUint8)
	}
	if value, ok := puo.mutation.RoundID(); ok {
		_spec.SetField(pool.FieldRoundID, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.AddedRoundID(); ok {
		_spec.AddField(pool.FieldRoundID, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.Mode(); ok {
		_spec.SetField(pool.FieldMode, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(pool.FieldType, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.AddedType(); ok {
		_spec.AddField(pool.FieldType, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.LambNum(); ok {
		_spec.SetField(pool.FieldLambNum, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedLambNum(); ok {
		_spec.AddField(pool.FieldLambNum, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Remark(); ok {
		_spec.SetField(pool.FieldRemark, field.TypeString, value)
	}
	_node = &Pool{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
