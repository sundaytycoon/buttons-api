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
			Comment("해당 row는 쓸 수 있는지? 없는지?"),

		field.Enum("type").
			Values("DESKTOP", "MOBILE").
			Comment("모바일 사용자인지, 데스크탑 사용자인지?"),

		field.Enum("os").
			Values("MAC", "ANDROID", "WINDOWS").
			Comment("어떤 운영체제에서 접속 했는지?"),

		field.Enum("platform").
			Values("native", "chrome", "safari", "explorer").
			Comment("브라우저라면, 어떤 브라우저에서 사용하고 있는지?"),
	}
}

// Edges of the UserDevice
func (UserDevice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("devices").
			Field("id").Unique(),
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
