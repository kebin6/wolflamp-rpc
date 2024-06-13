package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("File's name | 文件名称").
			Annotations(entsql.WithComments(true)),
		field.Uint8("file_type").
			Comment("File's type | 文件类型").
			Annotations(entsql.WithComments(true)),
		field.Uint64("size").
			Comment("File's size | 文件大小").
			Annotations(entsql.WithComments(true)),
		field.String("path").
			Comment("File's path | 文件路径"),
		field.String("user_id").
			Comment("User's UUID | 用户的 UUID").
			Annotations(entsql.WithComments(true)),
		field.String("md5").
			Comment("The md5 of the file | 文件的 md5").
			Annotations(entsql.WithComments(true)),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("banner", Banner.Type).Unique().StorageKey(edge.Column("file_id")),
	}
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fms_files"}, // fms means file management service
	}
}
