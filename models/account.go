package models

type Account struct {
	ID    string `json:"id"`    // base
	Name  string `json:"name"`  // base
	Email string `json:"email"` // base

	Avatar            *string   `json:"avatar,omitempty"`
	AvatarURL         *string   `json:"avatar_url,omitempty"`
	IsPasswordSet     *bool     `json:"is_password_set,omitempty"`
	InterfaceLanguage *string   `json:"interface_language,omitempty"`
	InterfaceTheme    *string   `json:"interface_theme,omitempty"`
	Timezone          *string   `json:"timezone,omitempty"`
	LastLoginAt       *UnixTime `json:"last_login_at,omitempty"`
	LastLoginIP       *string   `json:"last_login_ip,omitempty"`
	CreatedAt         *UnixTime `json:"created_at,omitempty"`

	Role         *string   `json:"role,omitempty"`           // account_with_role_fields
	Status       *string   `json:"status,omitempty"`         // account_with_role_fields
	LastActiveAt *UnixTime `json:"last_active_at,omitempty"` // account_with_role_fields
}
