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
	"github.com/kebin6/wolflamp-rpc/ent/origininvitecode"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// OriginInviteCodeUpdate is the builder for updating OriginInviteCode entities.
type OriginInviteCodeUpdate struct {
	config
	hooks    []Hook
	mutation *OriginInviteCodeMutation
}

// Where appends a list predicates to the OriginInviteCodeUpdate builder.
func (oicu *OriginInviteCodeUpdate) Where(ps ...predicate.OriginInviteCode) *OriginInviteCodeUpdate {
	oicu.mutation.Where(ps...)
	return oicu
}

// SetUpdatedAt sets the "updated_at" field.
func (oicu *OriginInviteCodeUpdate) SetUpdatedAt(t time.Time) *OriginInviteCodeUpdate {
	oicu.mutation.SetUpdatedAt(t)
	return oicu
}

// SetStatus sets the "status" field.
func (oicu *OriginInviteCodeUpdate) SetStatus(u uint8) *OriginInviteCodeUpdate {
	oicu.mutation.ResetStatus()
	oicu.mutation.SetStatus(u)
	return oicu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oicu *OriginInviteCodeUpdate) SetNillableStatus(u *uint8) *OriginInviteCodeUpdate {
	if u != nil {
		oicu.SetStatus(*u)
	}
	return oicu
}

// AddStatus adds u to the "status" field.
func (oicu *OriginInviteCodeUpdate) AddStatus(u int8) *OriginInviteCodeUpdate {
	oicu.mutation.AddStatus(u)
	return oicu
}

// ClearStatus clears the value of the "status" field.
func (oicu *OriginInviteCodeUpdate) ClearStatus() *OriginInviteCodeUpdate {
	oicu.mutation.ClearStatus()
	return oicu
}

// SetCode sets the "code" field.
func (oicu *OriginInviteCodeUpdate) SetCode(s string) *OriginInviteCodeUpdate {
	oicu.mutation.SetCode(s)
	return oicu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (oicu *OriginInviteCodeUpdate) SetNillableCode(s *string) *OriginInviteCodeUpdate {
	if s != nil {
		oicu.SetCode(*s)
	}
	return oicu
}

// Mutation returns the OriginInviteCodeMutation object of the builder.
func (oicu *OriginInviteCodeUpdate) Mutation() *OriginInviteCodeMutation {
	return oicu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oicu *OriginInviteCodeUpdate) Save(ctx context.Context) (int, error) {
	oicu.defaults()
	return withHooks(ctx, oicu.sqlSave, oicu.mutation, oicu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oicu *OriginInviteCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := oicu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oicu *OriginInviteCodeUpdate) Exec(ctx context.Context) error {
	_, err := oicu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicu *OriginInviteCodeUpdate) ExecX(ctx context.Context) {
	if err := oicu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oicu *OriginInviteCodeUpdate) defaults() {
	if _, ok := oicu.mutation.UpdatedAt(); !ok {
		v := origininvitecode.UpdateDefaultUpdatedAt()
		oicu.mutation.SetUpdatedAt(v)
	}
}

func (oicu *OriginInviteCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(origininvitecode.Table, origininvitecode.Columns, sqlgraph.NewFieldSpec(origininvitecode.FieldID, field.TypeUint64))
	if ps := oicu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oicu.mutation.UpdatedAt(); ok {
		_spec.SetField(origininvitecode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oicu.mutation.Status(); ok {
		_spec.SetField(origininvitecode.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := oicu.mutation.AddedStatus(); ok {
		_spec.AddField(origininvitecode.FieldStatus, field.TypeUint8, value)
	}
	if oicu.mutation.StatusCleared() {
		_spec.ClearField(origininvitecode.FieldStatus, field.TypeUint8)
	}
	if value, ok := oicu.mutation.Code(); ok {
		_spec.SetField(origininvitecode.FieldCode, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oicu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{origininvitecode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oicu.mutation.done = true
	return n, nil
}

// OriginInviteCodeUpdateOne is the builder for updating a single OriginInviteCode entity.
type OriginInviteCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OriginInviteCodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oicuo *OriginInviteCodeUpdateOne) SetUpdatedAt(t time.Time) *OriginInviteCodeUpdateOne {
	oicuo.mutation.SetUpdatedAt(t)
	return oicuo
}

// SetStatus sets the "status" field.
func (oicuo *OriginInviteCodeUpdateOne) SetStatus(u uint8) *OriginInviteCodeUpdateOne {
	oicuo.mutation.ResetStatus()
	oicuo.mutation.SetStatus(u)
	return oicuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oicuo *OriginInviteCodeUpdateOne) SetNillableStatus(u *uint8) *OriginInviteCodeUpdateOne {
	if u != nil {
		oicuo.SetStatus(*u)
	}
	return oicuo
}

// AddStatus adds u to the "status" field.
func (oicuo *OriginInviteCodeUpdateOne) AddStatus(u int8) *OriginInviteCodeUpdateOne {
	oicuo.mutation.AddStatus(u)
	return oicuo
}

// ClearStatus clears the value of the "status" field.
func (oicuo *OriginInviteCodeUpdateOne) ClearStatus() *OriginInviteCodeUpdateOne {
	oicuo.mutation.ClearStatus()
	return oicuo
}

// SetCode sets the "code" field.
func (oicuo *OriginInviteCodeUpdateOne) SetCode(s string) *OriginInviteCodeUpdateOne {
	oicuo.mutation.SetCode(s)
	return oicuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (oicuo *OriginInviteCodeUpdateOne) SetNillableCode(s *string) *OriginInviteCodeUpdateOne {
	if s != nil {
		oicuo.SetCode(*s)
	}
	return oicuo
}

// Mutation returns the OriginInviteCodeMutation object of the builder.
func (oicuo *OriginInviteCodeUpdateOne) Mutation() *OriginInviteCodeMutation {
	return oicuo.mutation
}

// Where appends a list predicates to the OriginInviteCodeUpdate builder.
func (oicuo *OriginInviteCodeUpdateOne) Where(ps ...predicate.OriginInviteCode) *OriginInviteCodeUpdateOne {
	oicuo.mutation.Where(ps...)
	return oicuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oicuo *OriginInviteCodeUpdateOne) Select(field string, fields ...string) *OriginInviteCodeUpdateOne {
	oicuo.fields = append([]string{field}, fields...)
	return oicuo
}

// Save executes the query and returns the updated OriginInviteCode entity.
func (oicuo *OriginInviteCodeUpdateOne) Save(ctx context.Context) (*OriginInviteCode, error) {
	oicuo.defaults()
	return withHooks(ctx, oicuo.sqlSave, oicuo.mutation, oicuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oicuo *OriginInviteCodeUpdateOne) SaveX(ctx context.Context) *OriginInviteCode {
	node, err := oicuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oicuo *OriginInviteCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := oicuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicuo *OriginInviteCodeUpdateOne) ExecX(ctx context.Context) {
	if err := oicuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oicuo *OriginInviteCodeUpdateOne) defaults() {
	if _, ok := oicuo.mutation.UpdatedAt(); !ok {
		v := origininvitecode.UpdateDefaultUpdatedAt()
		oicuo.mutation.SetUpdatedAt(v)
	}
}

func (oicuo *OriginInviteCodeUpdateOne) sqlSave(ctx context.Context) (_node *OriginInviteCode, err error) {
	_spec := sqlgraph.NewUpdateSpec(origininvitecode.Table, origininvitecode.Columns, sqlgraph.NewFieldSpec(origininvitecode.FieldID, field.TypeUint64))
	id, ok := oicuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OriginInviteCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oicuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, origininvitecode.FieldID)
		for _, f := range fields {
			if !origininvitecode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != origininvitecode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oicuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oicuo.mutation.UpdatedAt(); ok {
		_spec.SetField(origininvitecode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oicuo.mutation.Status(); ok {
		_spec.SetField(origininvitecode.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := oicuo.mutation.AddedStatus(); ok {
		_spec.AddField(origininvitecode.FieldStatus, field.TypeUint8, value)
	}
	if oicuo.mutation.StatusCleared() {
		_spec.ClearField(origininvitecode.FieldStatus, field.TypeUint8)
	}
	if value, ok := oicuo.mutation.Code(); ok {
		_spec.SetField(origininvitecode.FieldCode, field.TypeString, value)
	}
	_node = &OriginInviteCode{config: oicuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oicuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{origininvitecode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oicuo.mutation.done = true
	return _node, nil
}
