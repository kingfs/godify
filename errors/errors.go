package errors

import "fmt"

// ErrorResponse API错误响应结构
type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// APIError API错误
type APIError struct {
	StatusCode int         `json:"status_code"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("API error [%d:%s]: %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("API error [%d]: %s", e.StatusCode, e.Message)
}

// 预定义的错误类型
var (
	ErrAppUnavailable                           = &APIError{Code: "app_unavailable", Message: "App unavailable"}
	ErrNotChatApp                               = &APIError{Code: "not_chat_app", Message: "App is not a chat app"}
	ErrNotCompletionApp                         = &APIError{Code: "not_completion_app", Message: "App is not a completion app"}
	ErrConversationNotExists                    = &APIError{Code: "conversation_not_exists", Message: "Conversation not exists"}
	ErrConversationCompleted                    = &APIError{Code: "conversation_completed", Message: "Conversation completed"}
	ErrMessageNotExists                         = &APIError{Code: "message_not_exists", Message: "Message not exists"}
	ErrProviderNotInitialize                    = &APIError{Code: "provider_not_initialize", Message: "Provider not initialize"}
	ErrProviderQuotaExceeded                    = &APIError{Code: "provider_quota_exceeded", Message: "Provider quota exceeded"}
	ErrProviderModelCurrentlyNotSupport         = &APIError{Code: "provider_model_currently_not_support", Message: "Provider model currently not support"}
	ErrCompletionRequest                        = &APIError{Code: "completion_request_error", Message: "Completion request error"}
	ErrInvokeRateLimit                          = &APIError{Code: "invoke_rate_limit", Message: "Invoke rate limit"}
	ErrNoFileUploaded                           = &APIError{Code: "no_file_uploaded", Message: "No file uploaded"}
	ErrTooManyFiles                             = &APIError{Code: "too_many_files", Message: "Too many files"}
	ErrFileTooLarge                             = &APIError{Code: "file_too_large", Message: "File too large"}
	ErrUnsupportedFileType                      = &APIError{Code: "unsupported_file_type", Message: "Unsupported file type"}
	ErrAudioTooLarge                            = &APIError{Code: "audio_too_large", Message: "Audio too large"}
	ErrUnsupportedAudioType                     = &APIError{Code: "unsupported_audio_type", Message: "Unsupported audio type"}
	ErrProviderNotSupportSpeechToText           = &APIError{Code: "provider_not_support_speech_to_text", Message: "Provider not support speech to text"}
	ErrAppMoreLikeThisDisabled                  = &APIError{Code: "app_more_like_this_disabled", Message: "App more like this disabled"}
	ErrAppSuggestedQuestionsAfterAnswerDisabled = &APIError{Code: "app_suggested_questions_after_answer_disabled", Message: "App suggested questions after answer disabled"}
)

// IsAPIError 检查是否为API错误
func IsAPIError(err error) bool {
	_, ok := err.(*APIError)
	return ok
}

// GetAPIError 获取API错误详情
func GetAPIError(err error) *APIError {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr
	}
	return nil
}
