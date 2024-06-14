// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/round"
)

// 回合表
type Round struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Round Count | 当日累计回合数
	RoundCount uint32 `json:"round_count,omitempty"`
	// 累计第几回合
	TotalRoundCount uint64 `json:"total_round_count,omitempty"`
	// Start Time | 回合开始时间（包括倒计时）
	StartAt time.Time `json:"start_at,omitempty"`
	// Open Time | 回合开奖时间
	OpenAt time.Time `json:"open_at,omitempty"`
	// End Time | 回合结束时间
	EndAt time.Time `json:"end_at,omitempty"`
	// 选中的羊圈号码
	SelectedFold uint32 `json:"selected_fold,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoundQuery when eager-loading is set.
	Edges        RoundEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoundEdges holds the relations/edges for other nodes in the graph.
type RoundEdges struct {
	// Fold holds the value of the fold edge.
	Fold []*RoundLambFold `json:"fold,omitempty"`
	// Invest holds the value of the invest edge.
	Invest []*RoundInvest `json:"invest,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FoldOrErr returns the Fold value or an error if the edge
// was not loaded in eager-loading.
func (e RoundEdges) FoldOrErr() ([]*RoundLambFold, error) {
	if e.loadedTypes[0] {
		return e.Fold, nil
	}
	return nil, &NotLoadedError{edge: "fold"}
}

// InvestOrErr returns the Invest value or an error if the edge
// was not loaded in eager-loading.
func (e RoundEdges) InvestOrErr() ([]*RoundInvest, error) {
	if e.loadedTypes[1] {
		return e.Invest, nil
	}
	return nil, &NotLoadedError{edge: "invest"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Round) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case round.FieldID, round.FieldStatus, round.FieldRoundCount, round.FieldTotalRoundCount, round.FieldSelectedFold:
			values[i] = new(sql.NullInt64)
		case round.FieldCreatedAt, round.FieldUpdatedAt, round.FieldStartAt, round.FieldOpenAt, round.FieldEndAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Round fields.
func (r *Round) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case round.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint64(value.Int64)
		case round.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case round.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case round.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = uint8(value.Int64)
			}
		case round.FieldRoundCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field round_count", values[i])
			} else if value.Valid {
				r.RoundCount = uint32(value.Int64)
			}
		case round.FieldTotalRoundCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_round_count", values[i])
			} else if value.Valid {
				r.TotalRoundCount = uint64(value.Int64)
			}
		case round.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				r.StartAt = value.Time
			}
		case round.FieldOpenAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field open_at", values[i])
			} else if value.Valid {
				r.OpenAt = value.Time
			}
		case round.FieldEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				r.EndAt = value.Time
			}
		case round.FieldSelectedFold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field selected_fold", values[i])
			} else if value.Valid {
				r.SelectedFold = uint32(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Round.
// This includes values selected through modifiers, order, etc.
func (r *Round) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryFold queries the "fold" edge of the Round entity.
func (r *Round) QueryFold() *RoundLambFoldQuery {
	return NewRoundClient(r.config).QueryFold(r)
}

// QueryInvest queries the "invest" edge of the Round entity.
func (r *Round) QueryInvest() *RoundInvestQuery {
	return NewRoundClient(r.config).QueryInvest(r)
}

// Update returns a builder for updating this Round.
// Note that you need to call Round.Unwrap() before calling this method if this Round
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Round) Update() *RoundUpdateOne {
	return NewRoundClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Round entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Round) Unwrap() *Round {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Round is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Round) String() string {
	var builder strings.Builder
	builder.WriteString("Round(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", ")
	builder.WriteString("round_count=")
	builder.WriteString(fmt.Sprintf("%v", r.RoundCount))
	builder.WriteString(", ")
	builder.WriteString("total_round_count=")
	builder.WriteString(fmt.Sprintf("%v", r.TotalRoundCount))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(r.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("open_at=")
	builder.WriteString(r.OpenAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(r.EndAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("selected_fold=")
	builder.WriteString(fmt.Sprintf("%v", r.SelectedFold))
	builder.WriteByte(')')
	return builder.String()
}

// Rounds is a parsable slice of Round.
type Rounds []*Round
