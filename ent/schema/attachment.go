package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(32).NotEmpty(),
		field.String("filename").NotEmpty(),
		field.String("filepath").NotEmpty(),
		field.String("contentType").NotEmpty(),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Envelope.Type).Ref("attachments").Unique(),
	}
}
