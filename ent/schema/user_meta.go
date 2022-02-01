package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserMeta holds the schema definition for the UserOAuthProvider entity.
type UserMeta struct {
	//ent.Schema
}

// Fields of the UserMeta.
func (UserMeta) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("created_at").
			Default(time.Now).
			Comment("해당 row를 최초로 만든 시간은 언제인지?"),

		field.String("created_by").
			Comment("해당 row를 최초로 만든 주체는 누구인지?"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("해당 row를 수정한 마지막 시간은 언제인지"),
		field.String("updated_by").
			Comment("해당 row를 수정한 마지막 주체 누구인지"),

		field.String("profile").
			Comment("어떤 Oauth provider를 이용하였는지?"),
	}
}

// Edge of the UserMeta
func (UserMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("meta").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}

// Indexes of the UserMeta
func (UserMeta) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.
			Fields("user_id").
			Unique(),
	}
}

// Annotations of the UserMeta.
func (UserMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}
