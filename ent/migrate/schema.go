// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "data", Type: field.TypeJSON},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE", "DELETED"}, Default: "ACTIVE"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"USER", "ADMIN"}, Default: "USER"},
		{Name: "signup", Type: field.TypeBool, Default: false},
		{Name: "username", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserDevicesColumns holds the columns for the "user_devices" table.
	UserDevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE", "DELETED"}, Default: "ACTIVE"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"DESKTOP", "MOBILE"}},
		{Name: "os", Type: field.TypeEnum, Enums: []string{"MAC", "ANDROID", "WINDOWS"}},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"native", "chrome", "safari", "explorer"}},
		{Name: "user_device", Type: field.TypeString, Nullable: true},
	}
	// UserDevicesTable holds the schema information for the "user_devices" table.
	UserDevicesTable = &schema.Table{
		Name:       "user_devices",
		Columns:    UserDevicesColumns,
		PrimaryKey: []*schema.Column{UserDevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_devices_users_device",
				Columns:    []*schema.Column{UserDevicesColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userdevice_user_id",
				Unique:  true,
				Columns: []*schema.Column{UserDevicesColumns[1]},
			},
		},
	}
	// UserMetaColumns holds the columns for the "user_meta" table.
	UserMetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "profile", Type: field.TypeString},
		{Name: "user_meta", Type: field.TypeString, Nullable: true},
	}
	// UserMetaTable holds the schema information for the "user_meta" table.
	UserMetaTable = &schema.Table{
		Name:       "user_meta",
		Columns:    UserMetaColumns,
		PrimaryKey: []*schema.Column{UserMetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_meta_users_meta",
				Columns:    []*schema.Column{UserMetaColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usermeta_user_id",
				Unique:  true,
				Columns: []*schema.Column{UserMetaColumns[1]},
			},
		},
	}
	// UserOauthProvidersColumns holds the columns for the "user_oauth_providers" table.
	UserOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE", "DELETED"}, Default: "ACTIVE"},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"GOOGLE", "KAKAO"}},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "access_token", Type: field.TypeString, Unique: true},
		{Name: "refresh_token", Type: field.TypeString, Unique: true},
		{Name: "user_oauth_providers", Type: field.TypeString, Nullable: true},
	}
	// UserOauthProvidersTable holds the schema information for the "user_oauth_providers" table.
	UserOauthProvidersTable = &schema.Table{
		Name:       "user_oauth_providers",
		Columns:    UserOauthProvidersColumns,
		PrimaryKey: []*schema.Column{UserOauthProvidersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_oauth_providers_users_oauth_providers",
				Columns:    []*schema.Column{UserOauthProvidersColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "useroauthprovider_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserOauthProvidersColumns[1]},
			},
			{
				Name:    "useroauthprovider_status_provider_access_token",
				Unique:  true,
				Columns: []*schema.Column{UserOauthProvidersColumns[6], UserOauthProvidersColumns[7], UserOauthProvidersColumns[9]},
			},
			{
				Name:    "useroauthprovider_provider_refresh_token",
				Unique:  true,
				Columns: []*schema.Column{UserOauthProvidersColumns[7], UserOauthProvidersColumns[10]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsTable,
		UsersTable,
		UserDevicesTable,
		UserMetaTable,
		UserOauthProvidersTable,
	}
)

func init() {
	SessionsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	UserDevicesTable.ForeignKeys[0].RefTable = UsersTable
	UserDevicesTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	UserMetaTable.ForeignKeys[0].RefTable = UsersTable
	UserMetaTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	UserOauthProvidersTable.ForeignKeys[0].RefTable = UsersTable
	UserOauthProvidersTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
}
