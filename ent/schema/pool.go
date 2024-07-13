package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Pool 资金池
type Pool struct {
	ent.Schema
}

// Fields of the Exchange.
func (Pool) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("round_id").Default(0).
			Comment("回合ID").
			Annotations(entsql.WithComments(true)),
		field.String("mode").Default("").
			Comment("mode | 交易币种：coin,token").
			Annotations(entsql.WithComments(true)),
		field.Uint32("type").Default(0).
			Comment("type | 池类型：1=机器人资金池；2=奖金蓄水池；3=其他").
			Annotations(entsql.WithComments(true)),
		field.Float("lamb_num").Default(0).
			Comment("lamp amount | 羊数量").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Default("").
			Comment("备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Pool) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the Exchange.
func (Pool) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Pool) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_pool", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("资金池表"),
	}
}

// Mixin of the Exchange.
func (Pool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
