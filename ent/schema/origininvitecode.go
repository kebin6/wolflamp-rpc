package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// OriginInviteCode 原始邀请码
type OriginInviteCode struct {
	ent.Schema
}

// Fields of the OriginInviteCode.
func (OriginInviteCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Default("").
			Comment("邀请码").
			Annotations(entsql.WithComments(true)),
	}
}

func (OriginInviteCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("uniq_invite_code").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("邀请码")),
	}
}

// Edges of the File.
func (OriginInviteCode) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OriginInviteCode) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_origin_invite_code", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("原始邀请码管理表"),
	}
}

// Mixin of the OriginInviteCode.
func (OriginInviteCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
