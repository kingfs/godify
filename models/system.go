package models

type SetupStatus string

const (
	SetupStatusFinished   SetupStatus = "finished"
	SetupStatusNotStarted SetupStatus = "not_started"
)

// Version Dify版本信息
type Version struct {
	Version       string `json:"version"`
	ReleaseDate   string `json:"release_date"`
	ReleaseNotes  string `json:"release_notes"`
	CanAutoUpdate bool   `json:"can_auto_update"`
	Features      struct {
		CanReplaceLogo            bool `json:"can_replace_logo"`
		ModelLoadBalancingEnabled bool `json:"model_load_balancing_enabled"`
	} `json:"features"`
}

// SetupRequest 安装请求
type SetupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// StatusResponse 状态响应
type StatusResponse struct {
	Step    SetupStatus `json:"step"`
	SetupAt *string     `json:"setup_at"`
}

// ResultResponse 结果响应
type ResultResponse struct {
	Result string `json:"result"`
}
