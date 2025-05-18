package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Category хранит информацию о категории.
type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Comment("Название").Unique(),
		field.String("description").NotEmpty().Comment("Описание"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Category) Edges() []ent.Edge {
	return nil
}
