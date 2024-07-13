// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
)

// 回合羊圈情况表
type RoundLambFold struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 羊圈号
	FoldNo uint32 `json:"fold_no,omitempty"`
	// 投注羊数量
	LambNum uint32 `json:"lamb_num,omitempty"`
	// 所属回合ID
	RoundID uint64 `json:"round_id,omitempty"`
	// 投注盈亏结果
	ProfitAndLoss float32 `json:"profit_and_loss,omitempty"`
	// 当日第几回合
	RoundCount uint32 `json:"round_count,omitempty"`
	// 累计第几回合
	TotalRoundCount uint64 `json:"total_round_count,omitempty"`
	// 游戏类型：coin,token
	Mode string `json:"mode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoundLambFoldQuery when eager-loading is set.
	Edges        RoundLambFoldEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoundLambFoldEdges holds the relations/edges for other nodes in the graph.
type RoundLambFoldEdges struct {
	// Round holds the value of the round edge.
	Round *Round `json:"round,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoundOrErr returns the Round value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoundLambFoldEdges) RoundOrErr() (*Round, error) {
	if e.Round != nil {
		return e.Round, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: round.Label}
	}
	return nil, &NotLoadedError{edge: "round"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoundLambFold) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roundlambfold.FieldProfitAndLoss:
			values[i] = new(sql.NullFloat64)
		case roundlambfold.FieldID, roundlambfold.FieldFoldNo, roundlambfold.FieldLambNum, roundlambfold.FieldRoundID, roundlambfold.FieldRoundCount, roundlambfold.FieldTotalRoundCount:
			values[i] = new(sql.NullInt64)
		case roundlambfold.FieldMode:
			values[i] = new(sql.NullString)
		case roundlambfold.FieldCreatedAt, roundlambfold.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoundLambFold fields.
func (rlf *RoundLambFold) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roundlambfold.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rlf.ID = uint64(value.Int64)
		case roundlambfold.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rlf.CreatedAt = value.Time
			}
		case roundlambfold.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rlf.UpdatedAt = value.Time
			}
		case roundlambfold.FieldFoldNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fold_no", values[i])
			} else if value.Valid {
				rlf.FoldNo = uint32(value.Int64)
			}
		case roundlambfold.FieldLambNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lamb_num", values[i])
			} else if value.Valid {
				rlf.LambNum = uint32(value.Int64)
			}
		case roundlambfold.FieldRoundID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field round_id", values[i])
			} else if value.Valid {
				rlf.RoundID = uint64(value.Int64)
			}
		case roundlambfold.FieldProfitAndLoss:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field profit_and_loss", values[i])
			} else if value.Valid {
				rlf.ProfitAndLoss = float32(value.Float64)
			}
		case roundlambfold.FieldRoundCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field round_count", values[i])
			} else if value.Valid {
				rlf.RoundCount = uint32(value.Int64)
			}
		case roundlambfold.FieldTotalRoundCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_round_count", values[i])
			} else if value.Valid {
				rlf.TotalRoundCount = uint64(value.Int64)
			}
		case roundlambfold.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				rlf.Mode = value.String
			}
		default:
			rlf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoundLambFold.
// This includes values selected through modifiers, order, etc.
func (rlf *RoundLambFold) Value(name string) (ent.Value, error) {
	return rlf.selectValues.Get(name)
}

// QueryRound queries the "round" edge of the RoundLambFold entity.
func (rlf *RoundLambFold) QueryRound() *RoundQuery {
	return NewRoundLambFoldClient(rlf.config).QueryRound(rlf)
}

// Update returns a builder for updating this RoundLambFold.
// Note that you need to call RoundLambFold.Unwrap() before calling this method if this RoundLambFold
// was returned from a transaction, and the transaction was committed or rolled back.
func (rlf *RoundLambFold) Update() *RoundLambFoldUpdateOne {
	return NewRoundLambFoldClient(rlf.config).UpdateOne(rlf)
}

// Unwrap unwraps the RoundLambFold entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rlf *RoundLambFold) Unwrap() *RoundLambFold {
	_tx, ok := rlf.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoundLambFold is not a transactional entity")
	}
	rlf.config.driver = _tx.drv
	return rlf
}

// String implements the fmt.Stringer.
func (rlf *RoundLambFold) String() string {
	var builder strings.Builder
	builder.WriteString("RoundLambFold(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rlf.ID))
	builder.WriteString("created_at=")
	builder.WriteString(rlf.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rlf.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fold_no=")
	builder.WriteString(fmt.Sprintf("%v", rlf.FoldNo))
	builder.WriteString(", ")
	builder.WriteString("lamb_num=")
	builder.WriteString(fmt.Sprintf("%v", rlf.LambNum))
	builder.WriteString(", ")
	builder.WriteString("round_id=")
	builder.WriteString(fmt.Sprintf("%v", rlf.RoundID))
	builder.WriteString(", ")
	builder.WriteString("profit_and_loss=")
	builder.WriteString(fmt.Sprintf("%v", rlf.ProfitAndLoss))
	builder.WriteString(", ")
	builder.WriteString("round_count=")
	builder.WriteString(fmt.Sprintf("%v", rlf.RoundCount))
	builder.WriteString(", ")
	builder.WriteString("total_round_count=")
	builder.WriteString(fmt.Sprintf("%v", rlf.TotalRoundCount))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(rlf.Mode)
	builder.WriteByte(')')
	return builder.String()
}

// RoundLambFolds is a parsable slice of RoundLambFold.
type RoundLambFolds []*RoundLambFold
