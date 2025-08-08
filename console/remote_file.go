package console

import (
	"context"
	"fmt"
	"net/url"

	"github.com/kingfs/godify/client"
)

// RemoteFileInfo represents information about a remote file.
type RemoteFileInfo struct {
	FileType   string `json:"file_type"`
	FileLength int    `json:"file_length"`
}

// RemoteFileUploadRequest is the request to upload a file from a URL.
type RemoteFileUploadRequest struct {
	URL string `json:"url"`
}

// RemoteFileUploadResponse is the response after uploading a remote file.
type RemoteFileUploadResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	Extension string `json:"extension"`
	URL       string `json:"url"`
	MimeType  string `json:"mime_type"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

// GetRemoteFileInfo retrieves information about a remote file.
func (c *client.Client) ConsoleGetRemoteFileInfo(ctx context.Context, remoteURL string) (*RemoteFileInfo, error) {
	var result RemoteFileInfo
	encodedURL := url.PathEscape(remoteURL)
	path := fmt.Sprintf("/console/api/remote-files/%s", encodedURL)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleUploadRemoteFile uploads a file from a given URL.
func (c *client.Client) ConsoleUploadRemoteFile(ctx context.Context, req *RemoteFileUploadRequest) (*RemoteFileUploadResponse, error) {
	var result RemoteFileUploadResponse
	err := c.sendRequest(ctx, "POST", "/console/api/remote-files/upload", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
