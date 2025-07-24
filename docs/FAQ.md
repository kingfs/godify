# Dify Go SDK 常见问题解答 (FAQ)

## 🔧 安装与配置

### Q: 如何安装 Dify Go SDK？
```bash
go get github.com/kingfs/godify
```

### Q: 如何配置认证信息？
```go
// 方式1: 直接传入参数
client := dify.NewServiceClient("your-token", "https://api.dify.ai")

// 方式2: 使用环境变量
export DIFY_TOKEN="your-token"
export DIFY_BASE_URL="https://api.dify.ai"

// 方式3: 使用配置文件
client, err := client.NewClientFromConfig("config.yaml")
```

### Q: 支持哪些认证方式？
- **Bearer Token**: 适用于 Service API
- **API Key**: 适用于 Web API
- **Session Cookie**: 适用于 Console API

## 🚀 使用问题

### Q: 如何处理 API 错误？
```go
resp, err := client.Chat(ctx, req)
if err != nil {
    if errors.IsAPIError(err) {
        apiErr := errors.GetAPIError(err)
        switch apiErr.Code {
        case "app_unavailable":
            // 处理应用不可用
        case "quota_exceeded":
            // 处理配额超限
        case "rate_limit_exceeded":
            // 处理速率限制
        }
    }
}
```

### Q: 如何实现自动重试？
```go
// 使用内置重试机制
config := &client.ClientConfig{
    MaxRetries: 3,
    RetryDelay: time.Second,
}

// 或自定义重试逻辑
for attempt := 0; attempt < maxRetries; attempt++ {
    resp, err := client.Chat(ctx, req)
    if err == nil {
        break
    }
    time.Sleep(backoff * time.Duration(attempt))
}
```

### Q: 如何处理流式响应？
```go
handler := &client.JSONSSEHandler{
    OnEventFunc: func(eventType string, data map[string]interface{}) error {
        fmt.Printf("Event: %s, Data: %+v\n", eventType, data)
        return nil
    },
    OnErrorFunc: func(err error) {
        fmt.Printf("Error: %v\n", err)
    },
    OnCompleteFunc: func() {
        fmt.Println("Stream completed")
    },
}

client.ChatStream(ctx, req, handler)
```

## 🔍 调试与监控

### Q: 如何启用日志记录？
```go
// 配置日志
logConfig := &logger.Config{
    Level:  logger.InfoLevel,
    Format: "json",
    Output: "stdout",
}
log, _ := logger.NewLogger(logConfig)

// 创建客户端时传入日志器
config := &client.ClientConfig{
    Logger: log,
    // ... 其他配置
}
```

### Q: 如何监控性能指标？
```go
// 启用监控
metrics := metrics.NewMetrics(true)
config := &client.ClientConfig{
    Metrics: metrics,
}

// 启动监控服务器
go func() {
    ctx := context.Background()
    client.Metrics.StartMetricsServer(ctx, 8080)
}()

// 获取统计信息
stats := client.Metrics.GetStats()
fmt.Printf("成功率: %.2f%%\n", stats["success_rate"])
```

### Q: 如何查看请求日志？
```bash
# 设置环境变量
export DIFY_LOG_LEVEL="debug"

# 或在代码中设置
log.SetLevel(logger.DebugLevel)
```

## 📁 文件处理

### Q: 如何上传文件？
```go
fileData := []byte("file content")
file, err := client.UploadFile(ctx, "document.txt", fileData, "datasets")
if err != nil {
    log.Printf("上传失败: %v", err)
}
```

### Q: 支持哪些文件格式？
- **文本文件**: .txt, .md, .doc, .docx
- **音频文件**: .mp3, .wav, .m4a
- **图片文件**: .jpg, .png, .gif

### Q: 文件大小限制是多少？
- 单个文件最大: 10MB
- 音频文件最大: 25MB
- 总文件数限制: 10个

## 🔄 并发与性能

### Q: 如何实现并发请求？
```go
// 使用 goroutine
for _, req := range requests {
    go func(r *service.ChatRequest) {
        resp, err := client.Chat(ctx, r)
        // 处理响应
    }(req)
}

// 使用工作池
pool := make(chan struct{}, maxConcurrency)
for _, req := range requests {
    pool <- struct{}{}
    go func(r *service.ChatRequest) {
        defer func() { <-pool }()
        resp, err := client.Chat(ctx, r)
        // 处理响应
    }(req)
}
```

### Q: 如何优化连接池？
```go
config := &client.ClientConfig{
    MaxIdleConns:        200,
    MaxIdleConnsPerHost: 20,
    IdleConnTimeout:     time.Minute * 2,
}
```

## 🛠 开发与测试

### Q: 如何运行测试？
```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 运行特定包的测试
go test ./service/...
```

### Q: 如何模拟 API 响应？
```go
// 使用 httptest
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"answer": "模拟响应"}`))
}))
defer server.Close()

client := dify.NewServiceClient("token", server.URL)
```

### Q: 如何处理测试环境配置？
```go
// 使用环境变量区分环境
if os.Getenv("ENV") == "test" {
    client = dify.NewServiceClient("test-token", "http://localhost:8080")
} else {
    client = dify.NewServiceClient("prod-token", "https://api.dify.ai")
}
```

## 🔐 安全与认证

### Q: 如何安全存储 API 密钥？
```go
// 使用环境变量
token := os.Getenv("DIFY_TOKEN")

// 使用配置文件
config, err := config.LoadConfig("config.yaml")

// 使用密钥管理服务
token := getSecretFromVault("dify-api-key")
```

### Q: 如何验证 API 响应？
```go
// 验证响应状态
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("API 返回错误状态: %d", resp.StatusCode)
}

// 验证响应内容
if resp.Answer == "" {
    return fmt.Errorf("响应内容为空")
}
```

## 📊 错误代码对照表

| 错误代码 | 含义 | 建议处理方式 |
|---------|------|-------------|
| `app_unavailable` | 应用不可用 | 检查应用状态，稍后重试 |
| `quota_exceeded` | 配额超限 | 升级套餐或等待重置 |
| `rate_limit_exceeded` | 速率限制 | 降低请求频率 |
| `invalid_token` | 无效令牌 | 检查认证信息 |
| `permission_denied` | 权限不足 | 检查应用权限 |
| `conversation_not_exists` | 对话不存在 | 检查对话ID |
| `file_too_large` | 文件过大 | 压缩文件或分片上传 |

## 🆘 获取帮助

### Q: 在哪里可以找到更多帮助？
- **GitHub Issues**: [提交问题](https://github.com/kingfs/godify/issues)
- **文档**: [完整API文档](./API.md)
- **示例**: [使用示例](./examples/)
- **社区**: [Dify 社区](https://community.dify.ai)

### Q: 如何报告 Bug？
1. 检查是否是最新版本
2. 查看 [Issues](https://github.com/kingfs/godify/issues) 是否已有类似问题
3. 提供详细的错误信息和复现步骤
4. 包含系统信息和 Go 版本

### Q: 如何贡献代码？
1. Fork 项目
2. 创建功能分支
3. 编写测试
4. 提交 Pull Request
5. 等待代码审查

---

**💡 提示**: 如果这里没有找到您的问题，请查看 [完整API文档](./API.md) 或提交 [Issue](https://github.com/kingfs/godify/issues)。