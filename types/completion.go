package types

// File represents a file to be sent with a request.
type File struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url"`
}

// CompletionRequest is the request for creating a completion message.
type CompletionRequest struct {
	Inputs        map[string]interface{} `json:"inputs"`
	Query         string                 `json:"query,omitempty"`
	Files         []File                 `json:"files,omitempty"`
	ResponseMode  string                 `json:"response_mode"`
	RetrieverFrom string                 `json:"retriever_from,omitempty"`
	User          string                 `json:"user"`
}

// ChatRequest is the request for creating a chat message.
type ChatRequest struct {
	Inputs           map[string]interface{} `json:"inputs"`
	Query            string                 `json:"query"`
	Files            []File                 `json:"files,omitempty"`
	ResponseMode     string                 `json:"response_mode"`
	ConversationID   string                 `json:"conversation_id,omitempty"`
	RetrieverFrom    string                 `json:"retriever_from,omitempty"`
	AutoGenerateName bool                   `json:"auto_generate_name"`
	User             string                 `json:"user"`
}

// BlockingResponse is the response for a blocking chat or completion call.
type BlockingResponse struct {
	Event          string                 `json:"event"`
	MessageID      string                 `json:"message_id"`
	ConversationID string                 `json:"conversation_id"`
	Mode           string                 `json:"mode"`
	Answer         string                 `json:"answer"`
	Metadata       map[string]interface{} `json:"metadata"`
	CreatedAt      int64                  `json:"created_at"`
}

// StreamEvent represents a single event from a streaming response.
type StreamEvent struct {
	Event          string                 `json:"event"`
	MessageID      string                 `json:"message_id,omitempty"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	Mode           string                 `json:"mode,omitempty"`
	Answer         string                 `json:"answer,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      int64                  `json:"created_at,omitempty"`
	Error          string                 `json:"error,omitempty"`
	ErrorCode      int                    `json:"error_code,omitempty"`
}

// StopResponse is the response for a stop request.
type StopResponse struct {
	Result string `json:"result"`
}
