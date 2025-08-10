package types

// FileUploadResponse is the response from the file upload endpoint.
type FileUploadResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Extension  string `json:"extension"`
	MimeType   string `json:"mime_type"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  int64  `json:"created_at"`
	PreviewURL string `json:"preview_url"`
}
