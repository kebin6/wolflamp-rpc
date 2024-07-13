package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Statement 账单
type Statement struct {
	ent.Schema
}

// Fields of the Statement.
func (Statement) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("player_id").Default(0).
			Comment("Player 's id | 玩家ID，平台默认0").
			Annotations(entsql.WithComments(true)),
		field.Uint32("module").Default(0).
			Comment("Statement Module | 所属模块").
			Annotations(entsql.WithComments(true)),
		field.String("code").Default("").
			Comment("statement 's code | 账单编号").
			Annotations(entsql.WithComments(true)),
		field.Uint32("inout_type").Default(1).
			Comment("input or output | 收支类型：1=入账；2=出账").
			Annotations(entsql.WithComments(true)),
		field.Float("amount").Default(0).
			Comment("amount | 金额").
			Annotations(entsql.WithComments(true)),
		field.String("refer_id").Default("").
			Comment("refer id | 对应单ID").
			Annotations(entsql.WithComments(true)),
		field.String("mode").Default("").
			Comment("币种类别：coin/token").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Default("").
			Comment("remark | 备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Statement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("player_id").StorageKey("idx_player_id").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("玩家ID")),
		index.Fields("code").StorageKey("uniq_code").Unique().
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("账单号")),
		index.Fields("refer_id").StorageKey("idx_refer").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("对应单ID")),
	}
}

// Edges of the Statement.
func (Statement) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Statement) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_statement", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("账单表"),
	}
}

// Mixin of the Statement.
func (Statement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
