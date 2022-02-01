package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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

		field.Enum("status").
			Values("ACTIVE", "INACTIVE", "DELETED").
			Default("ACTIVE").
			Comment("해당 사용자는 서비스에서 유효한지 아닌지"),

		field.Enum("type").
			Values("USER", "ADMIN").
			Default("USER").
			Comment("일반 사용자인지, 어드민 인지, 스태프 인지.. 등"),

		field.Bool("signup").
			Default(false).
			Comment("회원가입이 잘 끝난 사용자인지 아닌지?"),

		field.String("username").
			Unique().
			Comment("서비스에서 유일한 사용자의 이름"),
	}
}

//Edge of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("oauth_providers", UserOAuthProvider.Type),
		edge.To("meta", UserMeta.Type),
	}
}

// Annotations of the Entity.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}
