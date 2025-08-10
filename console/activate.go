package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ActivateCheckResponse is the response from checking an activation token.
type ActivateCheckResponse struct {
	IsValid bool `json:"is_valid"`
	Data    struct {
		WorkspaceName string `json:"workspace_name"`
		WorkspaceID   string `json:"workspace_id"`
		Email         string `json:"email"`
	} `json:"data"`
}

// ActivateRequest is the request to activate an account.
type ActivateRequest struct {
	WorkspaceID       string `json:"workspace_id,omitempty"`
	Email             string `json:"email,omitempty"`
	Token             string `json:"token"`
	Name              string `json:"name"`
	InterfaceLanguage string `json:"interface_language"`
	Timezone          string `json:"timezone"`
}

// ActivateResponse is the response after activating an account.
type ActivateResponse struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

// CheckActivationToken checks if an activation token is valid.
func (c *client.Client) CheckActivationToken(ctx context.Context, token, workspaceID, email string) (*ActivateCheckResponse, error) {
	var result ActivateCheckResponse
	path := fmt.Sprintf("/console/api/activate/check?token=%s&workspace_id=%s&email=%s", token, workspaceID, email)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ActivateAccount activates a user account.
func (c *client.Client) ActivateAccount(ctx context.Context, req *ActivateRequest) (*ActivateResponse, error) {
	var result ActivateResponse
	err := c.sendRequest(ctx, "POST", "/console/api/activate", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
