package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UserDevice holds the schema definition for the User entity.
type UserDevice struct {
	ent.Schema
}

// Fields of the UserDevice.
func (UserDevice) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("user_id").Unique(),
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

		field.Enum("status").
			Values("ACTIVE", "INACTIVE", "DELETED").
			Default("ACTIVE").
			Comment("해당 사용자는 서비스에서 유효한지 아닌지"),

		field.Enum("os").
			Values("MAC", "ANDROID", "WINDOWS").
			Nillable().
			Comment("해당 사용자는 서비스에서 유효한지 아닌지"),

		field.Enum("device").
			Values("DESKTOP", "MOBILE").
			Nillable().
			Comment("해당 사용자는 서비스에서 유효한지 아닌지"),

		field.Enum("type").
			Values("safari", "chrome", "DELETED").
			Nillable().
			Comment("해당 사용자는 서비스에서 유효한지 아닌지"),

		field.String("profile").
			Comment("어떤 Oauth provider를 이용하였는지?"),
	}
}

// Edge of the UserDevice
func (UserDevice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("meta").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}

// Indexes of the UserDevice
func (UserDevice) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.
			Fields("user_id").
			Unique(),
	}
}

// Annotations of the UserDevice.
func (UserDevice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}
