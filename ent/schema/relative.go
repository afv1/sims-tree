package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Relative holds the schema definition for the Relative entity.
type Relative struct {
	ent.Schema
}

// Fields of the Relative.
func (Relative) Fields() []ent.Field {
	return []ent.Field{
        field.String("name").
            NotEmpty(),
        field.String("description").
            Nillable(),
        field.String("avatar_resource").
            Nillable(),
    }
}

// Edges of the Relative.
func (Relative) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("families", Family.Type).
            Ref("relatives").
            Unique(),
    }
}
