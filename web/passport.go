package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// PassportResponse is the response from the passport endpoint.
type PassportResponse struct {
	AccessToken string `json:"access_token"`
}

// GetPassport retrieves a web API passport (access token) for a given app.
func (c *client.Client) GetPassport(ctx context.Context, appCode, userID string) (*PassportResponse, error) {
	var result PassportResponse
	path := "/api/passport"
	if userID != "" {
		path += fmt.Sprintf("?user_id=%s", userID)
	}

	headers := map[string]string{
		"X-App-Code": appCode,
	}

	err := c.sendRequest(ctx, "GET", path, nil, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
