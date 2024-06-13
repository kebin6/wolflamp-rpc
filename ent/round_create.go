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
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
)

// RoundCreate is the builder for creating a Round entity.
type RoundCreate struct {
	config
	mutation *RoundMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoundCreate) SetCreatedAt(t time.Time) *RoundCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoundCreate) SetNillableCreatedAt(t *time.Time) *RoundCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoundCreate) SetUpdatedAt(t time.Time) *RoundCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoundCreate) SetNillableUpdatedAt(t *time.Time) *RoundCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoundCreate) SetStatus(u uint8) *RoundCreate {
	rc.mutation.SetStatus(u)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoundCreate) SetNillableStatus(u *uint8) *RoundCreate {
	if u != nil {
		rc.SetStatus(*u)
	}
	return rc
}

// SetRoundCount sets the "round_count" field.
func (rc *RoundCreate) SetRoundCount(u uint32) *RoundCreate {
	rc.mutation.SetRoundCount(u)
	return rc
}

// SetStartAt sets the "start_at" field.
func (rc *RoundCreate) SetStartAt(t time.Time) *RoundCreate {
	rc.mutation.SetStartAt(t)
	return rc
}

// SetOpenAt sets the "open_at" field.
func (rc *RoundCreate) SetOpenAt(t time.Time) *RoundCreate {
	rc.mutation.SetOpenAt(t)
	return rc
}

// SetEndAt sets the "end_at" field.
func (rc *RoundCreate) SetEndAt(t time.Time) *RoundCreate {
	rc.mutation.SetEndAt(t)
	return rc
}

// SetSelectedFold sets the "selected_fold" field.
func (rc *RoundCreate) SetSelectedFold(u uint32) *RoundCreate {
	rc.mutation.SetSelectedFold(u)
	return rc
}

// SetNillableSelectedFold sets the "selected_fold" field if the given value is not nil.
func (rc *RoundCreate) SetNillableSelectedFold(u *uint32) *RoundCreate {
	if u != nil {
		rc.SetSelectedFold(*u)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoundCreate) SetID(u uint64) *RoundCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddFoldIDs adds the "fold" edge to the RoundLambFold entity by IDs.
func (rc *RoundCreate) AddFoldIDs(ids ...uint64) *RoundCreate {
	rc.mutation.AddFoldIDs(ids...)
	return rc
}

// AddFold adds the "fold" edges to the RoundLambFold entity.
func (rc *RoundCreate) AddFold(r ...*RoundLambFold) *RoundCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddFoldIDs(ids...)
}

// AddInvestIDs adds the "invest" edge to the RoundInvest entity by IDs.
func (rc *RoundCreate) AddInvestIDs(ids ...uint64) *RoundCreate {
	rc.mutation.AddInvestIDs(ids...)
	return rc
}

// AddInvest adds the "invest" edges to the RoundInvest entity.
func (rc *RoundCreate) AddInvest(r ...*RoundInvest) *RoundCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddInvestIDs(ids...)
}

// Mutation returns the RoundMutation object of the builder.
func (rc *RoundCreate) Mutation() *RoundMutation {
	return rc.mutation
}

// Save creates the Round in the database.
func (rc *RoundCreate) Save(ctx context.Context) (*Round, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoundCreate) SaveX(ctx context.Context) *Round {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoundCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoundCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoundCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := round.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := round.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := round.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.SelectedFold(); !ok {
		v := round.DefaultSelectedFold
		rc.mutation.SetSelectedFold(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoundCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Round.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Round.updated_at"`)}
	}
	if _, ok := rc.mutation.RoundCount(); !ok {
		return &ValidationError{Name: "round_count", err: errors.New(`ent: missing required field "Round.round_count"`)}
	}
	if _, ok := rc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "Round.start_at"`)}
	}
	if _, ok := rc.mutation.OpenAt(); !ok {
		return &ValidationError{Name: "open_at", err: errors.New(`ent: missing required field "Round.open_at"`)}
	}
	if _, ok := rc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "Round.end_at"`)}
	}
	if _, ok := rc.mutation.SelectedFold(); !ok {
		return &ValidationError{Name: "selected_fold", err: errors.New(`ent: missing required field "Round.selected_fold"`)}
	}
	return nil
}

func (rc *RoundCreate) sqlSave(ctx context.Context) (*Round, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoundCreate) createSpec() (*Round, *sqlgraph.CreateSpec) {
	var (
		_node = &Round{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(round.Table, sqlgraph.NewFieldSpec(round.FieldID, field.TypeUint64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(round.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(round.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(round.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.RoundCount(); ok {
		_spec.SetField(round.FieldRoundCount, field.TypeUint32, value)
		_node.RoundCount = value
	}
	if value, ok := rc.mutation.StartAt(); ok {
		_spec.SetField(round.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := rc.mutation.OpenAt(); ok {
		_spec.SetField(round.FieldOpenAt, field.TypeTime, value)
		_node.OpenAt = value
	}
	if value, ok := rc.mutation.EndAt(); ok {
		_spec.SetField(round.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := rc.mutation.SelectedFold(); ok {
		_spec.SetField(round.FieldSelectedFold, field.TypeUint32, value)
		_node.SelectedFold = value
	}
	if nodes := rc.mutation.FoldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.FoldTable,
			Columns: []string{round.FoldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roundlambfold.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.InvestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.InvestTable,
			Columns: []string{round.InvestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roundinvest.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoundCreateBulk is the builder for creating many Round entities in bulk.
type RoundCreateBulk struct {
	config
	err      error
	builders []*RoundCreate
}

// Save creates the Round entities in the database.
func (rcb *RoundCreateBulk) Save(ctx context.Context) ([]*Round, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Round, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoundMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoundCreateBulk) SaveX(ctx context.Context) []*Round {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoundCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoundCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
