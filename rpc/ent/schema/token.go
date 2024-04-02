package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"

	"github.com/toutmost/admin-common/orm/ent/mixins"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Comment(" User's UUID | 用户的UUID").
			Annotations(entsql.WithComments(true)),
		field.String("username").
			Comment("Username | 用户名").
			Default("unknown").
			Annotations(entsql.WithComments(true)),
		field.String("token").
			Comment("Token string | Token 字符串").
			Annotations(entsql.WithComments(true)),
		field.String("source").
			Comment("Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）").
			Annotations(entsql.WithComments(true)),
		field.Time("expired_at").
			Comment(" Expire time | 过期时间").
			Annotations(entsql.WithComments(true)),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Token) Edges() []ent.Edge {
	return nil
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid"),
		index.Fields("expired_at"),
	}
}

func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_tokens"},
	}
}
