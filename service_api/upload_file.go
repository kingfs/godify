package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// DocumentUploadFile represents the details of a file uploaded for a document.
type DocumentUploadFile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	Extension   string `json:"extension"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
	MimeType    string `json:"mime_type"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`
}

// GetDocumentUploadFile retrieves the upload file details for a specific document.
func (c *client.Client) GetDocumentUploadFile(ctx context.Context, datasetID, documentID string) (*DocumentUploadFile, error) {
	var result DocumentUploadFile
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/upload-file", datasetID, documentID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
