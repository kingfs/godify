package console

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// ConsoleLoginRequest is the request for console email/password login.
type ConsoleLoginRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me,omitempty"`
}

// LoginResponse is the response after a successful login, containing an access token.
type LoginResponse struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

// TokenPairResponse is the response containing both access and refresh tokens.
type TokenPairResponse struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

// SendEmailCodeRequest is the request to send a login code via email.
type SendEmailCodeRequest struct {
	Email    string `json:"email"`
	Language string `json:"language,omitempty"`
}

// SendEmailCodeResponse is the response after sending a login code.
type SendEmailCodeResponse struct {
	Result string `json:"result"`
	Data   string `json:"data"` // This is the token
}

// CheckEmailCodeRequest is the request to log in using an email code.
type CheckEmailCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Token string `json:"token"`
}

// RefreshTokenRequest is the request to refresh an access token.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// LogoutResponse is the response from the logout endpoint.
type LogoutResponse struct {
	Result string `json:"result"`
}

// Login performs console email and password authentication.
func (c *client.Client) ConsoleLogin(ctx context.Context, req *ConsoleLoginRequest) (*TokenPairResponse, error) {
	var result TokenPairResponse
	err := c.sendRequest(ctx, "POST", "/console/api/login", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Logout logs out the current user.
func (c *client.Client) Logout(ctx context.Context) (*LogoutResponse, error) {
	var result LogoutResponse
	err := c.sendRequest(ctx, "GET", "/console/api/logout", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


// ConsoleSendLoginEmailCode sends a one-time login code to the user's email.
func (c *client.Client) ConsoleSendLoginEmailCode(ctx context.Context, req *SendEmailCodeRequest) (*SendEmailCodeResponse, error) {
	var result SendEmailCodeResponse
	err := c.sendRequest(ctx, "POST", "/console/api/email-code-login", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleLoginWithEmailCode performs login using the one-time code.
func (c *client.Client) ConsoleLoginWithEmailCode(ctx context.Context, req *CheckEmailCodeRequest) (*TokenPairResponse, error) {
	var result TokenPairResponse
	err := c.sendRequest(ctx, "POST", "/console/api/email-code-login/validity", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleSendResetPasswordEmail sends a reset password email.
func (c *client.Client) ConsoleSendResetPasswordEmail(ctx context.Context, req *SendEmailCodeRequest) (*SendEmailCodeResponse, error) {
	var result SendEmailCodeResponse
	err := c.sendRequest(ctx, "POST", "/console/api/reset-password", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RefreshToken refreshes an access token.
func (c *client.Client) RefreshToken(ctx context.Context, refreshToken string) (*TokenPairResponse, error) {
	var result TokenPairResponse
	req := RefreshTokenRequest{RefreshToken: refreshToken}
	err := c.sendRequest(ctx, "POST", "/console/api/refresh-token", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
