// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
)

// RoundLambFoldCreate is the builder for creating a RoundLambFold entity.
type RoundLambFoldCreate struct {
	config
	mutation *RoundLambFoldMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rlfc *RoundLambFoldCreate) SetCreatedAt(t time.Time) *RoundLambFoldCreate {
	rlfc.mutation.SetCreatedAt(t)
	return rlfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableCreatedAt(t *time.Time) *RoundLambFoldCreate {
	if t != nil {
		rlfc.SetCreatedAt(*t)
	}
	return rlfc
}

// SetUpdatedAt sets the "updated_at" field.
func (rlfc *RoundLambFoldCreate) SetUpdatedAt(t time.Time) *RoundLambFoldCreate {
	rlfc.mutation.SetUpdatedAt(t)
	return rlfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableUpdatedAt(t *time.Time) *RoundLambFoldCreate {
	if t != nil {
		rlfc.SetUpdatedAt(*t)
	}
	return rlfc
}

// SetFoldNo sets the "fold_no" field.
func (rlfc *RoundLambFoldCreate) SetFoldNo(u uint32) *RoundLambFoldCreate {
	rlfc.mutation.SetFoldNo(u)
	return rlfc
}

// SetNillableFoldNo sets the "fold_no" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableFoldNo(u *uint32) *RoundLambFoldCreate {
	if u != nil {
		rlfc.SetFoldNo(*u)
	}
	return rlfc
}

// SetLambNum sets the "lamb_num" field.
func (rlfc *RoundLambFoldCreate) SetLambNum(u uint32) *RoundLambFoldCreate {
	rlfc.mutation.SetLambNum(u)
	return rlfc
}

// SetNillableLambNum sets the "lamb_num" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableLambNum(u *uint32) *RoundLambFoldCreate {
	if u != nil {
		rlfc.SetLambNum(*u)
	}
	return rlfc
}

// SetRoundID sets the "round_id" field.
func (rlfc *RoundLambFoldCreate) SetRoundID(u uint64) *RoundLambFoldCreate {
	rlfc.mutation.SetRoundID(u)
	return rlfc
}

// SetNillableRoundID sets the "round_id" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableRoundID(u *uint64) *RoundLambFoldCreate {
	if u != nil {
		rlfc.SetRoundID(*u)
	}
	return rlfc
}

// SetRoundCount sets the "round_count" field.
func (rlfc *RoundLambFoldCreate) SetRoundCount(u uint32) *RoundLambFoldCreate {
	rlfc.mutation.SetRoundCount(u)
	return rlfc
}

// SetNillableRoundCount sets the "round_count" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableRoundCount(u *uint32) *RoundLambFoldCreate {
	if u != nil {
		rlfc.SetRoundCount(*u)
	}
	return rlfc
}

// SetProfitAndLoss sets the "profit_and_loss" field.
func (rlfc *RoundLambFoldCreate) SetProfitAndLoss(f float32) *RoundLambFoldCreate {
	rlfc.mutation.SetProfitAndLoss(f)
	return rlfc
}

// SetNillableProfitAndLoss sets the "profit_and_loss" field if the given value is not nil.
func (rlfc *RoundLambFoldCreate) SetNillableProfitAndLoss(f *float32) *RoundLambFoldCreate {
	if f != nil {
		rlfc.SetProfitAndLoss(*f)
	}
	return rlfc
}

// SetID sets the "id" field.
func (rlfc *RoundLambFoldCreate) SetID(u uint64) *RoundLambFoldCreate {
	rlfc.mutation.SetID(u)
	return rlfc
}

// SetRound sets the "round" edge to the Round entity.
func (rlfc *RoundLambFoldCreate) SetRound(r *Round) *RoundLambFoldCreate {
	return rlfc.SetRoundID(r.ID)
}

// Mutation returns the RoundLambFoldMutation object of the builder.
func (rlfc *RoundLambFoldCreate) Mutation() *RoundLambFoldMutation {
	return rlfc.mutation
}

// Save creates the RoundLambFold in the database.
func (rlfc *RoundLambFoldCreate) Save(ctx context.Context) (*RoundLambFold, error) {
	rlfc.defaults()
	return withHooks(ctx, rlfc.sqlSave, rlfc.mutation, rlfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rlfc *RoundLambFoldCreate) SaveX(ctx context.Context) *RoundLambFold {
	v, err := rlfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlfc *RoundLambFoldCreate) Exec(ctx context.Context) error {
	_, err := rlfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlfc *RoundLambFoldCreate) ExecX(ctx context.Context) {
	if err := rlfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rlfc *RoundLambFoldCreate) defaults() {
	if _, ok := rlfc.mutation.CreatedAt(); !ok {
		v := roundlambfold.DefaultCreatedAt()
		rlfc.mutation.SetCreatedAt(v)
	}
	if _, ok := rlfc.mutation.UpdatedAt(); !ok {
		v := roundlambfold.DefaultUpdatedAt()
		rlfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rlfc.mutation.FoldNo(); !ok {
		v := roundlambfold.DefaultFoldNo
		rlfc.mutation.SetFoldNo(v)
	}
	if _, ok := rlfc.mutation.LambNum(); !ok {
		v := roundlambfold.DefaultLambNum
		rlfc.mutation.SetLambNum(v)
	}
	if _, ok := rlfc.mutation.RoundID(); !ok {
		v := roundlambfold.DefaultRoundID
		rlfc.mutation.SetRoundID(v)
	}
	if _, ok := rlfc.mutation.RoundCount(); !ok {
		v := roundlambfold.DefaultRoundCount
		rlfc.mutation.SetRoundCount(v)
	}
	if _, ok := rlfc.mutation.ProfitAndLoss(); !ok {
		v := roundlambfold.DefaultProfitAndLoss
		rlfc.mutation.SetProfitAndLoss(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rlfc *RoundLambFoldCreate) check() error {
	if _, ok := rlfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RoundLambFold.created_at"`)}
	}
	if _, ok := rlfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RoundLambFold.updated_at"`)}
	}
	if _, ok := rlfc.mutation.FoldNo(); !ok {
		return &ValidationError{Name: "fold_no", err: errors.New(`ent: missing required field "RoundLambFold.fold_no"`)}
	}
	if _, ok := rlfc.mutation.LambNum(); !ok {
		return &ValidationError{Name: "lamb_num", err: errors.New(`ent: missing required field "RoundLambFold.lamb_num"`)}
	}
	if _, ok := rlfc.mutation.ProfitAndLoss(); !ok {
		return &ValidationError{Name: "profit_and_loss", err: errors.New(`ent: missing required field "RoundLambFold.profit_and_loss"`)}
	}
	return nil
}

func (rlfc *RoundLambFoldCreate) sqlSave(ctx context.Context) (*RoundLambFold, error) {
	if err := rlfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rlfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rlfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rlfc.mutation.id = &_node.ID
	rlfc.mutation.done = true
	return _node, nil
}

func (rlfc *RoundLambFoldCreate) createSpec() (*RoundLambFold, *sqlgraph.CreateSpec) {
	var (
		_node = &RoundLambFold{config: rlfc.config}
		_spec = sqlgraph.NewCreateSpec(roundlambfold.Table, sqlgraph.NewFieldSpec(roundlambfold.FieldID, field.TypeUint64))
	)
	if id, ok := rlfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rlfc.mutation.CreatedAt(); ok {
		_spec.SetField(roundlambfold.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rlfc.mutation.UpdatedAt(); ok {
		_spec.SetField(roundlambfold.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rlfc.mutation.FoldNo(); ok {
		_spec.SetField(roundlambfold.FieldFoldNo, field.TypeUint32, value)
		_node.FoldNo = value
	}
	if value, ok := rlfc.mutation.LambNum(); ok {
		_spec.SetField(roundlambfold.FieldLambNum, field.TypeUint32, value)
		_node.LambNum = value
	}
	if value, ok := rlfc.mutation.RoundCount(); ok {
		_spec.SetField(roundlambfold.FieldRoundCount, field.TypeUint32, value)
		_node.RoundCount = value
	}
	if value, ok := rlfc.mutation.ProfitAndLoss(); ok {
		_spec.SetField(roundlambfold.FieldProfitAndLoss, field.TypeFloat32, value)
		_node.ProfitAndLoss = value
	}
	if nodes := rlfc.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roundlambfold.RoundTable,
			Columns: []string{roundlambfold.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoundID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoundLambFoldCreateBulk is the builder for creating many RoundLambFold entities in bulk.
type RoundLambFoldCreateBulk struct {
	config
	err      error
	builders []*RoundLambFoldCreate
}

// Save creates the RoundLambFold entities in the database.
func (rlfcb *RoundLambFoldCreateBulk) Save(ctx context.Context) ([]*RoundLambFold, error) {
	if rlfcb.err != nil {
		return nil, rlfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rlfcb.builders))
	nodes := make([]*RoundLambFold, len(rlfcb.builders))
	mutators := make([]Mutator, len(rlfcb.builders))
	for i := range rlfcb.builders {
		func(i int, root context.Context) {
			builder := rlfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoundLambFoldMutation)
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
					_, err = mutators[i+1].Mutate(root, rlfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rlfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rlfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rlfcb *RoundLambFoldCreateBulk) SaveX(ctx context.Context) []*RoundLambFold {
	v, err := rlfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlfcb *RoundLambFoldCreateBulk) Exec(ctx context.Context) error {
	_, err := rlfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlfcb *RoundLambFoldCreateBulk) ExecX(ctx context.Context) {
	if err := rlfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
