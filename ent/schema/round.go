package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Round 游戏回合数据
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("round_count").
			Immutable().
			Comment("Round Count | 当日累计回合数").
			Annotations(entsql.WithComments(true)),
		field.Time("start_at").
			Immutable().
			Comment("Start Time | 回合开始时间（包括倒计时）").
			Annotations(entsql.WithComments(true)),
		field.Time("open_at").
			Immutable().
			Comment("Open Time | 回合开奖时间").
			Annotations(entsql.WithComments(true)),
		field.Time("end_at").
			Immutable().
			Comment("End Time | 回合结束时间").
			Annotations(entsql.WithComments(true)),
		field.Uint32("selected_fold").
			Default(0).
			Comment("选中的羊圈号码").
			Annotations(entsql.WithComments(true)),
	}
}

func (Round) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fold", RoundLambFold.Type),
		edge.To("invest", RoundInvest.Type),
	}
}

func (Round) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_round", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("回合表"),
	}

}

// Mixin of the Round.
func (Round) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
