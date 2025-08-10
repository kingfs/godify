package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// Note: The following structs are duplicates from the web package.
// They should be refactored into a common types package.

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

// Limitation represents a usage and limit for a feature.
type Limitation struct {
	Size  int `json:"size"`
	Limit int `json:"limit"`
}

// Subscription represents the billing subscription details.
type Subscription struct {
	Plan     string `json:"plan"`
	Interval string `json:"interval"`
}

// Billing represents the billing status and features.
type Billing struct {
	Enabled      bool         `json:"enabled"`
	Subscription Subscription `json:"subscription"`
}

// Education represents the education license status.
type Education struct {
	Enabled   bool `json:"enabled"`
	Activated bool `json:"activated"`
}

// Features represents the feature flags and limits for a tenant.
type Features struct {
	Billing                   Billing           `json:"billing"`
	Education                 Education         `json:"education"`
	Members                   Limitation        `json:"members"`
	Apps                      Limitation        `json:"apps"`
	VectorSpace               Limitation        `json:"vector_space"`
	KnowledgeRateLimit        int               `json:"knowledge_rate_limit"`
	AnnotationQuotaLimit      Limitation        `json:"annotation_quota_limit"`
	DocumentsUploadQuota      Limitation        `json:"documents_upload_quota"`
	DocsProcessing            string            `json:"docs_processing"`
	CanReplaceLogo            bool              `json:"can_replace_logo"`
	ModelLoadBalancingEnabled bool              `json:"model_load_balancing_enabled"`
	DatasetOperatorEnabled    bool              `json:"dataset_operator_enabled"`
	WebappCopyrightEnabled    bool              `json:"webapp_copyright_enabled"`
	WorkspaceMembers          LicenseLimitation `json:"workspace_members"`
	IsAllowTransferWorkspace  bool              `json:"is_allow_transfer_workspace"`
}

// GetFeatures retrieves the feature flags and limits for the current tenant.
func (c *client.Client) GetFeatures(ctx context.Context) (*Features, error) {
	var result Features
	err := c.sendRequest(ctx, "GET", "/console/api/features", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


// ConsoleGetSystemFeatures retrieves the system-level feature flags.
func (c *client.Client) ConsoleGetSystemFeatures(ctx context.Context) (*SystemFeatures, error) {
	var result SystemFeatures
	err := c.sendRequest(ctx, "GET", "/console/api/system-features", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
