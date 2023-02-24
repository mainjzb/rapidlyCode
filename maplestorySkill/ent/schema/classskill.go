package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ClassSkill holds the schema definition for the ClassSkill entity.
type ClassSkill struct {
	ent.Schema
}

// Fields of the ClassSkill.
func (ClassSkill) Fields() []ent.Field {
	return []ent.Field{
		field.String("class").NotEmpty(),
		field.String("job_advancement"),
		field.String("detail").Optional(),
	}
}

// Edges of the ClassSkill.
func (ClassSkill) Edges() []ent.Edge {
	return nil
}
