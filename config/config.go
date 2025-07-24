package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	// 基础配置
	BaseURL    string        `mapstructure:"base_url" yaml:"base_url" json:"base_url"`
	AuthType   string        `mapstructure:"auth_type" yaml:"auth_type" json:"auth_type"`
	Token      string        `mapstructure:"token" yaml:"token" json:"token"`
	Timeout    time.Duration `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	MaxRetries int           `mapstructure:"max_retries" yaml:"max_retries" json:"max_retries"`
	
	// 工作空间配置
	WorkspaceID string `mapstructure:"workspace_id" yaml:"workspace_id" json:"workspace_id"`
	
	// 日志配置
	LogLevel  string `mapstructure:"log_level" yaml:"log_level" json:"log_level"`
	LogFormat string `mapstructure:"log_format" yaml:"log_format" json:"log_format"`
	LogOutput string `mapstructure:"log_output" yaml:"log_output" json:"log_output"`
	LogFile   string `mapstructure:"log_file" yaml:"log_file" json:"log_file"`
	
	// 监控配置
	EnableMetrics bool   `mapstructure:"enable_metrics" yaml:"enable_metrics" json:"enable_metrics"`
	MetricsPort   int    `mapstructure:"metrics_port" yaml:"metrics_port" json:"metrics_port"`
	HealthCheck   bool   `mapstructure:"health_check" yaml:"health_check" json:"health_check"`
	
	// 连接池配置
	MaxIdleConns        int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxIdleConnsPerHost int           `mapstructure:"max_idle_conns_per_host" yaml:"max_idle_conns_per_host" json:"max_idle_conns_per_host"`
	IdleConnTimeout     time.Duration `mapstructure:"idle_conn_timeout" yaml:"idle_conn_timeout" json:"idle_conn_timeout"`
}

// LoadConfig 加载配置
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	
	// 设置默认值
	setDefaults(config)
	
	// 从环境变量加载
	loadFromEnv(config)
	
	// 从配置文件加载（如果指定了路径）
	if configPath != "" {
		if err := loadFromFile(config, configPath); err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}
	
	// 验证配置
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	
	return config, nil
}

// setDefaults 设置默认值
func setDefaults(config *Config) {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.dify.ai"
	}
	if config.AuthType == "" {
		config.AuthType = "bearer"
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}
	if config.LogFormat == "" {
		config.LogFormat = "text"
	}
	if config.LogOutput == "" {
		config.LogOutput = "stdout"
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 100
	}
	if config.MaxIdleConnsPerHost == 0 {
		config.MaxIdleConnsPerHost = 10
	}
	if config.IdleConnTimeout == 0 {
		config.IdleConnTimeout = 90 * time.Second
	}
	if config.MetricsPort == 0 {
		config.MetricsPort = 8080
	}
}

// loadFromEnv 从环境变量加载配置
func loadFromEnv(config *Config) {
	// 基础配置
	if v := os.Getenv("DIFY_BASE_URL"); v != "" {
		config.BaseURL = v
	}
	if v := os.Getenv("DIFY_AUTH_TYPE"); v != "" {
		config.AuthType = v
	}
	if v := os.Getenv("DIFY_TOKEN"); v != "" {
		config.Token = v
	}
	if v := os.Getenv("DIFY_TIMEOUT"); v != "" {
		if timeout, err := time.ParseDuration(v); err == nil {
			config.Timeout = timeout
		}
	}
	if v := os.Getenv("DIFY_MAX_RETRIES"); v != "" {
		if retries, err := strconv.Atoi(v); err == nil {
			config.MaxRetries = retries
		}
	}
	
	// 工作空间配置
	if v := os.Getenv("DIFY_WORKSPACE_ID"); v != "" {
		config.WorkspaceID = v
	}
	
	// 日志配置
	if v := os.Getenv("DIFY_LOG_LEVEL"); v != "" {
		config.LogLevel = v
	}
	if v := os.Getenv("DIFY_LOG_FORMAT"); v != "" {
		config.LogFormat = v
	}
	if v := os.Getenv("DIFY_LOG_OUTPUT"); v != "" {
		config.LogOutput = v
	}
	if v := os.Getenv("DIFY_LOG_FILE"); v != "" {
		config.LogFile = v
	}
	
	// 监控配置
	if v := os.Getenv("DIFY_ENABLE_METRICS"); v != "" {
		config.EnableMetrics = strings.ToLower(v) == "true"
	}
	if v := os.Getenv("DIFY_METRICS_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			config.MetricsPort = port
		}
	}
	if v := os.Getenv("DIFY_HEALTH_CHECK"); v != "" {
		config.HealthCheck = strings.ToLower(v) == "true"
	}
	
	// 连接池配置
	if v := os.Getenv("DIFY_MAX_IDLE_CONNS"); v != "" {
		if conns, err := strconv.Atoi(v); err == nil {
			config.MaxIdleConns = conns
		}
	}
	if v := os.Getenv("DIFY_MAX_IDLE_CONNS_PER_HOST"); v != "" {
		if conns, err := strconv.Atoi(v); err == nil {
			config.MaxIdleConnsPerHost = conns
		}
	}
	if v := os.Getenv("DIFY_IDLE_CONN_TIMEOUT"); v != "" {
		if timeout, err := time.ParseDuration(v); err == nil {
			config.IdleConnTimeout = timeout
		}
	}
}

// loadFromFile 从配置文件加载
func loadFromFile(config *Config, configPath string) error {
	v := viper.New()
	v.SetConfigFile(configPath)
	
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	
	return v.Unmarshal(config)
}

// validateConfig 验证配置
func validateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("base_url is required")
	}
	if config.Token == "" {
		return fmt.Errorf("token is required")
	}
	if config.Timeout <= 0 {
		return fmt.Errorf("timeout must be positive")
	}
	if config.MaxRetries < 0 {
		return fmt.Errorf("max_retries must be non-negative")
	}
	
	return nil
}

// GetEnvWithDefault 获取环境变量，如果不存在则返回默认值
func GetEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}