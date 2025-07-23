package models

// DocumentMetadata 文档元数据
type DocumentMetadata struct {
	ID         string   `json:"id"`
	DocumentID string   `json:"document_id"`
	DatasetID  string   `json:"dataset_id"`
	Key        string   `json:"key"`
	Value      string   `json:"value"`
	CreatedBy  string   `json:"created_by"`
	CreatedAt  UnixTime `json:"created_at"`
	UpdatedBy  *string  `json:"updated_by"`
	UpdatedAt  UnixTime `json:"updated_at"`
}

// Document 文档结构
type Document struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	DatasetID      string   `json:"dataset_id"`
	DocumentType   string   `json:"document_type"`
	DataSourceType string   `json:"data_source_type"`
	CreatedBy      string   `json:"created_by"`
	CreatedAt      UnixTime `json:"created_at"`
	UpdatedBy      *string  `json:"updated_by"`
	UpdatedAt      UnixTime `json:"updated_at"`
}
