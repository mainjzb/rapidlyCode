package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("type").NotEmpty(),
		field.String("job_advancement"),
		field.String("icon").Optional(),
		field.String("icon_name").Optional(),
		field.Int("require_level").Optional(),
		field.String("detail").Optional(),
		field.Int("max_level"),
		field.String("description").Optional(),
		field.String("mechanics_level").Optional(),
		field.String("mechanics_detail").Optional(),
		// field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return nil
}
