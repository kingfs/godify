package models

// Conversation 对话结构
type Conversation struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Inputs       map[string]interface{} `json:"inputs"`
	Introduction string                 `json:"introduction"`
	CreatedAt    UnixTime               `json:"created_at"`
	UpdatedAt    UnixTime               `json:"updated_at"`
}

// ConversationListResponse 对话列表响应
type ConversationListResponse struct {
	InfiniteScrollPagination
	Data []Conversation `json:"data"`
}

// ConversationRenameRequest 对话重命名请求
type ConversationRenameRequest struct {
	Name         *string `json:"name,omitempty"`
	AutoGenerate bool    `json:"auto_generate,omitempty"`
}

// SimpleConversation 简单对话信息
type SimpleConversation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ConversationOperationResponse 对话操作响应
type ConversationOperationResponse struct {
	Result string `json:"result"`
}
