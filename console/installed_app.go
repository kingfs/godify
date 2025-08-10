package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledApp represents an app installed from the explore section.
type InstalledApp struct {
	ID                string      `json:"id"`
	App               App         `json:"app"`
	AppOwnerTenantID  string      `json:"app_owner_tenant_id"`
	IsPinned          bool        `json:"is_pinned"`
	LastUsedAt        int64       `json:"last_used_at"`
	Editable          bool        `json:"editable"`
	Uninstallable     bool        `json:"uninstallable"`
}

// InstalledAppListResponse is the response for listing installed apps.
type InstalledAppListResponse struct {
	InstalledApps []InstalledApp `json:"installed_apps"`
}

// InstallAppRequest is the request to install an app.
type InstallAppRequest struct {
	AppID string `json:"app_id"`
}

// UpdateInstalledAppRequest is the request to update an installed app.
type UpdateInstalledAppRequest struct {
	IsPinned *bool `json:"is_pinned,omitempty"`
}

// GetInstalledApps retrieves a list of installed apps.
func (c *client.Client) GetInstalledApps(ctx context.Context, appID string) (*InstalledAppListResponse, error) {
	var result InstalledAppListResponse
	path := "/console/api/installed-apps"
	if appID != "" {
		path += fmt.Sprintf("?app_id=%s", appID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstallApp installs an app from the explore section.
func (c *client.Client) InstallApp(ctx context.Context, appID string) (*types.StopResponse, error) {
	var result types.StopResponse
	req := InstallAppRequest{AppID: appID}
	err := c.sendRequest(ctx, "POST", "/console/api/installed-apps", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateInstalledApp updates the properties of an installed app (e.g., pinning).
func (c *client.Client) UpdateInstalledApp(ctx context.Context, installedAppID string, req *UpdateInstalledAppRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s", installedAppID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UninstallApp uninstalls an app.
func (c *client.Client) UninstallApp(ctx context.Context, installedAppID string) error {
	path := fmt.Sprintf("/console/api/installed-apps/%s", installedAppID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
