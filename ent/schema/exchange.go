package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Exchange 兑换
type Exchange struct {
	ent.Schema
}

// Fields of the Exchange.
func (Exchange) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("player_id").Default(0).
			Comment("player 's id | 兑换者ID").
			Annotations(entsql.WithComments(true)),
		field.String("transaction_id").Default("").
			Comment("transaction id | 交易ID").
			Annotations(entsql.WithComments(true)),
		field.String("mode").Default("").
			Comment("mode | 交易币种：coin,token").
			Annotations(entsql.WithComments(true)),
		field.Uint32("type").Default(0).
			Comment("type | 兑换类型：1=币兑羊；2=羊兑币").
			Annotations(entsql.WithComments(true)),
		field.Uint32("coin_num").Default(0).
			Comment("coin amount| 币数量").
			Annotations(entsql.WithComments(true)),
		field.Uint32("lamp_num").Default(0).
			Comment("lamp amount | 羊数量").
			Annotations(entsql.WithComments(true)),
		field.String("gcics_order_id").Default("").
			Comment("GCICS平台订单号 ").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Default("").
			Comment("备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Exchange) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("player_id").StorageKey("idx_player_id").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("兑换者ID")),
	}
}

// Edges of the Exchange.
func (Exchange) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Exchange) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_exchange", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("兑换表"),
	}
}

// Mixin of the Exchange.
func (Exchange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
