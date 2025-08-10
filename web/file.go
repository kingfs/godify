package web

import (
	"context"
	"io"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// UploadFile uploads a file for use in the web application.
// The file should be an io.Reader, and filename should be provided.
// The user parameter is the ID of the end-user.
// The source can be "datasets" or empty.
func (c *client.Client) WebUploadFile(ctx context.Context, user, source, filename string, file io.Reader) (*types.FileUploadResponse, error) {
	var result types.FileUploadResponse
	fields := map[string]string{
		"user": user,
	}
	if source != "" {
		fields["source"] = source
	}

	err := c.SendMultipartRequest(ctx, "/api/files/upload", file, filename, fields, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
