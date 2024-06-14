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

// RoundLambFold 游戏回合羊圈数据
type RoundLambFold struct {
	ent.Schema
}

// Fields of the RoundLambFold.
func (RoundLambFold) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("fold_no").
			Default(0).
			Comment("羊圈号").
			Annotations(entsql.WithComments(true)),
		field.Uint32("lamb_num").
			Default(0).
			Comment("投注羊数量").
			Annotations(entsql.WithComments(true)),
		field.Uint64("round_id").
			Optional().
			Default(0).
			Comment("所属回合ID").
			Annotations(entsql.WithComments(true)),
		field.Float32("profit_and_loss").
			Default(0).
			Comment("投注盈亏结果").
			Annotations(entsql.WithComments(true)),
		field.Uint32("round_count").
			Optional().
			Default(0).
			Comment("当日第几回合").
			Annotations(entsql.WithComments(true)),
		field.Uint64("total_round_count").
			Optional().
			Default(0).
			Comment("累计第几回合").
			Annotations(entsql.WithComments(true)),
	}
}

func (RoundLambFold) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("round_id").StorageKey("idx_round").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("回合ID")),
		index.Fields("total_round_count").StorageKey("idx_total_count").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("累计第几回合")),
	}
}

// Edges of the RoundLambFold.
func (RoundLambFold) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("round", Round.Type).Ref("fold").Field("round_id").Unique(),
	}
}

func (RoundLambFold) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_round_lamb_fold", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("回合羊圈情况表"),
	}

}

// Mixin of the RoundLambFold.
func (RoundLambFold) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}
