package types

// AppParametersResponse is the response for the GetAppParameters endpoint.
type AppParametersResponse struct {
	OpeningStatement             string                   `json:"opening_statement"`
	SuggestedQuestions           []string                 `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer map[string]interface{} `json:"suggested_questions_after_answer"`
	SpeechToText                 map[string]interface{}   `json:"speech_to_text"`
	TextToSpeech                 map[string]interface{}   `json:"text_to_speech"`
	RetrieverResource            map[string]interface{}   `json:"retriever_resource"`
	AnnotationReply              map[string]interface{}   `json:"annotation_reply"`
	MoreLikeThis                 map[string]interface{}   `json:"more_like_this"`
	UserInputForm                []map[string]interface{} `json:"user_input_form"`
	SensitiveWordAvoidance       map[string]interface{}   `json:"sensitive_word_avoidance"`
	FileUpload                   map[string]interface{}   `json:"file_upload"`
	SystemParameters             SystemParameters         `json:"system_parameters"`
}

// SystemParameters defines the system parameters for an app.
type SystemParameters struct {
	ImageFileSizeLimit       int `json:"image_file_size_limit"`
	VideoFileSizeLimit       int `json:"video_file_size_limit"`
	AudioFileSizeLimit       int `json:"audio_file_size_limit"`
	FileSizeLimit            int `json:"file_size_limit"`
	WorkflowFileUploadLimit int `json:"workflow_file_upload_limit"`
}

// AppMetaResponse is the response for the GetAppMeta endpoint.
type AppMetaResponse struct {
	ToolIcons map[string]interface{} `json:"tool_icons"`
}

// AppInfoResponse is the response for the GetAppInfo endpoint.
type AppInfoResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Mode        string   `json:"mode"`
	AuthorName  string   `json:"author_name"`
}

// Site represents the configuration of a web app site.
type Site struct {
	Title                   string `json:"title"`
	ChatColorTheme          string `json:"chat_color_theme"`
	ChatColorThemeInverted  bool   `json:"chat_color_theme_inverted"`
	IconType                string `json:"icon_type"`
	Icon                    string `json:"icon"`
	IconBackground          string `json:"icon_background"`
	IconURL                 string `json:"icon_url"`
	Description             string `json:"description"`
	Copyright               string `json:"copyright"`
	PrivacyPolicy           string `json:"privacy_policy"`
	CustomDisclaimer        string `json:"custom_disclaimer"`
	DefaultLanguage         string `json:"default_language"`
	ShowWorkflowSteps       bool   `json:"show_workflow_steps"`
	UseIconAsAnswerIcon     bool   `json:"use_icon_as_answer_icon"`
}
