package models

// DocumentForAPI API文档
type DocumentForAPI struct {
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	DataSourceType        string    `json:"data_source_type"`
	DocForm               string    `json:"doc_form"`
	DocLanguage           string    `json:"doc_language"`
	IndexingTechnique     string    `json:"indexing_technique"`
	CreatedBy             string    `json:"created_by"`
	CreatedAt             UnixTime  `json:"created_at"`
	UpdatedAt             UnixTime  `json:"updated_at"`
	SegmentCount          int       `json:"segment_count"`
	WordCount             int       `json:"word_count"`
	Status                string    `json:"status"`
	ErrorMessage          string    `json:"error_message,omitempty"`
	HitCount              int       `json:"hit_count"`
	DisplayStatus         string    `json:"display_status"`
	IndexingStatus        string    `json:"indexing_status"`
	CompletedAt           *UnixTime `json:"completed_at,omitempty"`
	ErrorCount            int       `json:"error_count"`
	CompletedSegmentCount int       `json:"completed_segment_count"`
	IsEmpty               bool      `json:"is_empty"`
	IsPaused              bool      `json:"is_paused"`
	IsArchived            bool      `json:"is_archived"`
	IsDeleted             bool      `json:"is_deleted"`
}

// CreateDocumentByTextRequest 通过文本创建文档请求
type CreateDocumentByTextRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	Content           string `json:"content"`
	DocForm           string `json:"doc_form,omitempty"`
	DocLanguage       string `json:"doc_language,omitempty"`
	IndexingTechnique string `json:"indexing_technique,omitempty"`
}

// CreateDocumentByFileRequest 通过文件创建文档请求
type CreateDocumentByFileRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	DocForm           string `json:"doc_form,omitempty"`
	DocLanguage       string `json:"doc_language,omitempty"`
	IndexingTechnique string `json:"indexing_technique,omitempty"`
}

// UpdateDocumentByTextRequest 通过文本更新文档请求
type UpdateDocumentByTextRequest struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Content           string `json:"content"`
	DocForm           string `json:"doc_form,omitempty"`
	DocLanguage       string `json:"doc_language,omitempty"`
	IndexingTechnique string `json:"indexing_technique,omitempty"`
}

// UpdateDocumentByFileRequest 通过文件更新文档请求
type UpdateDocumentByFileRequest struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	DocForm           string `json:"doc_form,omitempty"`
	DocLanguage       string `json:"doc_language,omitempty"`
	IndexingTechnique string `json:"indexing_technique,omitempty"`
}

// SegmentForAPI API分段
type SegmentForAPI struct {
	ID             string    `json:"id"`
	DocumentID     string    `json:"document_id"`
	DocumentName   string    `json:"document_name"`
	DatasetID      string    `json:"dataset_id"`
	DatasetName    string    `json:"dataset_name"`
	Content        string    `json:"content"`
	Answer         string    `json:"answer"`
	WordCount      int       `json:"word_count"`
	CharCount      int       `json:"char_count"`
	Position       int       `json:"position"`
	IsEmpty        bool      `json:"is_empty"`
	HitCount       int       `json:"hit_count"`
	IndexingStatus string    `json:"indexing_status"`
	CompletedAt    *UnixTime `json:"completed_at,omitempty"`
	ErrorMessage   string    `json:"error_message,omitempty"`
	CreatedAt      UnixTime  `json:"created_at"`
	UpdatedAt      UnixTime  `json:"updated_at"`
}

// CreateSegmentsRequest 创建分段请求
type CreateSegmentsRequest struct {
	Segments []SegmentData `json:"segments"`
}

// SegmentData 分段数据
type SegmentData struct {
	Content string `json:"content"`
	Answer  string `json:"answer,omitempty"`
}

// UpdateSegmentRequest 更新分段请求
type UpdateSegmentRequest struct {
	Content string `json:"content"`
	Answer  string `json:"answer,omitempty"`
}

// MetadataForAPI API元数据
type MetadataForAPI struct {
	ID        string   `json:"id"`
	DatasetID string   `json:"dataset_id"`
	Key       string   `json:"key"`
	Value     string   `json:"value"`
	CreatedAt UnixTime `json:"created_at"`
	UpdatedAt UnixTime `json:"updated_at"`
}

// MetadataListResponse 元数据列表响应
type MetadataListResponse struct {
	Data    []MetadataForAPI `json:"data"`
	HasMore bool             `json:"has_more"`
	Limit   int              `json:"limit"`
	Total   int              `json:"total"`
	Page    int              `json:"page"`
}

// CreateMetadataRequest 创建元数据请求
type CreateMetadataRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UpdateMetadataRequest 更新元数据请求
type UpdateMetadataRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HitTestingResult 命中测试结果
type HitTestingResult struct {
	SegmentID   string  `json:"segment_id"`
	DocumentID  string  `json:"document_id"`
	Content     string  `json:"content"`
	Answer      string  `json:"answer"`
	Score       float64 `json:"score"`
	DatasetID   string  `json:"dataset_id"`
	DatasetName string  `json:"dataset_name"`
}
