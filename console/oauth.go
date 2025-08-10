package console

import (
	"context"
	"errors"
	"fmt"

	"github.com/kingfs/godify/client"
)

// GetOAuthLoginURL retrieves the OAuth login URL for a given provider.
// It makes a request to the endpoint and expects a redirect. It returns the URL
// from the 'Location' header of the 302 response.
func (c *client.Client) GetOAuthLoginURL(ctx context.Context, provider, inviteToken string) (string, error) {
	path := fmt.Sprintf("/console/api/oauth/login/%s", provider)
	if inviteToken != "" {
		path += fmt.Sprintf("?invite_token=%s", inviteToken)
	}

	err := c.sendRequest(ctx, "GET", path, nil, nil, nil)

	var redirectErr *client.RedirectError
	if errors.As(err, &redirectErr) {
		return redirectErr.Location, nil
	}

	if err == nil {
		return "", fmt.Errorf("expected a redirect, but got a successful response")
	}

	return "", err
}
