package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// VersionFeatures represents the features section of the version response.
type VersionFeatures struct {
	CanReplaceLogo            bool `json:"can_replace_logo"`
	ModelLoadBalancingEnabled bool `json:"model_load_balancing_enabled"`
}

// VersionResponse is the response from the version endpoint.
type VersionResponse struct {
	Version       string          `json:"version"`
	ReleaseDate   string          `json:"release_date"`
	ReleaseNotes  string          `json:"release_notes"`
	CanAutoUpdate bool            `json:"can_auto_update"`
	Features      VersionFeatures `json:"features"`
}

// GetVersion retrieves the version information of the Dify instance.
func (c *client.Client) GetVersion(ctx context.Context, currentVersion string) (*VersionResponse, error) {
	var result VersionResponse
	path := fmt.Sprintf("/console/api/version?current_version=%s", currentVersion)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
