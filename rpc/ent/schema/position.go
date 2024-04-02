package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/toutmost/admin-common/orm/ent/mixins"
)

type Position struct {
	ent.Schema
}

func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Position Name | 职位名称").
			Annotations(entsql.WithComments(true)),
		field.String("code").
			Comment("The code of position | 职位编码").
			Annotations(entsql.WithComments(true)),
		field.String("remark").Optional().
			Comment("Remark | 备注").
			Annotations(entsql.WithComments(true)),
	}
}

func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("positions"),
	}
}

func (Position) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique(),
	}
}

func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_positions"},
	}
}
