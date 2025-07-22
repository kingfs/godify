package models

// Workflow 工作流
type Workflow struct {
	ID          string                 `json:"id"`
	AppID       string                 `json:"app_id"`
	Type        string                 `json:"type"`
	Version     string                 `json:"version"`
	Graph       map[string]interface{} `json:"graph"`
	Features    map[string]interface{} `json:"features"`
	CreatedBy   string                 `json:"created_by"`
	CreatedAt   UnixTime               `json:"created_at"`
	UpdatedBy   string                 `json:"updated_by"`
	UpdatedAt   UnixTime               `json:"updated_at"`
	Environment string                 `json:"environment"`
}

// WorkflowRunRequest 工作流运行请求
type WorkflowRunRequest struct {
	Inputs map[string]interface{}   `json:"inputs"`
	Files  []map[string]interface{} `json:"files,omitempty"`
}

// WorkflowRunResponse 工作流运行响应
type WorkflowRunResponse struct {
	WorkflowRunID string                 `json:"workflow_run_id"`
	TaskID        string                 `json:"task_id"`
	Data          map[string]interface{} `json:"data"`
	Event         string                 `json:"event"`
	CreatedAt     UnixTime               `json:"created_at"`
}

// WorkflowStopResponse 工作流停止响应
type WorkflowStopResponse struct {
	Result string `json:"result"`
}

// WorkflowNodeExecution 工作流节点执行
type WorkflowNodeExecution struct {
	ID                string                 `json:"id"`
	Index             int                    `json:"index"`
	NodeID            string                 `json:"node_id"`
	NodeType          string                 `json:"node_type"`
	Title             string                 `json:"title"`
	Inputs            map[string]interface{} `json:"inputs"`
	ProcessData       map[string]interface{} `json:"process_data"`
	Outputs           map[string]interface{} `json:"outputs"`
	Status            string                 `json:"status"`
	Error             string                 `json:"error,omitempty"`
	ElapsedTime       float64                `json:"elapsed_time"`
	ExecutionMetadata map[string]interface{} `json:"execution_metadata"`
	CreatedAt         UnixTime               `json:"created_at"`
	FinishedAt        *UnixTime              `json:"finished_at"`
}

// WorkflowRun 工作流运行记录
type WorkflowRun struct {
	ID             string                  `json:"id"`
	WorkflowID     string                  `json:"workflow_id"`
	TriggerUser    string                  `json:"trigger_user"`
	Status         string                  `json:"status"`
	Inputs         map[string]interface{}  `json:"inputs"`
	Outputs        map[string]interface{}  `json:"outputs"`
	Error          string                  `json:"error,omitempty"`
	ElapsedTime    float64                 `json:"elapsed_time"`
	TotalTokens    int                     `json:"total_tokens"`
	TotalSteps     int                     `json:"total_steps"`
	CreatedAt      UnixTime                `json:"created_at"`
	FinishedAt     *UnixTime               `json:"finished_at"`
	NodeExecutions []WorkflowNodeExecution `json:"node_executions"`
}

// WorkflowRunListResponse 工作流运行列表响应
type WorkflowRunListResponse struct {
	Data    []WorkflowRun `json:"data"`
	HasMore bool          `json:"has_more"`
	Limit   int           `json:"limit"`
	Total   int           `json:"total"`
	Page    int           `json:"page"`
}

// DraftWorkflowRequest 草稿工作流请求
type DraftWorkflowRequest struct {
	Graph    map[string]interface{} `json:"graph"`
	Features map[string]interface{} `json:"features"`
	Hash     string                 `json:"hash,omitempty"`
}

// WorkflowPublishRequest 工作流发布请求
type WorkflowPublishRequest struct {
	Description string `json:"description,omitempty"`
}

// WorkflowImportRequest 工作流导入请求
type WorkflowImportRequest struct {
	Data string `json:"data"`
	Name string `json:"name,omitempty"`
}

// DatasetForAPI Service API数据集结构
type DatasetForAPI struct {
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description"`
	Permission             string                 `json:"permission"`
	DataSourceType         string                 `json:"data_source_type"`
	IndexingTechnique      string                 `json:"indexing_technique"`
	AppCount               int                    `json:"app_count"`
	DocumentCount          int                    `json:"document_count"`
	WordCount              int                    `json:"word_count"`
	CreatedBy              string                 `json:"created_by"`
	CreatedAt              UnixTime               `json:"created_at"`
	UpdatedAt              UnixTime               `json:"updated_at"`
	EmbeddingModel         string                 `json:"embedding_model,omitempty"`
	EmbeddingModelProvider string                 `json:"embedding_model_provider,omitempty"`
	EmbeddingAvailable     bool                   `json:"embedding_available"`
	RetrievalModelDict     map[string]interface{} `json:"retrieval_model_dict,omitempty"`
	Tags                   []string               `json:"tags"`
}

// DatasetListForAPIResponse Service API数据集列表响应
type DatasetListForAPIResponse struct {
	Data    []DatasetForAPI `json:"data"`
	HasMore bool            `json:"has_more"`
	Limit   int             `json:"limit"`
	Total   int             `json:"total"`
	Page    int             `json:"page"`
}

// CreateDatasetForAPIRequest Service API创建数据集请求
type CreateDatasetForAPIRequest struct {
	Name                   string                 `json:"name"`
	Description            string                 `json:"description,omitempty"`
	IndexingTechnique      string                 `json:"indexing_technique,omitempty"`
	Permission             string                 `json:"permission,omitempty"`
	ExternalKnowledgeAPI   map[string]interface{} `json:"external_knowledge_api,omitempty"`
	ExternalKnowledgeID    string                 `json:"external_knowledge_id,omitempty"`
	ExternalRetrievalModel map[string]interface{} `json:"external_retrieval_model,omitempty"`
}

// DatasetDocument 数据集文档
type DatasetDocument struct {
	ID                 string                 `json:"id"`
	Position           int                    `json:"position"`
	DataSourceType     string                 `json:"data_source_type"`
	DataSourceInfo     map[string]interface{} `json:"data_source_info"`
	DatasetProcessRule map[string]interface{} `json:"dataset_process_rule"`
	Name               string                 `json:"name"`
	CreatedFrom        string                 `json:"created_from"`
	CreatedBy          string                 `json:"created_by"`
	CreatedAt          UnixTime               `json:"created_at"`
	Tokens             int                    `json:"tokens"`
	IndexingStatus     string                 `json:"indexing_status"`
	Error              string                 `json:"error,omitempty"`
	Enabled            bool                   `json:"enabled"`
	DisabledAt         *UnixTime              `json:"disabled_at"`
	DisabledBy         string                 `json:"disabled_by,omitempty"`
	Archived           bool                   `json:"archived"`
	DisplayStatus      string                 `json:"display_status"`
	WordCount          int                    `json:"word_count"`
	HitCount           int                    `json:"hit_count"`
}

// DocumentListResponse 文档列表响应
type DocumentListResponse struct {
	Data    []DatasetDocument `json:"data"`
	HasMore bool              `json:"has_more"`
	Limit   int               `json:"limit"`
	Total   int               `json:"total"`
	Page    int               `json:"page"`
}

// DocumentSegment 文档片段
type DocumentSegment struct {
	ID          string    `json:"id"`
	Position    int       `json:"position"`
	DocumentID  string    `json:"document_id"`
	Content     string    `json:"content"`
	Answer      string    `json:"answer,omitempty"`
	WordCount   int       `json:"word_count"`
	Tokens      int       `json:"tokens"`
	Keywords    []string  `json:"keywords"`
	Index       int       `json:"index"`
	HitCount    int       `json:"hit_count"`
	Enabled     bool      `json:"enabled"`
	DisabledAt  *UnixTime `json:"disabled_at"`
	DisabledBy  string    `json:"disabled_by,omitempty"`
	Status      string    `json:"status"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   UnixTime  `json:"created_at"`
	IndexingAt  *UnixTime `json:"indexing_at"`
	CompletedAt *UnixTime `json:"completed_at"`
	Error       string    `json:"error,omitempty"`
	StoppedAt   *UnixTime `json:"stopped_at"`
}

// SegmentListResponse 片段列表响应
type SegmentListResponse struct {
	Data    []DocumentSegment `json:"data"`
	HasMore bool              `json:"has_more"`
	Limit   int               `json:"limit"`
	Total   int               `json:"total"`
	Page    int               `json:"page"`
}

// HitTestingRequest 命中测试请求
type HitTestingRequest struct {
	Query                  string                 `json:"query"`
	RetrievalModel         map[string]interface{} `json:"retrieval_model,omitempty"`
	ExternalRetrievalModel map[string]interface{} `json:"external_retrieval_model,omitempty"`
}

// HitTestingResponse 命中测试响应
type HitTestingResponse struct {
	Query   string          `json:"query"`
	Records []HitTestRecord `json:"records"`
}

// HitTestRecord 命中测试记录
type HitTestRecord struct {
	Index      int       `json:"index"`
	Position   int       `json:"position"`
	Source     string    `json:"source"`
	Score      float64   `json:"score"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	TSNEVector []float64 `json:"tsne_position,omitempty"`
}

// Tag 标签
type Tag struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	CreatedBy string   `json:"created_by"`
	CreatedAt UnixTime `json:"created_at"`
}

// TagListResponse 标签列表响应
type TagListResponse struct {
	Data []Tag `json:"data"`
}
