package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Family holds the schema definition for the Family entity.
type Family struct {
	ent.Schema
}

// Fields of the Family.
func (Family) Fields() []ent.Field {
    return []ent.Field{
        field.String("surname"),
        field.String("description"),
        field.String("logo_resource").
            Nillable(),
    }
}

// Edges of the Family.
func (Family) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("owner", User.Type).
            Ref("families").
            Unique(),
        edge.To("relatives", Relative.Type),
    }
}
