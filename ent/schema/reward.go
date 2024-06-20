package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Reward 奖励
type Reward struct {
	ent.Schema
}

// Fields of the Reward.
func (Reward) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("to_id").Default(0).
			Comment("reward to player 's id | 奖励对象ID").
			Annotations(entsql.WithComments(true)),
		field.Uint64("contributor_id").Default(0).
			Comment("contributor 's id | 贡献者ID,平台为0").
			Annotations(entsql.WithComments(true)),
		field.String("contributor_email").Default("").
			Comment("contributor 's email | 贡献者邮箱，平台为system").
			Annotations(entsql.WithComments(true)),
		field.Uint32("contributor_level").Default(1).
			Comment("contributor level | 贡献者属于奖励对象第几级").
			Annotations(entsql.WithComments(true)),
		field.Float32("num").Default(0).
			Comment("lamb amount | 数量").
			Annotations(entsql.WithComments(true)),
		field.String("formula").Default("").
			Comment("计算公式").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Default("").
			Comment("remark | 备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Reward) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("to_id").StorageKey("idx_to").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("奖励对象ID")),
	}
}

// Edges of the Reward.
func (Reward) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Reward) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_reward", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("奖励表"),
	}
}

// Mixin of the Reward.
func (Reward) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
