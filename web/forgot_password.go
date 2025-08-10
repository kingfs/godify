package web

import (
	"context"

	"github.com/kingfs/godify/client"
)

// ForgotPasswordRequest is the request to send a forgot password email.
type ForgotPasswordRequest struct {
	Email    string `json:"email"`
	Language string `json:"language,omitempty"`
}

// ForgotPasswordResponse is the response after sending a forgot password email.
type ForgotPasswordResponse struct {
	Result string `json:"result"`
	Data   string `json:"data"` // This is the token
}

// CheckForgotPasswordCodeRequest is the request to validate a forgot password code.
type CheckForgotPasswordCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Token string `json:"token"`
}

// CheckForgotPasswordCodeResponse is the response after validating a code.
type CheckForgotPasswordCodeResponse struct {
	IsValid bool   `json:"is_valid"`
	Email   string `json:"email"`
	Token   string `json:"token"` // This is a new token for the reset step
}

// ResetPasswordRequest is the request to reset the password.
type ResetPasswordRequest struct {
	Token           string `json:"token"`
	NewPassword     string `json:"new_password"`
	PasswordConfirm string `json:"password_confirm"`
}

// ResetPasswordResponse is the response after resetting the password.
type ResetPasswordResponse struct {
	Result string `json:"result"`
}

// ForgotPassword sends a request to initiate the password reset process.
func (c *client.Client) ForgotPassword(ctx context.Context, req *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	var result ForgotPasswordResponse
	err := c.sendRequest(ctx, "POST", "/api/forgot-password", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckForgotPasswordCode validates the code sent to the user's email.
func (c *client.Client) CheckForgotPasswordCode(ctx context.Context, req *CheckForgotPasswordCodeRequest) (*CheckForgotPasswordCodeResponse, error) {
	var result CheckForgotPasswordCodeResponse
	err := c.sendRequest(ctx, "POST", "/api/forgot-password/validity", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetPassword sets a new password using the token from the code check step.
func (c *client.Client) ResetPassword(ctx context.Context, req *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	var result ResetPasswordResponse
	err := c.sendRequest(ctx, "POST", "/api/forgot-password/resets", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
