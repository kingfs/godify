package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// Account represents a user account.
type Account struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Avatar             string `json:"avatar"`
	Status             string `json:"status"`
	InterfaceLanguage  string `json:"interface_language"`
	InterfaceTheme     string `json:"interface_theme"`
	Timezone           string `json:"timezone"`
	LastLoginAt        int64  `json:"last_login_at"`
	LastLoginIP        string `json:"last_login_ip"`
	Created_at         int64  `json:"created_at"`
}

// InitAccountRequest is the request to initialize an account.
type InitAccountRequest struct {
	InvitationCode    string `json:"invitation_code,omitempty"`
	InterfaceLanguage string `json:"interface_language"`
	Timezone          string `json:"timezone"`
}

// UpdateNameRequest is the request to update the account name.
type UpdateNameRequest struct {
	Name string `json:"name"`
}

// UpdateAvatarRequest is the request to update the account avatar.
type UpdateAvatarRequest struct {
	Avatar string `json:"avatar"`
}

// UpdateInterfaceLanguageRequest is the request to update the interface language.
type UpdateInterfaceLanguageRequest struct {
	InterfaceLanguage string `json:"interface_language"`
}

// UpdateInterfaceThemeRequest is the request to update the interface theme.
type UpdateInterfaceThemeRequest struct {
	InterfaceTheme string `json:"interface_theme"`
}

// UpdateTimezoneRequest is the request to update the timezone.
type UpdateTimezoneRequest struct {
	Timezone string `json:"timezone"`
}

// UpdatePasswordRequest is the request to update the password.
type UpdatePasswordRequest struct {
	Password          string `json:"password"`
	NewPassword       string `json:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password"`
}

// ... and many more structs and functions for this large API file.

// InitAccount initializes a new user account.
func (c *client.Client) InitAccount(ctx context.Context, req *InitAccountRequest) error {
	return c.sendRequest(ctx, "POST", "/console/api/account/init", req, nil, nil)
}

// GetAccountProfile retrieves the profile of the current user.
func (c *client.Client) GetAccountProfile(ctx context.Context) (*Account, error) {
	var result Account
	err := c.sendRequest(ctx, "GET", "/console/api/account/profile", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountName updates the name of the current user.
func (c *client.Client) UpdateAccountName(ctx context.Context, name string) (*Account, error) {
	var result Account
	req := UpdateNameRequest{Name: name}
	err := c.sendRequest(ctx, "POST", "/console/api/account/name", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountAvatar updates the avatar of the current user.
func (c *client.Client) UpdateAccountAvatar(ctx context.Context, avatar string) (*Account, error) {
	var result Account
	req := UpdateAvatarRequest{Avatar: avatar}
	err := c.sendRequest(ctx, "POST", "/console/api/account/avatar", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountInterfaceLanguage updates the interface language of the current user.
func (c *client.Client) UpdateAccountInterfaceLanguage(ctx context.Context, language string) (*Account, error) {
	var result Account
	req := UpdateInterfaceLanguageRequest{InterfaceLanguage: language}
	err := c.sendRequest(ctx, "POST", "/console/api/account/interface-language", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountInterfaceTheme updates the interface theme of the current user.
func (c *client.Client) UpdateAccountInterfaceTheme(ctx context.Context, theme string) (*Account, error) {
	var result Account
	req := UpdateInterfaceThemeRequest{InterfaceTheme: theme}
	err := c.sendRequest(ctx, "POST", "/console/api/account/interface-theme", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountTimezone updates the timezone of the current user.
func (c *client.Client) UpdateAccountTimezone(ctx context.Context, timezone string) (*Account, error) {
	var result Account
	req := UpdateTimezoneRequest{Timezone: timezone}
	err := c.sendRequest(ctx, "POST", "/console/api/account/timezone", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccountPassword updates the password of the current user.
func (c *client.Client) UpdateAccountPassword(ctx context.Context, req *UpdatePasswordRequest) error {
	return c.sendRequest(ctx, "POST", "/console/api/account/password", req, nil, nil)
}

// AccountIntegrate represents an OAuth integration for an account.
type AccountIntegrate struct {
	ID        string `json:"id"`
	Provider  string `json:"provider"`
	CreatedAt int64  `json:"created_at"`
	IsBound   bool   `json:"is_bound"`
	Link      string `json:"link"`
}

// AccountIntegrateListResponse is the response for listing account integrations.
type AccountIntegrateListResponse struct {
	Data []AccountIntegrate `json:"data"`
}

// GetAccountIntegrates retrieves the list of OAuth integrations for the current user.
func (c *client.Client) GetAccountIntegrates(ctx context.Context) (*AccountIntegrateListResponse, error) {
	var result AccountIntegrateListResponse
	err := c.sendRequest(ctx, "GET", "/console/api/account/integrates", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAccountDeleteVerify sends a verification code for account deletion.
func (c *client.Client) GetAccountDeleteVerify(ctx context.Context) (string, error) {
	var result struct {
		Data string `json:"data"`
	}
	err := c.sendRequest(ctx, "GET", "/console/api/account/delete/verify", nil, &result, nil)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}

// DeleteAccountRequest is the request to delete an account.
type DeleteAccountRequest struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}

// DeleteAccount deletes the current user's account.
func (c *client.Client) DeleteAccount(ctx context.Context, req *DeleteAccountRequest) error {
	return c.sendRequest(ctx, "POST", "/console/api/account/delete", req, nil, nil)
}

// ChangeEmailSendCodeRequest is the request to send a code for changing email.
type ChangeEmailSendCodeRequest struct {
	Email    string `json:"email"`
	Language string `json:"language,omitempty"`
	Phase    string `json:"phase,omitempty"`
	Token    string `json:"token,omitempty"`
}

// ChangeEmailCheckCodeRequest is the request to check the code for changing email.
type ChangeEmailCheckCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Token string `json:"token"`
}

// ChangeEmailResetRequest is the request to reset the email.
type ChangeEmailResetRequest struct {
	NewEmail string `json:"new_email"`
	Token    string `json:"token"`
}

// SendChangeEmailCode sends a verification code for changing an email address.
func (c *client.Client) SendChangeEmailCode(ctx context.Context, req *ChangeEmailSendCodeRequest) (*SendEmailCodeResponse, error) {
	var result SendEmailCodeResponse
	err := c.sendRequest(ctx, "POST", "/console/api/account/change-email", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckChangeEmailCode checks the verification code for changing an email address.
func (c *client.Client) CheckChangeEmailCode(ctx context.Context, req *ChangeEmailCheckCodeRequest) (*CheckForgotPasswordCodeResponse, error) {
	var result CheckForgotPasswordCodeResponse
	err := c.sendRequest(ctx, "POST", "/console/api/account/change-email/validity", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetEmail changes the user's email address.
func (c *client.Client) ResetEmail(ctx context.Context, req *ChangeEmailResetRequest) (*Account, error) {
	var result Account
	err := c.sendRequest(ctx, "POST", "/console/api/account/change-email/reset", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckEmailUnique checks if an email is unique.
func (c *client.Client) CheckEmailUnique(ctx context.Context, email string) error {
	req := map[string]string{"email": email}
	return c.sendRequest(ctx, "POST", "/console/api/account/change-email/check-email-unique", req, nil, nil)
}
