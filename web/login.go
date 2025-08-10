package web

import (
	"context"

	"github.com/kingfs/godify/client"
)

// LoginRequest is the request for email/password login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is the response after a successful login.
type LoginResponse struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken string `json:"access_token"`
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

// CheckEmailCodeResponse is the response after a successful email code login.
type CheckEmailCodeResponse struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

// Login performs email and password authentication.
func (c *client.Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var result LoginResponse
	err := c.sendRequest(ctx, "POST", "/api/login", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SendLoginEmailCode sends a one-time login code to the user's email.
func (c *client.Client) SendLoginEmailCode(ctx context.Context, req *SendEmailCodeRequest) (*SendEmailCodeResponse, error) {
	var result SendEmailCodeResponse
	err := c.sendRequest(ctx, "POST", "/api/email-code-login", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// LoginWithEmailCode performs login using the one-time code.
func (c *client.Client) LoginWithEmailCode(ctx context.Context, req *CheckEmailCodeRequest) (*CheckEmailCodeResponse, error) {
	var result CheckEmailCodeResponse
	err := c.sendRequest(ctx, "POST", "/api/email-code-login/validity", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
