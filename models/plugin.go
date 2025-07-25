package models

// PluginDebuggingKeyResponse 用于接收调试密钥接口返回值
type PluginDebuggingKeyResponse struct {
	Key  string `json:"key"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// PluginListResponse 用于接收插件列表接口返回值
type PluginListResponse struct {
	Plugins []PluginItem `json:"plugins"`
	Total   int          `json:"total"`
}

// PluginItem 需根据实际插件字段定义，暂用 map 占位
type PluginItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	LatestVersion string `json:"latest_version"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// PluginListLatestVersionsResponse 用于接收插件最新版本列表接口返回值
type PluginListLatestVersionsResponse struct {
	Versions []PluginVersionItem `json:"versions"`
}

// PluginVersionItem 需根据实际字段定义，暂用 map 占位

type PluginVersionItem map[string]any

// PluginListInstallationsFromIdsResponse 用于接收根据ID获取插件安装信息接口返回值
type PluginListInstallationsFromIdsResponse struct {
	Plugins []PluginInstallationItem `json:"plugins"`
}

// PluginInstallationItem 需根据实际字段定义，暂用 map 占位
// TODO: 根据实际字段细化
type PluginInstallationItem map[string]any

// PluginUploadResponse 用于接收插件上传相关接口返回值
// TODO: 根据实际返回字段细化
type PluginUploadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// PluginInstallResponse 用于接收插件安装相关接口返回值
// TODO: 根据实际返回字段细化
type PluginInstallResponse struct {
	AllInstalled bool   `json:"all_installed"`
	TaskID       string `json:"task_id"`
}

// PluginFetchManifestResponse 用于接收插件 manifest 接口返回值
// TODO: PluginManifest 需根据实际字段细化
type PluginFetchManifestResponse struct {
	Manifest PluginManifest `json:"manifest"`
}

type PluginManifest map[string]any

// PluginFetchInstallTasksResponse 用于接收插件安装任务列表接口返回值
// TODO: PluginInstallTask 需根据实际字段细化
type PluginFetchInstallTasksResponse struct {
	Tasks []PluginInstallTask `json:"tasks"`
}

type PluginInstallTask map[string]any

// PluginFetchInstallTaskResponse 用于接收单个插件安装任务接口返回值
// TODO: PluginInstallTask 需根据实际字段细化
type PluginFetchInstallTaskResponse struct {
	Task PluginInstallTask `json:"task"`
}

// PluginSimpleSuccessResponse 用于接收仅返回 success 字段的接口
type PluginSimpleSuccessResponse struct {
	Success bool `json:"success"`
}

// PluginFetchPermissionResponse 用于接收插件权限接口返回值
type PluginFetchPermissionResponse struct {
	InstallPermission string `json:"install_permission"`
	DebugPermission   string `json:"debug_permission"`
}

// PluginFetchDynamicSelectOptionsResponse 用于接收动态选项接口返回值
// TODO: PluginDynamicOption 需根据实际字段细化
type PluginFetchDynamicSelectOptionsResponse struct {
	Options []PluginDynamicOption `json:"options"`
}

type PluginDynamicOption map[string]any
