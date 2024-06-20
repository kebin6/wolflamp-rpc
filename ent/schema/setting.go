package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Setting 系统设置
type Setting struct {
	ent.Schema
}

var withComment = true

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").
			Comment("show name | 配置名称").
			Annotations(entsql.WithComments(true)),
		field.String("module").Default("").
			Comment("module | 模块").
			Annotations(entsql.WithComments(true)),
		field.Text("value").Default("").Nillable().
			Comment("value of setting | 配置值"),
	}
}

func (Setting) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the File.
func (Setting) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Setting) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_setting", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("系统配置表"),
	}
}

// Mixin of the Setting.
func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
