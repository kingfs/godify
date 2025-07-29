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
	MaxActiveRequests   int    `json:"max_active_requests,omitempty"`
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

// TokenPair 令牌对
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Result string    `json:"result"`
	Data   TokenPair `json:"data,omitempty"`
	Code   string    `json:"code,omitempty"`
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

// ConsoleAppDetail 控制台应用详情
type ConsoleAppDetail struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Description         string        `json:"description"`
	Mode                string        `json:"mode"`
	IconType            string        `json:"icon_type"`
	Icon                string        `json:"icon"`
	IconBackground      string        `json:"icon_background"`
	IconURL             *string       `json:"icon_url"`
	EnableSite          bool          `json:"enable_site"`
	EnableAPI           bool          `json:"enable_api"`
	ModelConfig         interface{}   `json:"model_config"`
	Workflow            *WorkflowInfo `json:"workflow"`
	Site                *SiteInfo     `json:"site"`
	APIBaseURL          string        `json:"api_base_url"`
	UseIconAsAnswerIcon bool          `json:"use_icon_as_answer_icon"`
	CreatedBy           string        `json:"created_by"`
	CreatedAt           UnixTime      `json:"created_at"`
	UpdatedBy           *string       `json:"updated_by"`
	UpdatedAt           UnixTime      `json:"updated_at"`
	DeletedTools        []interface{} `json:"deleted_tools"`
	AccessMode          *string       `json:"access_mode"`
}

// WorkflowInfo 工作流信息
type WorkflowInfo struct {
	ID        string   `json:"id"`
	CreatedBy string   `json:"created_by"`
	CreatedAt UnixTime `json:"created_at"`
	UpdatedBy *string  `json:"updated_by"`
	UpdatedAt UnixTime `json:"updated_at"`
}

// SiteInfo 站点信息
type SiteInfo struct {
	AccessToken            string   `json:"access_token"`
	Code                   string   `json:"code"`
	Title                  string   `json:"title"`
	IconType               string   `json:"icon_type"`
	Icon                   string   `json:"icon"`
	IconBackground         string   `json:"icon_background"`
	IconURL                *string  `json:"icon_url"`
	Description            *string  `json:"description"`
	DefaultLanguage        string   `json:"default_language"`
	ChatColorTheme         *string  `json:"chat_color_theme"`
	ChatColorThemeInverted bool     `json:"chat_color_theme_inverted"`
	CustomizeDomain        *string  `json:"customize_domain"`
	Copyright              *string  `json:"copyright"`
	PrivacyPolicy          *string  `json:"privacy_policy"`
	CustomDisclaimer       string   `json:"custom_disclaimer"`
	CustomizeTokenStrategy string   `json:"customize_token_strategy"`
	PromptPublic           bool     `json:"prompt_public"`
	AppBaseURL             string   `json:"app_base_url"`
	ShowWorkflowSteps      bool     `json:"show_workflow_steps"`
	UseIconAsAnswerIcon    bool     `json:"use_icon_as_answer_icon"`
	CreatedBy              string   `json:"created_by"`
	CreatedAt              UnixTime `json:"created_at"`
	UpdatedBy              string   `json:"updated_by"`
	UpdatedAt              UnixTime `json:"updated_at"`
}

// AppsChatMessageListApiResponse 聊天消息列表响应
// 对应 /apps/<uuid:app_id>/chat-messages 接口
// 参考返回示例结构体

type AppsChatMessageListApiResponse struct {
	Data    []AppsChatMessage `json:"data"`
	HasMore bool              `json:"has_more"`
	Limit   int               `json:"limit"`
}

type AppsChatMessage struct {
	AgentThoughts           []interface{}          `json:"agent_thoughts"`
	Annotation              interface{}            `json:"annotation"`
	AnnotationHitHistory    interface{}            `json:"annotation_hit_history"`
	Answer                  string                 `json:"answer"`
	AnswerTokens            int                    `json:"answer_tokens"`
	ConversationID          string                 `json:"conversation_id"`
	CreatedAt               float64                `json:"created_at"`
	Error                   interface{}            `json:"error"`
	Feedbacks               []interface{}          `json:"feedbacks"`
	FromAccountID           string                 `json:"from_account_id"`
	FromEndUserID           interface{}            `json:"from_end_user_id"`
	FromSource              string                 `json:"from_source"`
	ID                      string                 `json:"id"`
	Inputs                  map[string]interface{} `json:"inputs"`
	Message                 []AppsMessage          `json:"message"`
	MessageFiles            []interface{}          `json:"message_files"`
	MessageTokens           int                    `json:"message_tokens"`
	Metadata                map[string]interface{} `json:"metadata"`
	ParentMessageID         interface{}            `json:"parent_message_id"`
	ProviderResponseLatency float64                `json:"provider_response_latency"`
	Query                   string                 `json:"query"`
	Status                  string                 `json:"status"`
	WorkflowRunID           interface{}            `json:"workflow_run_id"`
}

type AppsMessage struct {
	Files []interface{} `json:"files"`
	Role  string        `json:"role"`
	Text  string        `json:"text"`
}

// AppsMessageApiResponse 消息详情响应
// 对应 /apps/<uuid:app_id>/messages/<uuid:message_id> 接口
// 参考返回示例结构体

type AppsMessageApiResponse struct {
	AgentThoughts           []interface{}          `json:"agent_thoughts"`
	Annotation              interface{}            `json:"annotation"`
	AnnotationHitHistory    interface{}            `json:"annotation_hit_history"`
	Answer                  string                 `json:"answer"`
	AnswerTokens            int                    `json:"answer_tokens"`
	ConversationID          string                 `json:"conversation_id"`
	CreatedAt               float64                `json:"created_at"`
	Error                   interface{}            `json:"error"`
	Feedbacks               []interface{}          `json:"feedbacks"`
	FromAccountID           string                 `json:"from_account_id"`
	FromEndUserID           interface{}            `json:"from_end_user_id"`
	FromSource              string                 `json:"from_source"`
	ID                      string                 `json:"id"`
	Inputs                  map[string]interface{} `json:"inputs"`
	Message                 []AppsMessage          `json:"message"`
	MessageFiles            []interface{}          `json:"message_files"`
	MessageTokens           int                    `json:"message_tokens"`
	Metadata                map[string]interface{} `json:"metadata"`
	ParentMessageID         interface{}            `json:"parent_message_id"`
	ProviderResponseLatency float64                `json:"provider_response_latency"`
	Query                   string                 `json:"query"`
	Status                  string                 `json:"status"`
	WorkflowRunID           interface{}            `json:"workflow_run_id"`
}

// Workspace 工作区信息
// 对应 /workspaces 接口返回的单个 workspace
// 例如 map[created_at:1.753153422e+09 current:true id:8b72a5d2-31cb-4ca5-a5e2-b7e4b79064b3 name:admin's Workspace plan:sandbox status:normal]
type Workspace struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Plan      string  `json:"plan"`
	Status    string  `json:"status"`
	Current   bool    `json:"current"`
	CreatedAt float64 `json:"created_at"`
}

// WorkspacesApiResponse 工作区列表响应
// 对应 map[workspaces:[...]]
type WorkspacesApiResponse struct {
	Workspaces []Workspace `json:"workspaces"`
}

// WorkspaceMember 工作区成员信息
// 对应 /workspaces/current/members 接口返回的 accounts 字段
// 例如 map[avatar:<nil> avatar_url:<nil> created_at:1.753153422e+09 email:admin@chaitin.net id:53846eb7-1ac0-461d-8684-d22e56ee477c last_active_at:1.753347058e+09 last_login_at:1.753153422e+09 name:admin role:owner status:active]
type WorkspaceMember struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Role         string  `json:"role"`
	Status       string  `json:"status"`
	Avatar       *string `json:"avatar"`
	AvatarURL    *string `json:"avatar_url"`
	CreatedAt    float64 `json:"created_at"`
	LastActiveAt float64 `json:"last_active_at"`
	LastLoginAt  float64 `json:"last_login_at"`
}

// WorkspacesCurrentMembersApiResponse 工作区当前成员接口响应
// 对应 map[accounts:[...]]
type WorkspacesCurrentMembersApiResponse struct {
	Accounts []WorkspaceMember `json:"accounts"`
}

// WorkspaceInviteEmailApiResponse 工作区成员邀请邮件接口响应
// 对应 /workspaces/current/members/invite-email
// 返回示例: {"result": "success", "invitation_results": [...], "tenant_id": "..."}
type WorkspaceInviteEmailApiResponse struct {
	Result            string                      `json:"result"`
	InvitationResults []WorkspaceInvitationResult `json:"invitation_results"`
	TenantID          string                      `json:"tenant_id"`
}

type WorkspaceInvitationResult struct {
	Status  string `json:"status"`
	Email   string `json:"email"`
	URL     string `json:"url,omitempty"`
	Message string `json:"message,omitempty"`
}

// WorkspaceOperationResponse 工作区成员操作通用响应
// 用于 /workspaces/current/members/<uuid:member_id> 删除成员等
// 返回示例: {"result": "success", "tenant_id": "..."}
type WorkspaceOperationResponse struct {
	Result   string `json:"result"`
	TenantID string `json:"tenant_id"`
}

// WorkspaceUpdateRoleResponse 工作区成员角色更新响应
// 用于 /workspaces/current/members/<uuid:member_id>/update-role
// 返回示例: {"result": "success"}
type WorkspaceUpdateRoleResponse struct {
	Result string `json:"result"`
}

// WorkspacesCurrentDatasetOperatorsApiResponse 工作区数据集操作员成员列表响应
// 对应 /workspaces/current/dataset-operators
// 返回示例: {"result": "success", "accounts": [...]}
type WorkspacesCurrentDatasetOperatorsApiResponse struct {
	Result   string            `json:"result"`
	Accounts []WorkspaceMember `json:"accounts"`
}

type ModelType string

const (
	LLM            ModelType = "llm"
	TEXT_EMBEDDING ModelType = "text-embedding"
	RERANK         ModelType = "rerank"
	SPEECH2TEXT    ModelType = "speech2text"
	MODERATION     ModelType = "moderation"
	TTS            ModelType = "tts"
)

type ModelProvidersResponse struct {
	Data []ModelProvider `json:"data"`
}

type ModelProvider struct {
	Background               *string               `json:"background"`
	ConfigurateMethods       []string              `json:"configurate_methods"`
	CustomConfiguration      CustomConfiguration   `json:"custom_configuration"`
	Description              map[string]string     `json:"description"`
	Help                     *string               `json:"help"`
	IconLarge                *string               `json:"icon_large"`
	IconSmall                map[string]string     `json:"icon_small"`
	Label                    map[string]string     `json:"label"`
	ModelCredentialSchema    ModelCredentialSchema `json:"model_credential_schema"`
	Model                    ModelInfo             `json:"model"`
	PreferredProviderType    string                `json:"preferred_provider_type"`
	Provider                 string                `json:"provider"`
	ProviderCredentialSchema interface{}           `json:"provider_credential_schema"`
	SupportedModelTypes      []string              `json:"supported_model_types"`
	SystemConfiguration      SystemConfiguration   `json:"system_configuration"`
	TenantID                 string                `json:"tenant_id"`
}

type CustomConfiguration struct {
	Status string `json:"status"`
}

type SystemConfiguration struct {
	CurrentQuotaType    *string     `json:"current_quota_type"`
	Enabled             bool        `json:"enabled"`
	QuotaConfigurations interface{} `json:"quota_configurations"`
}

// 这里是 model_credential_schema 字段的结构体
type ModelCredentialSchema struct {
	CredentialFormSchemas []CredentialFormSchema `json:"credential_form_schemas"`
}

type CredentialFormSchema struct {
	Default     interface{}        `json:"default"`
	Label       map[string]string  `json:"label"`
	MaxLength   int                `json:"max_length"`
	Options     []CredentialOption `json:"options"`
	Placeholder map[string]string  `json:"placeholder"`
	Required    bool               `json:"required"`
	ShowOn      []ShowOnCondition  `json:"show_on"`
	Type        string             `json:"type"`
	Variable    string             `json:"variable"`
}

type CredentialOption struct {
	Label  map[string]string `json:"label"`
	ShowOn []ShowOnCondition `json:"show_on"`
	Value  string            `json:"value"`
}

type ShowOnCondition struct {
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

// model 字段
type ModelInfo struct {
	Label       map[string]string `json:"label"`
	Placeholder map[string]string `json:"placeholder"`
}

type ProviderCredentialsResponse struct{}

// ModelProviderModelsResponse 用于 /workspaces/current/model-providers/<provider>/models 返回
// 参考 Python 返回示例和实际 JSON 字段

type ModelProviderModelsResponse struct {
	Data []ModelProviderModel `json:"data"`
}

type ModelProviderModel struct {
	Model                string                 `json:"model"`
	Label                map[string]string      `json:"label"`
	ModelType            string                 `json:"model_type"`
	Features             []string               `json:"features"`
	FetchFrom            string                 `json:"fetch_from"`
	ModelProperties      map[string]interface{} `json:"model_properties"`
	Deprecated           bool                   `json:"deprecated"`
	Status               string                 `json:"status"`
	LoadBalancingEnabled bool                   `json:"load_balancing_enabled"`
	Provider             ModelProviderInfo      `json:"provider"`
}

type ModelProviderInfo struct {
	Provider            string            `json:"provider"`
	Label               map[string]string `json:"label"`
	IconSmall           map[string]string `json:"icon_small"`
	IconLarge           *string           `json:"icon_large"`
	SupportedModelTypes []string          `json:"supported_model_types"`
	Models              []string          `json:"models"`
	TenantID            string            `json:"tenant_id"`
}

// ToolProviderListDetailResponse 工具提供者列表详细响应
type (
	ToolProviderListDetailResponse []ToolProviderDetail
	ToolProviderListResponse       struct {
		Data []ToolProvider `json:"data"`
	}
)

type ToolProvider struct {
	Provider  string            `json:"provider"`
	Label     map[string]string `json:"label"`
	IconSmall map[string]string `json:"icon_small"`
	IconLarge *string           `json:"icon_large"`
}

// ToolProviderDetail 工具提供者详细信息
type ToolProviderDetail struct {
	AllowDelete            bool                   `json:"allow_delete"`
	Author                 string                 `json:"author"`
	Description            map[string]string      `json:"description"`
	Icon                   string                 `json:"icon"`
	IconDark               *string                `json:"icon_dark"`
	ID                     string                 `json:"id"`
	IsTeamAuthorization    bool                   `json:"is_team_authorization"`
	Label                  map[string]string      `json:"label"`
	Labels                 []string               `json:"labels"`
	Name                   string                 `json:"name"`
	PluginID               *string                `json:"plugin_id"`
	PluginUniqueIdentifier string                 `json:"plugin_unique_identifier"`
	TeamCredentials        map[string]interface{} `json:"team_credentials"`
	Tools                  []interface{}          `json:"tools"`
	Type                   string                 `json:"type"`
}

// BuiltinToolListResponse 内置工具列表响应
type BuiltinToolListResponse []BuiltinTool

// BuiltinTool 内置工具信息
type BuiltinTool struct {
	Author       string            `json:"author"`
	Description  map[string]string `json:"description"`
	Label        map[string]string `json:"label"`
	Labels       []string          `json:"labels"`
	Name         string            `json:"name"`
	OutputSchema interface{}       `json:"output_schema"`
	Parameters   []ToolParameter   `json:"parameters"`
}

// ToolParameter 工具参数定义
type ToolParameter struct {
	AutoGenerate     *bool             `json:"auto_generate"`
	Default          interface{}       `json:"default"`
	Form             string            `json:"form"`
	HumanDescription map[string]string `json:"human_description"`
	InputSchema      interface{}       `json:"input_schema"`
	Label            map[string]string `json:"label"`
	LLMDescription   string            `json:"llm_description"`
	Max              interface{}       `json:"max"`
	Min              interface{}       `json:"min"`
	Name             string            `json:"name"`
	Options          []interface{}     `json:"options"`
	Placeholder      interface{}       `json:"placeholder"`
	Precision        interface{}       `json:"precision"`
	Required         bool              `json:"required"`
	Scope            interface{}       `json:"scope"`
	Template         interface{}       `json:"template"`
	Type             string            `json:"type"`
}

// ToolProviderApiEntity 工具提供商 API 实体
// 用于 /workspaces/current/tool-provider/builtin/{provider}/info 返回
type ToolProviderEntity struct {
	ID                     string                 `json:"id"`
	Author                 string                 `json:"author"`
	Name                   string                 `json:"name"`
	Description            map[string]string      `json:"description"`
	Icon                   interface{}            `json:"icon"`
	IconDark               *string                `json:"icon_dark"`
	Label                  map[string]string      `json:"label"`
	Type                   string                 `json:"type"`
	MaskedCredentials      map[string]interface{} `json:"masked_credentials"`
	IsTeamAuthorization    bool                   `json:"is_team_authorization"`
	PluginID               *string                `json:"plugin_id"`
	PluginUniqueIdentifier *string                `json:"plugin_unique_identifier"`
	Tools                  []interface{}          `json:"tools"`
	Labels                 []string               `json:"labels"`
	AllowDelete            bool                   `json:"allow_delete"`
	OriginalCredentials    map[string]interface{} `json:"original_credentials,omitempty"`
	ServerURL              *string                `json:"server_url,omitempty"`
	UpdatedAt              *int64                 `json:"updated_at,omitempty"`
	ServerIdentifier       *string                `json:"server_identifier,omitempty"`
}

// ToolBuiltinProviderCredentialsSchemaResponse 内置工具提供商凭据模式响应
// 对应 /workspaces/current/tool-provider/builtin/{provider}/credentials_schema 接口
type ToolBuiltinProviderCredentialsSchemaResponse []ToolBuiltinProviderCredentialSchema

// ToolBuiltinProviderCredentialSchema 内置工具提供商凭据模式
type ToolBuiltinProviderCredentialSchema struct {
	Default      interface{}                 `json:"default"`
	Label        map[string]string           `json:"label"`
	MaxLength    *int                        `json:"max_length,omitempty"`
	Options      []ToolBuiltinProviderOption `json:"options,omitempty"`
	Placeholder  map[string]string           `json:"placeholder,omitempty"`
	Required     bool                        `json:"required"`
	ShowOn       []ToolBuiltinProviderShowOn `json:"show_on,omitempty"`
	Type         string                      `json:"type"`
	Variable     string                      `json:"variable"`
	Help         map[string]string           `json:"help,omitempty"`
	Min          interface{}                 `json:"min,omitempty"`
	Max          interface{}                 `json:"max,omitempty"`
	Precision    interface{}                 `json:"precision,omitempty"`
	Scope        interface{}                 `json:"scope,omitempty"`
	Template     interface{}                 `json:"template,omitempty"`
	Form         string                      `json:"form,omitempty"`
	AutoGenerate *bool                       `json:"auto_generate,omitempty"`
}

// ToolBuiltinProviderOption 内置工具提供商选项
type ToolBuiltinProviderOption struct {
	Label  map[string]string           `json:"label"`
	ShowOn []ToolBuiltinProviderShowOn `json:"show_on,omitempty"`
	Value  string                      `json:"value"`
}

// ToolBuiltinProviderShowOn 内置工具提供商显示条件
type ToolBuiltinProviderShowOn struct {
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

// ApiToolListResponse 对应 GetToolApiProviderListTools 返回的工具列表
type ApiToolListResponse []ApiTool

// ApiTool API工具信息
type ApiTool struct {
	Author       string            `json:"author"`
	Description  map[string]string `json:"description"`
	Label        map[string]string `json:"label"`
	Labels       []string          `json:"labels"`
	Name         string            `json:"name"`
	OutputSchema interface{}       `json:"output_schema"`
	Parameters   []ToolParameter   `json:"parameters"`
}

// AppImportResponse 应用导入响应
type AppImportResponse struct {
	Status             string                 `json:"status"`
	AppID              string                 `json:"app_id,omitempty"`
	AppMode            string                 `json:"app_mode,omitempty"`
	CurrentDSLVersion  string                 `json:"current_dsl_version,omitempty"`
	Error              string                 `json:"error,omitempty"`
	ID                 string                 `json:"id,omitempty"`
	ImportedDSLVersion string                 `json:"imported_dsl_version,omitempty"`
	Data               map[string]interface{} `json:"data,omitempty"`
}
