// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/player"
)

// PlayerCreate is the builder for creating a Player entity.
type PlayerCreate struct {
	config
	mutation *PlayerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PlayerCreate) SetCreatedAt(t time.Time) *PlayerCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableCreatedAt(t *time.Time) *PlayerCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PlayerCreate) SetUpdatedAt(t time.Time) *PlayerCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableUpdatedAt(t *time.Time) *PlayerCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PlayerCreate) SetStatus(u uint8) *PlayerCreate {
	pc.mutation.SetStatus(u)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableStatus(u *uint8) *PlayerCreate {
	if u != nil {
		pc.SetStatus(*u)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PlayerCreate) SetName(s string) *PlayerCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableName(s *string) *PlayerCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetEmail sets the "email" field.
func (pc *PlayerCreate) SetEmail(s string) *PlayerCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableEmail(s *string) *PlayerCreate {
	if s != nil {
		pc.SetEmail(*s)
	}
	return pc
}

// SetPassword sets the "password" field.
func (pc *PlayerCreate) SetPassword(s string) *PlayerCreate {
	pc.mutation.SetPassword(s)
	return pc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (pc *PlayerCreate) SetNillablePassword(s *string) *PlayerCreate {
	if s != nil {
		pc.SetPassword(*s)
	}
	return pc
}

// SetTransactionPassword sets the "transaction_password" field.
func (pc *PlayerCreate) SetTransactionPassword(s string) *PlayerCreate {
	pc.mutation.SetTransactionPassword(s)
	return pc
}

// SetNillableTransactionPassword sets the "transaction_password" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableTransactionPassword(s *string) *PlayerCreate {
	if s != nil {
		pc.SetTransactionPassword(*s)
	}
	return pc
}

// SetCoinLamb sets the "coin_lamb" field.
func (pc *PlayerCreate) SetCoinLamb(f float32) *PlayerCreate {
	pc.mutation.SetCoinLamb(f)
	return pc
}

// SetNillableCoinLamb sets the "coin_lamb" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableCoinLamb(f *float32) *PlayerCreate {
	if f != nil {
		pc.SetCoinLamb(*f)
	}
	return pc
}

// SetTokenLamb sets the "token_lamb" field.
func (pc *PlayerCreate) SetTokenLamb(f float32) *PlayerCreate {
	pc.mutation.SetTokenLamb(f)
	return pc
}

// SetNillableTokenLamb sets the "token_lamb" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableTokenLamb(f *float32) *PlayerCreate {
	if f != nil {
		pc.SetTokenLamb(*f)
	}
	return pc
}

// SetRank sets the "rank" field.
func (pc *PlayerCreate) SetRank(u uint32) *PlayerCreate {
	pc.mutation.SetRank(u)
	return pc
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableRank(u *uint32) *PlayerCreate {
	if u != nil {
		pc.SetRank(*u)
	}
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PlayerCreate) SetAmount(f float64) *PlayerCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableAmount(f *float64) *PlayerCreate {
	if f != nil {
		pc.SetAmount(*f)
	}
	return pc
}

// SetDepositAddress sets the "deposit_address" field.
func (pc *PlayerCreate) SetDepositAddress(s string) *PlayerCreate {
	pc.mutation.SetDepositAddress(s)
	return pc
}

// SetNillableDepositAddress sets the "deposit_address" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableDepositAddress(s *string) *PlayerCreate {
	if s != nil {
		pc.SetDepositAddress(*s)
	}
	return pc
}

// SetInvitedNum sets the "invited_num" field.
func (pc *PlayerCreate) SetInvitedNum(u uint32) *PlayerCreate {
	pc.mutation.SetInvitedNum(u)
	return pc
}

// SetNillableInvitedNum sets the "invited_num" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableInvitedNum(u *uint32) *PlayerCreate {
	if u != nil {
		pc.SetInvitedNum(*u)
	}
	return pc
}

// SetTotalIncome sets the "total_income" field.
func (pc *PlayerCreate) SetTotalIncome(f float64) *PlayerCreate {
	pc.mutation.SetTotalIncome(f)
	return pc
}

// SetNillableTotalIncome sets the "total_income" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableTotalIncome(f *float64) *PlayerCreate {
	if f != nil {
		pc.SetTotalIncome(*f)
	}
	return pc
}

// SetProfitAndLoss sets the "profit_and_loss" field.
func (pc *PlayerCreate) SetProfitAndLoss(f float32) *PlayerCreate {
	pc.mutation.SetProfitAndLoss(f)
	return pc
}

// SetNillableProfitAndLoss sets the "profit_and_loss" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableProfitAndLoss(f *float32) *PlayerCreate {
	if f != nil {
		pc.SetProfitAndLoss(*f)
	}
	return pc
}

// SetInviteCode sets the "invite_code" field.
func (pc *PlayerCreate) SetInviteCode(s string) *PlayerCreate {
	pc.mutation.SetInviteCode(s)
	return pc
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableInviteCode(s *string) *PlayerCreate {
	if s != nil {
		pc.SetInviteCode(*s)
	}
	return pc
}

// SetInviterID sets the "inviter_id" field.
func (pc *PlayerCreate) SetInviterID(u uint64) *PlayerCreate {
	pc.mutation.SetInviterID(u)
	return pc
}

// SetNillableInviterID sets the "inviter_id" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableInviterID(u *uint64) *PlayerCreate {
	if u != nil {
		pc.SetInviterID(*u)
	}
	return pc
}

// SetInvitedCode sets the "invited_code" field.
func (pc *PlayerCreate) SetInvitedCode(s string) *PlayerCreate {
	pc.mutation.SetInvitedCode(s)
	return pc
}

// SetNillableInvitedCode sets the "invited_code" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableInvitedCode(s *string) *PlayerCreate {
	if s != nil {
		pc.SetInvitedCode(*s)
	}
	return pc
}

// SetSystemCommission sets the "system_commission" field.
func (pc *PlayerCreate) SetSystemCommission(f float32) *PlayerCreate {
	pc.mutation.SetSystemCommission(f)
	return pc
}

// SetNillableSystemCommission sets the "system_commission" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableSystemCommission(f *float32) *PlayerCreate {
	if f != nil {
		pc.SetSystemCommission(*f)
	}
	return pc
}

// SetGcicsUserID sets the "gcics_user_id" field.
func (pc *PlayerCreate) SetGcicsUserID(u uint64) *PlayerCreate {
	pc.mutation.SetGcicsUserID(u)
	return pc
}

// SetNillableGcicsUserID sets the "gcics_user_id" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableGcicsUserID(u *uint64) *PlayerCreate {
	if u != nil {
		pc.SetGcicsUserID(*u)
	}
	return pc
}

// SetGcicsToken sets the "gcics_token" field.
func (pc *PlayerCreate) SetGcicsToken(s string) *PlayerCreate {
	pc.mutation.SetGcicsToken(s)
	return pc
}

// SetNillableGcicsToken sets the "gcics_token" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableGcicsToken(s *string) *PlayerCreate {
	if s != nil {
		pc.SetGcicsToken(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PlayerCreate) SetID(u uint64) *PlayerCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetInviter sets the "inviter" edge to the Player entity.
func (pc *PlayerCreate) SetInviter(p *Player) *PlayerCreate {
	return pc.SetInviterID(p.ID)
}

// AddInviteeIDs adds the "invitees" edge to the Player entity by IDs.
func (pc *PlayerCreate) AddInviteeIDs(ids ...uint64) *PlayerCreate {
	pc.mutation.AddInviteeIDs(ids...)
	return pc
}

// AddInvitees adds the "invitees" edges to the Player entity.
func (pc *PlayerCreate) AddInvitees(p ...*Player) *PlayerCreate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddInviteeIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pc *PlayerCreate) Mutation() *PlayerMutation {
	return pc.mutation
}

// Save creates the Player in the database.
func (pc *PlayerCreate) Save(ctx context.Context) (*Player, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlayerCreate) SaveX(ctx context.Context) *Player {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlayerCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlayerCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlayerCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := player.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := player.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := player.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := player.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Email(); !ok {
		v := player.DefaultEmail
		pc.mutation.SetEmail(v)
	}
	if _, ok := pc.mutation.Password(); !ok {
		v := player.DefaultPassword
		pc.mutation.SetPassword(v)
	}
	if _, ok := pc.mutation.TransactionPassword(); !ok {
		v := player.DefaultTransactionPassword
		pc.mutation.SetTransactionPassword(v)
	}
	if _, ok := pc.mutation.CoinLamb(); !ok {
		v := player.DefaultCoinLamb
		pc.mutation.SetCoinLamb(v)
	}
	if _, ok := pc.mutation.TokenLamb(); !ok {
		v := player.DefaultTokenLamb
		pc.mutation.SetTokenLamb(v)
	}
	if _, ok := pc.mutation.Rank(); !ok {
		v := player.DefaultRank
		pc.mutation.SetRank(v)
	}
	if _, ok := pc.mutation.Amount(); !ok {
		v := player.DefaultAmount
		pc.mutation.SetAmount(v)
	}
	if _, ok := pc.mutation.DepositAddress(); !ok {
		v := player.DefaultDepositAddress
		pc.mutation.SetDepositAddress(v)
	}
	if _, ok := pc.mutation.InvitedNum(); !ok {
		v := player.DefaultInvitedNum
		pc.mutation.SetInvitedNum(v)
	}
	if _, ok := pc.mutation.TotalIncome(); !ok {
		v := player.DefaultTotalIncome
		pc.mutation.SetTotalIncome(v)
	}
	if _, ok := pc.mutation.ProfitAndLoss(); !ok {
		v := player.DefaultProfitAndLoss
		pc.mutation.SetProfitAndLoss(v)
	}
	if _, ok := pc.mutation.InviteCode(); !ok {
		v := player.DefaultInviteCode
		pc.mutation.SetInviteCode(v)
	}
	if _, ok := pc.mutation.InviterID(); !ok {
		v := player.DefaultInviterID
		pc.mutation.SetInviterID(v)
	}
	if _, ok := pc.mutation.InvitedCode(); !ok {
		v := player.DefaultInvitedCode
		pc.mutation.SetInvitedCode(v)
	}
	if _, ok := pc.mutation.SystemCommission(); !ok {
		v := player.DefaultSystemCommission
		pc.mutation.SetSystemCommission(v)
	}
	if _, ok := pc.mutation.GcicsUserID(); !ok {
		v := player.DefaultGcicsUserID
		pc.mutation.SetGcicsUserID(v)
	}
	if _, ok := pc.mutation.GcicsToken(); !ok {
		v := player.DefaultGcicsToken
		pc.mutation.SetGcicsToken(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlayerCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Player.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Player.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Player.name"`)}
	}
	if _, ok := pc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Player.email"`)}
	}
	if _, ok := pc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Player.password"`)}
	}
	if _, ok := pc.mutation.TransactionPassword(); !ok {
		return &ValidationError{Name: "transaction_password", err: errors.New(`ent: missing required field "Player.transaction_password"`)}
	}
	if _, ok := pc.mutation.CoinLamb(); !ok {
		return &ValidationError{Name: "coin_lamb", err: errors.New(`ent: missing required field "Player.coin_lamb"`)}
	}
	if _, ok := pc.mutation.TokenLamb(); !ok {
		return &ValidationError{Name: "token_lamb", err: errors.New(`ent: missing required field "Player.token_lamb"`)}
	}
	if _, ok := pc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New(`ent: missing required field "Player.rank"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Player.amount"`)}
	}
	if _, ok := pc.mutation.DepositAddress(); !ok {
		return &ValidationError{Name: "deposit_address", err: errors.New(`ent: missing required field "Player.deposit_address"`)}
	}
	if _, ok := pc.mutation.InvitedNum(); !ok {
		return &ValidationError{Name: "invited_num", err: errors.New(`ent: missing required field "Player.invited_num"`)}
	}
	if _, ok := pc.mutation.TotalIncome(); !ok {
		return &ValidationError{Name: "total_income", err: errors.New(`ent: missing required field "Player.total_income"`)}
	}
	if _, ok := pc.mutation.ProfitAndLoss(); !ok {
		return &ValidationError{Name: "profit_and_loss", err: errors.New(`ent: missing required field "Player.profit_and_loss"`)}
	}
	if _, ok := pc.mutation.InviteCode(); !ok {
		return &ValidationError{Name: "invite_code", err: errors.New(`ent: missing required field "Player.invite_code"`)}
	}
	if _, ok := pc.mutation.InvitedCode(); !ok {
		return &ValidationError{Name: "invited_code", err: errors.New(`ent: missing required field "Player.invited_code"`)}
	}
	if _, ok := pc.mutation.SystemCommission(); !ok {
		return &ValidationError{Name: "system_commission", err: errors.New(`ent: missing required field "Player.system_commission"`)}
	}
	if _, ok := pc.mutation.GcicsUserID(); !ok {
		return &ValidationError{Name: "gcics_user_id", err: errors.New(`ent: missing required field "Player.gcics_user_id"`)}
	}
	if _, ok := pc.mutation.GcicsToken(); !ok {
		return &ValidationError{Name: "gcics_token", err: errors.New(`ent: missing required field "Player.gcics_token"`)}
	}
	return nil
}

func (pc *PlayerCreate) sqlSave(ctx context.Context) (*Player, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlayerCreate) createSpec() (*Player, *sqlgraph.CreateSpec) {
	var (
		_node = &Player{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(player.Table, sqlgraph.NewFieldSpec(player.FieldID, field.TypeUint64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(player.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(player.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(player.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(player.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.Password(); ok {
		_spec.SetField(player.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := pc.mutation.TransactionPassword(); ok {
		_spec.SetField(player.FieldTransactionPassword, field.TypeString, value)
		_node.TransactionPassword = value
	}
	if value, ok := pc.mutation.CoinLamb(); ok {
		_spec.SetField(player.FieldCoinLamb, field.TypeFloat32, value)
		_node.CoinLamb = value
	}
	if value, ok := pc.mutation.TokenLamb(); ok {
		_spec.SetField(player.FieldTokenLamb, field.TypeFloat32, value)
		_node.TokenLamb = value
	}
	if value, ok := pc.mutation.Rank(); ok {
		_spec.SetField(player.FieldRank, field.TypeUint32, value)
		_node.Rank = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(player.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.DepositAddress(); ok {
		_spec.SetField(player.FieldDepositAddress, field.TypeString, value)
		_node.DepositAddress = value
	}
	if value, ok := pc.mutation.InvitedNum(); ok {
		_spec.SetField(player.FieldInvitedNum, field.TypeUint32, value)
		_node.InvitedNum = value
	}
	if value, ok := pc.mutation.TotalIncome(); ok {
		_spec.SetField(player.FieldTotalIncome, field.TypeFloat64, value)
		_node.TotalIncome = value
	}
	if value, ok := pc.mutation.ProfitAndLoss(); ok {
		_spec.SetField(player.FieldProfitAndLoss, field.TypeFloat32, value)
		_node.ProfitAndLoss = value
	}
	if value, ok := pc.mutation.InviteCode(); ok {
		_spec.SetField(player.FieldInviteCode, field.TypeString, value)
		_node.InviteCode = value
	}
	if value, ok := pc.mutation.InvitedCode(); ok {
		_spec.SetField(player.FieldInvitedCode, field.TypeString, value)
		_node.InvitedCode = value
	}
	if value, ok := pc.mutation.SystemCommission(); ok {
		_spec.SetField(player.FieldSystemCommission, field.TypeFloat32, value)
		_node.SystemCommission = value
	}
	if value, ok := pc.mutation.GcicsUserID(); ok {
		_spec.SetField(player.FieldGcicsUserID, field.TypeUint64, value)
		_node.GcicsUserID = value
	}
	if value, ok := pc.mutation.GcicsToken(); ok {
		_spec.SetField(player.FieldGcicsToken, field.TypeString, value)
		_node.GcicsToken = value
	}
	if nodes := pc.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   player.InviterTable,
			Columns: []string{player.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InviterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.InviteesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.InviteesTable,
			Columns: []string{player.InviteesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerCreateBulk is the builder for creating many Player entities in bulk.
type PlayerCreateBulk struct {
	config
	err      error
	builders []*PlayerCreate
}

// Save creates the Player entities in the database.
func (pcb *PlayerCreateBulk) Save(ctx context.Context) ([]*Player, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Player, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlayerCreateBulk) SaveX(ctx context.Context) []*Player {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlayerCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlayerCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
