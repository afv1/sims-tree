package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("login").
            Unique(),
        field.String("password"),
        field.String("comment").
            Optional(),
        field.Time("created_at").
            Default(time.Now),
//        field.Time("deleted_at").Optional().Nillable(),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("families", Family.Type),
    }
}
