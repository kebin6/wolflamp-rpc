package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Order 订单
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("player_id").Default(0).
			Comment("player id | 玩家ID").
			Annotations(entsql.WithComments(true)),
		field.String("type").Default("").
			Comment("order type | 订单类型: deposit=充值订单；withdraw=提现订单").
			Annotations(entsql.WithComments(true)),
		field.String("code").Default("").
			Comment("order code | 订单号").
			Annotations(entsql.WithComments(true)),
		field.String("transaction_id").Default("").
			Comment("transaction is | 交易号").
			Annotations(entsql.WithComments(true)),
		field.String("from_address").Default("").
			Comment("transfer from address | 汇出钱包地址"),
		field.String("to_address").Default("").
			Comment("transfer to address | 汇入钱包地址").
			Annotations(entsql.WithComments(true)),
		field.Float("num").Default(0).
			Comment("coin amount | 数量").
			Annotations(entsql.WithComments(true)),
		field.Float("handling_fee").Default(0).
			Comment("handling_fee | 对提币操作生效，提币手续费").
			Annotations(entsql.WithComments(true)),
		field.String("network").Default("").
			Comment("网络").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Default("").
			Comment("remark | 备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("player_id").StorageKey("idx_player_id").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("玩家ID")),
		index.Fields("code").Unique().StorageKey("uniq_code").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("订单号")),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Order) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_order", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("订单表"),
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
