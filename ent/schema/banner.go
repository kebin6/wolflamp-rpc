package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Banner 轮播图
type Banner struct {
	ent.Schema
}

// Fields of the Banner.
func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.String("endpoint").Default("").
			Comment("Endpoint | 使用端").
			Annotations(entsql.WithComments(true)),
		field.String("module").Default("").
			Comment("Module | 模块").
			Annotations(entsql.WithComments(true)),
		field.Uint8("file_type").Default(2).
			Comment("File's type | 文件类型").
			Annotations(entsql.WithComments(true)),
		field.String("path").Default("").
			Comment("File path | 文件路径").
			Annotations(entsql.WithComments(true)),
		field.String("jump_url").Default("").
			Comment("Jump url | 跳转地址").
			Annotations(entsql.WithComments(true)).Optional(),
	}
}

func (Banner) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the Banner.
func (Banner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file", File.Type).Ref("banner").Unique().Required(),
	}
}

func (Banner) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_banner", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("轮播图配置表"),
	}
}

// Mixin of the Banner.
func (Banner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
