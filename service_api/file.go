package service_api

import (
	"context"
	"io"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// UploadFile uploads a file for use in an application.
// The file should be an io.Reader, and filename should be provided.
// The user parameter is the ID of the end-user.
func (c *client.Client) UploadFile(ctx context.Context, user string, file io.Reader, filename string) (*types.FileUploadResponse, error) {
	var result types.FileUploadResponse
	fields := map[string]string{
		"user": user,
	}

	err := c.SendMultipartRequest(ctx, "/v1/files/upload", file, filename, fields, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
