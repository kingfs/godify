package models

import (
	"encoding/json"
	"time"
)

// UnixTime 自定义时间类型，用于处理Dify API返回的Unix时间戳
type UnixTime struct {
	time.Time
}

// UnmarshalJSON 从Unix时间戳反序列化
func (ut *UnixTime) UnmarshalJSON(data []byte) error {
	// 处理null值
	if string(data) == "null" {
		return nil
	}

	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		// 如果不是数字，尝试解析为字符串格式的时间
		var timeStr string
		if err2 := json.Unmarshal(data, &timeStr); err2 != nil {
			return err // 返回原始错误
		}

		// 尝试解析RFC3339格式
		t, err3 := time.Parse(time.RFC3339, timeStr)
		if err3 != nil {
			return err3
		}
		ut.Time = t
		return nil
	}

	ut.Time = time.Unix(timestamp, 0)
	return nil
}

// MarshalJSON 序列化为Unix时间戳
func (ut UnixTime) MarshalJSON() ([]byte, error) {
	if ut.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ut.Unix())
}

// PaginationResponse 分页响应基础结构
type PaginationResponse struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	HasMore bool `json:"has_more"`
}

// InfiniteScrollPagination 无限滚动分页
type InfiniteScrollPagination struct {
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
}

// FileUpload 文件上传结构
type FileUpload struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Size      int64    `json:"size"`
	Extension string   `json:"extension"`
	MimeType  string   `json:"mime_type"`
	CreatedBy string   `json:"created_by"`
	CreatedAt UnixTime `json:"created_at"`
	URL       string   `json:"url,omitempty"`
}

// AppMode 应用模式
type AppMode string

const (
	AppModeCompletion   AppMode = "completion"
	AppModeChat         AppMode = "chat"
	AppModeAgentChat    AppMode = "agent-chat"
	AppModeAdvancedChat AppMode = "advanced-chat"
	AppModeWorkflow     AppMode = "workflow"
)

// ResponseMode 响应模式
type ResponseMode string

const (
	ResponseModeBlocking  ResponseMode = "blocking"
	ResponseModeStreaming ResponseMode = "streaming"
)

// AppInfo 应用基本信息
type AppInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Mode        AppMode  `json:"mode"`
	AuthorName  string   `json:"author_name"`
}

// AppMeta 应用元数据
type AppMeta struct {
	ToolIcons map[string]string `json:"tool_icons"`
}

// AppParameters 应用参数
type AppParameters struct {
	OpeningStatement              string                 `json:"opening_statement,omitempty"`
	SuggestedQuestions            []interface{}          `json:"suggested_questions,omitempty"`
	SuggestedQuestionsAfterAnswer map[string]interface{} `json:"suggested_questions_after_answer,omitempty"`
	SpeechToText                  map[string]interface{} `json:"speech_to_text,omitempty"`
	TextToSpeech                  map[string]interface{} `json:"text_to_speech,omitempty"`
	RetrieverResource             map[string]interface{} `json:"retriever_resource,omitempty"`
	AnnotationReply               map[string]interface{} `json:"annotation_reply,omitempty"`
	MoreLikeThis                  map[string]interface{} `json:"more_like_this,omitempty"`
	UserInputForm                 []interface{}          `json:"user_input_form"`
	SensitiveWordAvoidance        map[string]interface{} `json:"sensitive_word_avoidance,omitempty"`
	FileUpload                    map[string]interface{} `json:"file_upload,omitempty"`
	SystemParameters              *SystemParameters      `json:"system_parameters,omitempty"`
}

// SystemParameters 系统参数
type SystemParameters struct {
	ImageFileSizeLimit      int `json:"image_file_size_limit,omitempty"`
	VideoFileSizeLimit      int `json:"video_file_size_limit,omitempty"`
	AudioFileSizeLimit      int `json:"audio_file_size_limit,omitempty"`
	FileSizeLimit           int `json:"file_size_limit,omitempty"`
	WorkflowFileUploadLimit int `json:"workflow_file_upload_limit,omitempty"`
}

// UserInputFormItem 用户输入表单项 (原始格式，保持灵活性)
type UserInputFormItem map[string]interface{}

// ParagraphConfig 段落配置
type ParagraphConfig struct {
	Label       string `json:"label"`
	Variable    string `json:"variable"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Description string `json:"description,omitempty"`
}

// SelectConfig 选择配置
type SelectConfig struct {
	Label       string               `json:"label"`
	Variable    string               `json:"variable"`
	Required    bool                 `json:"required"`
	Default     string               `json:"default"`
	Description string               `json:"description,omitempty"`
	Options     []SelectOptionConfig `json:"options"`
}

// SelectOptionConfig 选择选项配置
type SelectOptionConfig struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// TextInputConfig 文本输入配置
type TextInputConfig struct {
	Label       string `json:"label"`
	Variable    string `json:"variable"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Description string `json:"description,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
}

// NumberConfig 数字配置
type NumberConfig struct {
	Label       string  `json:"label"`
	Variable    string  `json:"variable"`
	Required    bool    `json:"required"`
	Default     float64 `json:"default"`
	Description string  `json:"description,omitempty"`
	Min         float64 `json:"min,omitempty"`
	Max         float64 `json:"max,omitempty"`
}

// FileUploadConfig 文件上传配置 (保持灵活性)
type FileUploadConfig map[string]interface{}

// FileTypeConfig 文件类型配置
type FileTypeConfig struct {
	Enabled    bool     `json:"enabled"`
	Number     int      `json:"number_limits,omitempty"`
	Extensions []string `json:"extensions,omitempty"`
	FileSize   int      `json:"file_size_limit,omitempty"`
}

// AccessMode 访问模式
type AccessMode string

const (
	AccessModePublic  AccessMode = "public"
	AccessModePrivate AccessMode = "private"
)

// WebAppAccessMode Web应用访问模式响应
type WebAppAccessMode struct {
	AccessMode AccessMode `json:"accessMode"`
}

// WebAppPermission Web应用权限响应
type WebAppPermission struct {
	Result bool `json:"result"`
}
