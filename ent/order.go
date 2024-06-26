// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/wolflamp-rpc/ent/order"
)

// 订单表
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// player id | 玩家ID
	PlayerID uint64 `json:"player_id,omitempty"`
	// order type | 订单类型: deposit=充值订单；withdraw=提现订单
	Type string `json:"type,omitempty"`
	// order code | 订单号
	Code string `json:"code,omitempty"`
	// transaction is | 交易号
	TransactionID string `json:"transaction_id,omitempty"`
	// transfer from address | 汇出钱包地址
	FromAddress string `json:"from_address,omitempty"`
	// transfer to address | 汇入钱包地址
	ToAddress string `json:"to_address,omitempty"`
	// coin amount | 数量
	Num float64 `json:"num,omitempty"`
	// handling_fee | 对提币操作生效，提币手续费
	HandlingFee float64 `json:"handling_fee,omitempty"`
	// 网络
	Network string `json:"network,omitempty"`
	// remark | 备注
	Remark       string `json:"remark,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldNum, order.FieldHandlingFee:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldStatus, order.FieldPlayerID:
			values[i] = new(sql.NullInt64)
		case order.FieldType, order.FieldCode, order.FieldTransactionID, order.FieldFromAddress, order.FieldToAddress, order.FieldNetwork, order.FieldRemark:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = uint64(value.Int64)
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = uint8(value.Int64)
			}
		case order.FieldPlayerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field player_id", values[i])
			} else if value.Valid {
				o.PlayerID = uint64(value.Int64)
			}
		case order.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = value.String
			}
		case order.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				o.Code = value.String
			}
		case order.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				o.TransactionID = value.String
			}
		case order.FieldFromAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from_address", values[i])
			} else if value.Valid {
				o.FromAddress = value.String
			}
		case order.FieldToAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to_address", values[i])
			} else if value.Valid {
				o.ToAddress = value.String
			}
		case order.FieldNum:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				o.Num = value.Float64
			}
		case order.FieldHandlingFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field handling_fee", values[i])
			} else if value.Valid {
				o.HandlingFee = value.Float64
			}
		case order.FieldNetwork:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field network", values[i])
			} else if value.Valid {
				o.Network = value.String
			}
		case order.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				o.Remark = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("player_id=")
	builder.WriteString(fmt.Sprintf("%v", o.PlayerID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(o.Type)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(o.Code)
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(o.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("from_address=")
	builder.WriteString(o.FromAddress)
	builder.WriteString(", ")
	builder.WriteString("to_address=")
	builder.WriteString(o.ToAddress)
	builder.WriteString(", ")
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", o.Num))
	builder.WriteString(", ")
	builder.WriteString("handling_fee=")
	builder.WriteString(fmt.Sprintf("%v", o.HandlingFee))
	builder.WriteString(", ")
	builder.WriteString("network=")
	builder.WriteString(o.Network)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(o.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
