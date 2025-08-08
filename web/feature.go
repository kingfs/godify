package web

import (
	"context"

	"github.com/kingfs/godify/client"
)

// LicenseLimitation represents the usage and limit for a licensed feature.
type LicenseLimitation struct {
	Enabled bool `json:"enabled"`
	Size    int  `json:"size"`
	Limit   int  `json:"limit"`
}

// License represents the software license information.
type License struct {
	Status     string            `json:"status"`
	ExpiredAt  string            `json:"expired_at"`
	Workspaces LicenseLimitation `json:"workspaces"`
}

// Branding represents the custom branding settings.
type Branding struct {
	Enabled          bool   `json:"enabled"`
	ApplicationTitle string `json:"application_title"`
	LoginPageLogo    string `json:"login_page_logo"`
	WorkspaceLogo    string `json:"workspace_logo"`
	Favicon          string `json:"favicon"`
}

// WebAppAuthSSO represents the SSO configuration for web apps.
type WebAppAuthSSO struct {
	Protocol string `json:"protocol"`
}

// WebAppAuth represents the authentication settings for web apps.
type WebAppAuth struct {
	Enabled                bool          `json:"enabled"`
	AllowSSO               bool          `json:"allow_sso"`
	SSOConfig              WebAppAuthSSO `json:"sso_config"`
	AllowEmailCodeLogin    bool          `json:"allow_email_code_login"`
	AllowEmailPasswordLogin bool          `json:"allow_email_password_login"`
}

// PluginInstallationPermission represents the plugin installation permissions.
type PluginInstallationPermission struct {
	PluginInstallationScope   string `json:"plugin_installation_scope"`
	RestrictToMarketplaceOnly bool   `json:"restrict_to_marketplace_only"`
}

// SystemFeatures represents the available system-level features.
type SystemFeatures struct {
	SSOEnforcedForSignin         bool                         `json:"sso_enforced_for_signin"`
	SSOEnforcedForSigninProtocol string                       `json:"sso_enforced_for_signin_protocol"`
	EnableMarketplace            bool                         `json:"enable_marketplace"`
	MaxPluginPackageSize         int                          `json:"max_plugin_package_size"`
	EnableEmailCodeLogin         bool                         `json:"enable_email_code_login"`
	EnableEmailPasswordLogin     bool                         `json:"enable_email_password_login"`
	EnableSocialOAuthLogin       bool                         `json:"enable_social_oauth_login"`
	IsAllowRegister              bool                         `json:"is_allow_register"`
	IsAllowCreateWorkspace       bool                         `json:"is_allow_create_workspace"`
	IsEmailSetup                 bool                         `json:"is_email_setup"`
	License                      License                      `json:"license"`
	Branding                     Branding                     `json:"branding"`
	WebAppAuth                   WebAppAuth                   `json:"webapp_auth"`
	PluginInstallationPermission PluginInstallationPermission `json:"plugin_installation_permission"`
	EnableChangeEmail            bool                         `json:"enable_change_email"`
}

// GetSystemFeatures retrieves the system-level feature flags.
func (c *client.Client) GetSystemFeatures(ctx context.Context) (*SystemFeatures, error) {
	var result SystemFeatures
	err := c.sendRequest(ctx, "GET", "/api/system-features", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
