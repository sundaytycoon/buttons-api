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

// UserOAuthProvider holds the schema definition for the UserOAuthProvider entity.
type UserOAuthProvider struct {
	ent.Schema
}

// Fields of the UserOAuthProvider.
func (UserOAuthProvider) Fields() []ent.Field {
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

		field.Enum("provider").
			Values("GOOGLE", "KAKAO").
			Comment("어떤 Oauth provider를 이용하였는지?"),

		field.Time("expiry").
			Comment("해당 토큰의 만료시간이 언제까지인지?"),

		field.String("access_token").
			Unique().
			Comment("access_token"),

		field.String("refresh_token").
			Unique().
			Comment("refresh_token"),
	}
}

// Edges of the UserOAuthProvider
func (UserOAuthProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("oauth_providers").
			Field("user_id").
			Required().
			Unique(),
	}
}

// Indexes of the UserOAuthProvider
func (UserOAuthProvider) Indexes() []ent.Index {
	return []ent.Index{
		// none index. 1: n
		index.
			Fields("user_id"),

		// unique index.
		index.
			Fields("status", "provider", "access_token").
			Unique(),
		// unique index.
		index.
			Fields("provider", "refresh_token").
			Unique(),
	}
}

// Annotations of the Entity.
func (UserOAuthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}
