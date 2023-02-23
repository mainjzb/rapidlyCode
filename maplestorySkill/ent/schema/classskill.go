package schema

import "entgo.io/ent"

// ClassSkill holds the schema definition for the ClassSkill entity.
type ClassSkill struct {
	ent.Schema
}

// Fields of the ClassSkill.
func (ClassSkill) Fields() []ent.Field {
	return nil
}

// Edges of the ClassSkill.
func (ClassSkill) Edges() []ent.Edge {
	return nil
}
