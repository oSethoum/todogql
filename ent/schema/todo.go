package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Annotations(entgql.OrderField("TEXT")),
		field.String("description").Annotations(entgql.OrderField("DESCRIPTION")),
		field.Bool("done").Annotations(entgql.OrderField("DONE")).Default(false),
	}
}

// Mixins of the Todo
func (Todo) Mixins() []ent.Mixin {
	return []ent.Mixin{}
}

// Edges of the Todo
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Unique(),
	}
}

// Annotations of the Todo
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
