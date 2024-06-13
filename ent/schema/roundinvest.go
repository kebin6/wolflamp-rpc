package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// RoundInvest 游戏玩家投注数据
type RoundInvest struct {
	ent.Schema
}

// Fields of the RoundInvest.
func (RoundInvest) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("player_id").
			Default(0).
			Comment("玩家ID").
			Annotations(entsql.WithComments(true)),
		field.String("player_email").
			Default("").
			Comment("玩家邮箱").
			Annotations(entsql.WithComments(true)),
		field.Uint32("fold_no").
			Default(0).
			Comment("羊圈号码").
			Annotations(entsql.WithComments(true)),
		field.Uint32("lamb_num").
			Default(0).
			Comment("投注羊数量").
			Annotations(entsql.WithComments(true)),
		field.Float32("profit_and_loss").
			Default(0).
			Comment("投注盈亏结果").
			Annotations(entsql.WithComments(true)),
		field.Uint64("round_id").
			Optional().
			Default(0).
			Comment("所属回合ID").
			Annotations(entsql.WithComments(true)),
		field.Uint32("round_count").
			Optional().
			Default(0).
			Comment("当日第几回合").
			Annotations(entsql.WithComments(true)),
	}
}

func (RoundInvest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("round_id").StorageKey("idx_round").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("回合ID")),
		index.Fields("player_id").StorageKey("idx_player").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("玩家ID")),
	}
}

// Edges of the RoundInvest.
func (RoundInvest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("round", Round.Type).Ref("invest").Field("round_id").Unique(),
	}
}

func (RoundInvest) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_round_invest", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("回合投注情况表"),
	}

}

// Mixin of the RoundInvest.
func (RoundInvest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}
