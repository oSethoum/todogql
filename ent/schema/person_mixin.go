package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// schema of the PersonMixin
type PersonMixin struct {
	mixin.Schema
}

// Fields of the PersonMixin.
func (PersonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Annotations(entgql.OrderField("FIRST_NAME")),
		field.String("last_name").Annotations(entgql.OrderField("LAST_NAME")),
		field.Time("birthdate").Annotations(entgql.OrderField("BIRTHDATE")),
	}
}

// Edges of the PersonMixin
func (PersonMixin) Edges() []ent.Edge {
	return []ent.Edge{}
}
