package models

// Message 消息结构
type Message struct {
	ID                 string                 `json:"id"`
	ConversationID     string                 `json:"conversation_id"`
	ParentMessageID    string                 `json:"parent_message_id,omitempty"`
	Inputs             map[string]interface{} `json:"inputs"`
	Query              string                 `json:"query"`
	Answer             string                 `json:"answer"`
	MessageFiles       []MessageFile          `json:"message_files"`
	Feedback           *MessageFeedback       `json:"feedback"`
	RetrieverResources []RetrieverResource    `json:"retriever_resources"`
	CreatedAt          UnixTime               `json:"created_at"`
	AgentThoughts      []AgentThought         `json:"agent_thoughts"`
	Metadata           map[string]interface{} `json:"metadata"`
	Status             string                 `json:"status"`
	Error              string                 `json:"error,omitempty"`
}

// MessageFile 消息文件
type MessageFile struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	BelongsTo string `json:"belongs_to"`
}

// MessageFeedback 消息反馈
type MessageFeedback struct {
	Rating  string `json:"rating"`
	Content string `json:"content,omitempty"`
}

// RetrieverResource 检索资源
type RetrieverResource struct {
	Position       int     `json:"position"`
	DatasetID      string  `json:"dataset_id"`
	DatasetName    string  `json:"dataset_name"`
	DocumentID     string  `json:"document_id"`
	DocumentName   string  `json:"document_name"`
	DataSourceType string  `json:"data_source_type"`
	SegmentID      string  `json:"segment_id"`
	Score          float64 `json:"score"`
	Content        string  `json:"content"`
}

// AgentThought Agent思考过程
type AgentThought struct {
	ID              string                 `json:"id"`
	MessageID       string                 `json:"message_id"`
	Position        int                    `json:"position"`
	Thought         string                 `json:"thought"`
	Tool            string                 `json:"tool"`
	ToolInput       map[string]interface{} `json:"tool_input"`
	ToolOutput      string                 `json:"tool_output"`
	CreatedAt       UnixTime               `json:"created_at"`
	MessageFiles    []MessageFile          `json:"message_files"`
	ToolProcessData map[string]interface{} `json:"tool_process_data"`
}

// MessageListResponse 消息列表响应
type MessageListResponse struct {
	InfiniteScrollPagination
	Data []Message `json:"data"`
}

// MessageFeedbackRequest 消息反馈请求
type MessageFeedbackRequest struct {
	Rating  *string `json:"rating,omitempty"` // "like", "dislike"
	Content *string `json:"content,omitempty"`
}

// SuggestedQuestionsResponse 建议问题响应
type SuggestedQuestionsResponse struct {
	Data []string `json:"data"`
}

// GenerateResponse 生成响应（用于completion和chat）
type GenerateResponse struct {
	Event          string                 `json:"event,omitempty"`
	MessageID      string                 `json:"message_id,omitempty"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	Mode           string                 `json:"mode,omitempty"`
	Answer         string                 `json:"answer,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      int64                  `json:"created_at,omitempty"`
	TaskID         string                 `json:"task_id,omitempty"`
	ID             string                 `json:"id,omitempty"`

	// 流式响应字段
	Delta string `json:"delta,omitempty"`

	// 错误字段
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
