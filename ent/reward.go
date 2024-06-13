// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
)

// 奖励表
type Reward struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// reward to player 's id | 奖励对象ID
	ToID uint64 `json:"to_id,omitempty"`
	// contributor 's id | 贡献者ID,平台为0
	ContributorID uint64 `json:"contributor_id,omitempty"`
	// contributor 's email | 贡献者邮箱，平台为system
	ContributorEmail string `json:"contributor_email,omitempty"`
	// contributor level | 贡献者属于奖励对象第几级
	ContributorLevel uint32 `json:"contributor_level,omitempty"`
	// lamb amount | 数量
	Num float32 `json:"num,omitempty"`
	// 计算公式
	Formula string `json:"formula,omitempty"`
	// remark | 备注
	Remark       string `json:"remark,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reward) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reward.FieldNum:
			values[i] = new(sql.NullFloat64)
		case reward.FieldID, reward.FieldStatus, reward.FieldToID, reward.FieldContributorID, reward.FieldContributorLevel:
			values[i] = new(sql.NullInt64)
		case reward.FieldContributorEmail, reward.FieldFormula, reward.FieldRemark:
			values[i] = new(sql.NullString)
		case reward.FieldCreatedAt, reward.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reward fields.
func (r *Reward) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reward.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint64(value.Int64)
		case reward.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reward.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reward.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = uint8(value.Int64)
			}
		case reward.FieldToID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_id", values[i])
			} else if value.Valid {
				r.ToID = uint64(value.Int64)
			}
		case reward.FieldContributorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contributor_id", values[i])
			} else if value.Valid {
				r.ContributorID = uint64(value.Int64)
			}
		case reward.FieldContributorEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contributor_email", values[i])
			} else if value.Valid {
				r.ContributorEmail = value.String
			}
		case reward.FieldContributorLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contributor_level", values[i])
			} else if value.Valid {
				r.ContributorLevel = uint32(value.Int64)
			}
		case reward.FieldNum:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				r.Num = float32(value.Float64)
			}
		case reward.FieldFormula:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field formula", values[i])
			} else if value.Valid {
				r.Formula = value.String
			}
		case reward.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				r.Remark = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reward.
// This includes values selected through modifiers, order, etc.
func (r *Reward) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Reward.
// Note that you need to call Reward.Unwrap() before calling this method if this Reward
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reward) Update() *RewardUpdateOne {
	return NewRewardClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reward entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reward) Unwrap() *Reward {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reward is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reward) String() string {
	var builder strings.Builder
	builder.WriteString("Reward(")
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
	builder.WriteString("to_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ToID))
	builder.WriteString(", ")
	builder.WriteString("contributor_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ContributorID))
	builder.WriteString(", ")
	builder.WriteString("contributor_email=")
	builder.WriteString(r.ContributorEmail)
	builder.WriteString(", ")
	builder.WriteString("contributor_level=")
	builder.WriteString(fmt.Sprintf("%v", r.ContributorLevel))
	builder.WriteString(", ")
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", r.Num))
	builder.WriteString(", ")
	builder.WriteString("formula=")
	builder.WriteString(r.Formula)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(r.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Rewards is a parsable slice of Reward.
type Rewards []*Reward
