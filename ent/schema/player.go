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

// Player 玩家信息
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("").
			Comment("player name | 用户名称").
			Annotations(entsql.WithComments(true)),
		field.String("email").
			Default("").
			Comment("email | 邮箱").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Default("").
			Comment("password | 密码").
			Annotations(entsql.WithComments(true)),
		field.String("transaction_password").
			Default("").
			Comment("transaction password | 交易密码").
			Annotations(entsql.WithComments(true)),
		field.Float32("lamp").
			Default(0).
			Comment("lamp | 小羊余额").
			Annotations(entsql.WithComments(true)),
		field.Uint32("rank").
			Default(0).
			Comment("rank | 代理等级").
			Annotations(entsql.WithComments(true)),
		field.Float("amount").
			Default(0.00).
			Comment("amount | 账户余额").
			Annotations(entsql.WithComments(true)),
		field.String("deposit_address").
			Default("").
			Comment("deposit address | 充值地址").
			Annotations(entsql.WithComments(true)),
		field.Uint32("invited_num").
			Default(0).
			Comment("invite people | 代理发展人数").
			Annotations(entsql.WithComments(true)),
		field.Float("total_income").
			Default(0.00).
			Comment("total income | 代理总收益").
			Annotations(entsql.WithComments(true)),
		field.Float32("profit_and_loss").
			Default(0).
			Comment("profit and loss | 总盈亏数").
			Annotations(entsql.WithComments(true)),
		field.Float32("recent_100_win_percent").
			Default(0).
			Comment("recent 100 win percent | 近100场胜率").
			Annotations(entsql.WithComments(true)),
		field.String("invite_code").
			Default("").
			Comment("invite code | 邀请码").
			Annotations(entsql.WithComments(true)),
		field.Uint64("inviter_id").Optional().
			Default(0).
			Comment("the inviter id | 邀请人ID").
			Annotations(entsql.WithComments(true)),
		field.String("invited_code").
			Default("").
			Comment("the inviter 's invite code | 邀请人邀请码").
			Annotations(entsql.WithComments(true)),
		field.Float32("system_commission").
			Default(0).
			Comment("system commission percent | 平台收益分成比例").
			Annotations(entsql.WithComments(true)),
	}
}

func (Player) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique().StorageKey("uniq_email").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("邮箱")),
		index.Fields("invite_code").Unique().StorageKey("uniq_invite_code").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("邀请码")),
		index.Fields("inviter_id").StorageKey("idx_inviter_id").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("邀请人ID")),
		index.Fields("invited_code").StorageKey("idx_invited_code").
			Annotations(entsql.Annotation{WithComments: &withComment}, schema.Comment("邀请人邀请码")),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		// 用户邀请到的人
		edge.To("invitees", Player.Type).
			From("inviter").
			Unique().
			Field("inviter_id"),
	}
}

func (Player) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "wl_player", WithComments: &withComment,
			Charset: "utf8mb4", Collation: "utf8mb4_general_ci"},
		schema.Comment("玩家表"),
	}
}

// Mixin of the Player.
func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}
