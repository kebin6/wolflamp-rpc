// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SettingCreate) SetCreatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SettingCreate) SetUpdatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *SettingCreate) SetStatus(u uint8) *SettingCreate {
	sc.mutation.SetStatus(u)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SettingCreate) SetNillableStatus(u *uint8) *SettingCreate {
	if u != nil {
		sc.SetStatus(*u)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SettingCreate) SetName(s string) *SettingCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *SettingCreate) SetNillableName(s *string) *SettingCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// SetModule sets the "module" field.
func (sc *SettingCreate) SetModule(s string) *SettingCreate {
	sc.mutation.SetModule(s)
	return sc
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (sc *SettingCreate) SetNillableModule(s *string) *SettingCreate {
	if s != nil {
		sc.SetModule(*s)
	}
	return sc
}

// SetValue sets the "value" field.
func (sc *SettingCreate) SetValue(s string) *SettingCreate {
	sc.mutation.SetValue(s)
	return sc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (sc *SettingCreate) SetNillableValue(s *string) *SettingCreate {
	if s != nil {
		sc.SetValue(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SettingCreate) SetID(u uint64) *SettingCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := setting.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := setting.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := setting.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Name(); !ok {
		v := setting.DefaultName
		sc.mutation.SetName(v)
	}
	if _, ok := sc.mutation.Module(); !ok {
		v := setting.DefaultModule
		sc.mutation.SetModule(v)
	}
	if _, ok := sc.mutation.Value(); !ok {
		v := setting.DefaultValue
		sc.mutation.SetValue(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Setting.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Setting.updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Setting.name"`)}
	}
	if _, ok := sc.mutation.Module(); !ok {
		return &ValidationError{Name: "module", err: errors.New(`ent: missing required field "Setting.module"`)}
	}
	if _, ok := sc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Setting.value"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeUint64))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(setting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(setting.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(setting.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Module(); ok {
		_spec.SetField(setting.FieldModule, field.TypeString, value)
		_node.Module = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
		_node.Value = &value
	}
	return _node, _spec
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	err      error
	builders []*SettingCreate
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}