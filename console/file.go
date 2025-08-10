package console

import (
	"context"
	"fmt"
	"io"

	"github.com/kingfs/godify/client"
)

// UploadConfig represents the file upload configuration.
type UploadConfig struct {
	FileSizeLimit          int `json:"file_size_limit"`
	BatchCountLimit        int `json:"batch_count_limit"`
	ImageFileSizeLimit     int `json:"image_file_size_limit"`
	VideoFileSizeLimit     int `json:"video_file_size_limit"`
	AudioFileSizeLimit     int `json:"audio_file_size_limit"`
	WorkflowFileUploadLimit int `json:"workflow_file_upload_limit"`
}

// File represents an uploaded file's details.
type File struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Extension  string `json:"extension"`
	MimeType   string `json:"mime_type"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  int64  `json:"created_at"`
	PreviewURL string `json:"preview_url"`
}

// FilePreview represents a text preview of a file.
type FilePreview struct {
	Content string `json:"content"`
}

// SupportedFileTypes represents the list of allowed file extensions.
type SupportedFileTypes struct {
	AllowedExtensions []string `json:"allowed_extensions"`
}

// GetUploadConfig retrieves the file upload configuration.
func (c *client.Client) GetUploadConfig(ctx context.Context) (*UploadConfig, error) {
	var result UploadConfig
	err := c.sendRequest(ctx, "GET", "/console/api/files/upload", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleUploadFile uploads a file from the console.
func (c *client.Client) ConsoleUploadFile(ctx context.Context, source string, file io.Reader, filename string) (*File, error) {
	var result File
	fields := map[string]string{}
	if source != "" {
		fields["source"] = source
	}
	err := c.SendMultipartRequest(ctx, "/console/api/files/upload", file, filename, fields, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFilePreview retrieves a text preview of a file.
func (c *client.Client) GetFilePreview(ctx context.Context, fileID string) (*FilePreview, error) {
	var result FilePreview
	path := fmt.Sprintf("/console/api/files/%s/preview", fileID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSupportedFileTypes retrieves the list of allowed file extensions for upload.
func (c *client.Client) GetSupportedFileTypes(ctx context.Context) (*SupportedFileTypes, error) {
	var result SupportedFileTypes
	err := c.sendRequest(ctx, "GET", "/console/api/files/support-type", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
