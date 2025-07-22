package models

// ConsoleApp Console应用信息
type ConsoleApp struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Mode           AppMode    `json:"mode"`
	IconType       string     `json:"icon_type,omitempty"`
	Icon           string     `json:"icon,omitempty"`
	IconBackground string     `json:"icon_background,omitempty"`
	EnableSite     bool       `json:"enable_site"`
	EnableAPI      bool       `json:"enable_api"`
	CreatedAt      UnixTime   `json:"created_at"`
	UpdatedAt      UnixTime   `json:"updated_at"`
	TenantID       string     `json:"tenant_id"`
	Site           *AppSite   `json:"site,omitempty"`
	APIBaseURL     string     `json:"api_base_url,omitempty"`
	Tags           []AppTag   `json:"tags"`
	AccessMode     AccessMode `json:"access_mode,omitempty"`
}

// AppSite 应用站点信息
type AppSite struct {
	AccessToken            string `json:"access_token"`
	Code                   string `json:"code"`
	Title                  string `json:"title"`
	Icon                   string `json:"icon"`
	IconBackground         string `json:"icon_background"`
	Description            string `json:"description"`
	DefaultLanguage        string `json:"default_language"`
	ChatColorTheme         string `json:"chat_color_theme"`
	ChatColorThemeInverted bool   `json:"chat_color_theme_inverted"`
	CustomDisclaimer       string `json:"custom_disclaimer"`
	CustomizeTokenStrategy string `json:"customize_token_strategy"`
	PromptPublic           bool   `json:"prompt_public"`
	Copyright              string `json:"copyright"`
	PrivacyPolicy          string `json:"privacy_policy"`
	ShowWorkflowSteps      bool   `json:"show_workflow_steps"`
}

// AppTag 应用标签
type AppTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ConsoleAppListResponse Console应用列表响应
type ConsoleAppListResponse struct {
	Data    []ConsoleApp `json:"data"`
	HasMore bool         `json:"has_more"`
	Limit   int          `json:"limit"`
	Total   int          `json:"total"`
	Page    int          `json:"page"`
}

// CreateAppRequest 创建应用请求
type CreateAppRequest struct {
	Name           string  `json:"name"`
	Description    string  `json:"description,omitempty"`
	Mode           AppMode `json:"mode"`
	IconType       string  `json:"icon_type,omitempty"`
	Icon           string  `json:"icon,omitempty"`
	IconBackground string  `json:"icon_background,omitempty"`
}

// UpdateAppRequest 更新应用请求
type UpdateAppRequest struct {
	Name                string `json:"name"`
	Description         string `json:"description,omitempty"`
	IconType            string `json:"icon_type,omitempty"`
	Icon                string `json:"icon,omitempty"`
	IconBackground      string `json:"icon_background,omitempty"`
	UseIconAsAnswerIcon bool   `json:"use_icon_as_answer_icon,omitempty"`
}

// CopyAppRequest 复制应用请求
type CopyAppRequest struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	IconType       string `json:"icon_type,omitempty"`
	Icon           string `json:"icon,omitempty"`
	IconBackground string `json:"icon_background,omitempty"`
}

// UpdateAppNameRequest 更新应用名称请求
type UpdateAppNameRequest struct {
	Name string `json:"name"`
}

// UpdateAppIconRequest 更新应用图标请求
type UpdateAppIconRequest struct {
	Icon           string `json:"icon,omitempty"`
	IconBackground string `json:"icon_background,omitempty"`
}

// UpdateAppSiteStatusRequest 更新应用站点状态请求
type UpdateAppSiteStatusRequest struct {
	EnableSite bool `json:"enable_site"`
}

// UpdateAppAPIStatusRequest 更新应用API状态请求
type UpdateAppAPIStatusRequest struct {
	EnableAPI bool `json:"enable_api"`
}

// AppExportResponse 应用导出响应
type AppExportResponse struct {
	Data string `json:"data"`
}

// AppTraceConfig 应用追踪配置
type AppTraceConfig struct {
	Enabled         bool   `json:"enabled"`
	TracingProvider string `json:"tracing_provider"`
}

// UpdateAppTraceRequest 更新应用追踪请求
type UpdateAppTraceRequest struct {
	Enabled         bool   `json:"enabled"`
	TracingProvider string `json:"tracing_provider"`
}

// Dataset 数据集
type Dataset struct {
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description"`
	Permission             string                 `json:"permission"`
	DataSourceType         string                 `json:"data_source_type"`
	IndexingTechnique      string                 `json:"indexing_technique"`
	CreatedBy              string                 `json:"created_by"`
	CreatedAt              UnixTime               `json:"created_at"`
	UpdatedAt              UnixTime               `json:"updated_at"`
	DocumentCount          int                    `json:"document_count"`
	WordCount              int                    `json:"word_count"`
	AppCount               int                    `json:"app_count"`
	EmbeddingModel         string                 `json:"embedding_model"`
	EmbeddingModelProvider string                 `json:"embedding_model_provider"`
	EmbeddingAvailable     bool                   `json:"embedding_available"`
	RetrievalModelDict     map[string]interface{} `json:"retrieval_model_dict"`
	Tags                   []string               `json:"tags"`
	PartialMemberList      []PartialMember        `json:"partial_member_list"`
}

// PartialMember 部分成员
type PartialMember struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	LastActive UnixTime `json:"last_active_at"`
}

// DatasetListResponse 数据集列表响应
type DatasetListResponse struct {
	Data    []Dataset `json:"data"`
	HasMore bool      `json:"has_more"`
	Limit   int       `json:"limit"`
	Total   int       `json:"total"`
	Page    int       `json:"page"`
}

// CreateDatasetRequest 创建数据集请求
type CreateDatasetRequest struct {
	Name                   string                 `json:"name"`
	Description            string                 `json:"description,omitempty"`
	Permission             string                 `json:"permission,omitempty"`
	IndexingTechnique      string                 `json:"indexing_technique,omitempty"`
	EmbeddingModel         string                 `json:"embedding_model,omitempty"`
	EmbeddingModelProvider string                 `json:"embedding_model_provider,omitempty"`
	RetrievalModel         map[string]interface{} `json:"retrieval_model,omitempty"`
}

// UpdateDatasetRequest 更新数据集请求
type UpdateDatasetRequest struct {
	Name           string                 `json:"name,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Permission     string                 `json:"permission,omitempty"`
	RetrievalModel map[string]interface{} `json:"retrieval_model,omitempty"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	RememberMe  bool   `json:"remember_me,omitempty"`
	InviteToken string `json:"invite_token,omitempty"`
	Language    string `json:"language,omitempty"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data,omitempty"`
	Code   string      `json:"code,omitempty"`
}

// APIKey API密钥
type APIKey struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Token      string    `json:"token"`
	LastUsedAt *UnixTime `json:"last_used_at"`
	CreatedAt  UnixTime  `json:"created_at"`
}

// APIKeyListResponse API密钥列表响应
type APIKeyListResponse struct {
	Data []APIKey `json:"data"`
}

// OperationResponse 通用操作响应
type OperationResponse struct {
	Result string `json:"result"`
}
