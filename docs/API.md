# Dify Golang SDK API 文档

## 目录

- [安装和初始化](#安装和初始化)
- [Service API](#service-api)
- [Web API](#web-api)
- [错误处理](#错误处理)
- [数据模型](#数据模型)

## 安装和初始化

### 安装

```bash
go get github.com/kingfs/godify
```

### 初始化客户端

```go
import "github.com/kingfs/godify"

// Service API 客户端 (面向开发者)
serviceClient := dify.NewServiceClient("your-app-api-token", "https://api.dify.ai")

// Web API 客户端 (面向最终用户)
webClient := dify.NewWebClient("your-web-api-key", "https://api.dify.ai")
```

## Service API

Service API 面向开发者，需要使用应用的 API Token 进行认证。

### 应用信息

#### 获取应用信息
```go
appInfo, err := serviceClient.GetAppInfo(ctx)
```

#### 获取应用参数
```go
params, err := serviceClient.GetAppParameters(ctx)
```

#### 获取应用元数据
```go
meta, err := serviceClient.GetAppMeta(ctx)
```

### 对话功能

#### 文本补全
```go
req := &service.CompletionRequest{
    Inputs: map[string]interface{}{
        "query": "你的问题",
    },
    Query: "你的问题",
    User:  "user-123",
    ResponseMode: models.ResponseModeBlocking, // 或 ResponseModeStreaming
}
resp, err := serviceClient.Completion(ctx, req)
```

#### 聊天对话
```go
req := &service.ChatRequest{
    Inputs: map[string]interface{}{},
    Query:  "你好",
    User:   "user-123",
    ConversationID: "existing-conversation-id", // 可选，继续已有对话
    AutoGenerateName: true,
}
resp, err := serviceClient.Chat(ctx, req)
```

#### 停止生成
```go
// 停止文本补全
err := serviceClient.CompletionStop(ctx, taskID, userID)

// 停止聊天
err := serviceClient.ChatStop(ctx, taskID, userID)
```

## Web API

Web API 面向最终用户，需要使用 Web API Key 进行认证。

### 应用信息

#### 获取应用元数据
```go
meta, err := webClient.GetAppMeta(ctx)
```

#### 获取应用参数
```go
params, err := webClient.GetAppParameters(ctx)
```

#### 检查Web应用访问模式
```go
accessMode, err := webClient.GetWebAppAccessMode(ctx, "app-id", "")
```

#### 检查Web应用权限
```go
permission, err := webClient.CheckWebAppPermission(ctx, "app-id")
```

### 对话功能

#### 文本补全
```go
req := &web.CompletionRequest{
    Inputs: map[string]interface{}{
        "query": "你的问题",
    },
    Query: "你的问题",
    ResponseMode: models.ResponseModeBlocking,
}
resp, err := webClient.Completion(ctx, req)
```

#### 聊天对话
```go
req := &web.ChatRequest{
    Inputs: map[string]interface{}{},
    Query:  "你好",
    ConversationID: "existing-conversation-id", // 可选
}
resp, err := webClient.Chat(ctx, req)
```

### 对话管理

#### 获取对话列表
```go
conversations, err := webClient.GetConversations(ctx, "", 20, nil, "-updated_at")
```

#### 重命名对话
```go
newName := "新的对话名称"
req := &models.ConversationRenameRequest{
    Name: &newName,
}
conv, err := webClient.RenameConversation(ctx, "conversation-id", req)
```

#### 置顶/取消置顶对话
```go
// 置顶
err := webClient.PinConversation(ctx, "conversation-id")

// 取消置顶
err := webClient.UnpinConversation(ctx, "conversation-id")
```

#### 删除对话
```go
err := webClient.DeleteConversation(ctx, "conversation-id")
```

### 消息管理

#### 获取消息列表
```go
messages, err := webClient.GetMessages(ctx, "conversation-id", "", 20)
```

#### 发送消息反馈
```go
rating := "like"
content := "很有帮助"
feedback := &models.MessageFeedbackRequest{
    Rating:  &rating,
    Content: &content,
}
err := webClient.SendMessageFeedback(ctx, "message-id", feedback)
```

#### 获取类似内容
```go
resp, err := webClient.GetMessageMoreLikeThis(ctx, "message-id", models.ResponseModeBlocking)
```

#### 获取建议问题
```go
questions, err := webClient.GetSuggestedQuestions(ctx, "message-id")
```

### 文件处理

#### 上传文件
```go
fileData := []byte("file content")
file, err := webClient.UploadFile(ctx, "filename.txt", fileData, "datasets")
```

#### 语音转文字
```go
audioData := []byte("audio file content")
result, err := webClient.AudioToText(ctx, audioData, "audio.wav")
```

#### 文字转语音
```go
req := &web.TextToAudioRequest{
    Text:  "要转换的文字",
    Voice: "default",
}
result, err := webClient.TextToAudio(ctx, req)
```

## 错误处理

SDK 提供了详细的错误处理机制：

```go
import "github.com/kingfs/godify/errors"

resp, err := client.Chat(ctx, req)
if err != nil {
    if errors.IsAPIError(err) {
        apiErr := errors.GetAPIError(err)
        fmt.Printf("API错误: %s (状态码: %d)\n", apiErr.Message, apiErr.StatusCode)
        
        // 检查特定错误类型
        switch apiErr.Code {
        case "app_unavailable":
            // 处理应用不可用错误
        case "conversation_not_exists":
            // 处理对话不存在错误
        }
    } else {
        // 处理其他类型错误
        fmt.Printf("其他错误: %v\n", err)
    }
}
```

### 预定义错误类型

- `ErrAppUnavailable` - 应用不可用
- `ErrNotChatApp` - 不是聊天应用
- `ErrNotCompletionApp` - 不是补全应用
- `ErrConversationNotExists` - 对话不存在
- `ErrMessageNotExists` - 消息不存在
- `ErrProviderQuotaExceeded` - 提供商配额超限
- `ErrInvokeRateLimit` - 调用频率限制
- `ErrFileTooLarge` - 文件过大
- `ErrUnsupportedFileType` - 不支持的文件类型

## 数据模型

### 主要模型类型

#### AppInfo - 应用信息
```go
type AppInfo struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Tags        []string `json:"tags"`
    Mode        AppMode  `json:"mode"`
    AuthorName  string   `json:"author_name"`
}
```

#### GenerateResponse - 生成响应
```go
type GenerateResponse struct {
    Event          string `json:"event,omitempty"`
    MessageID      string `json:"message_id,omitempty"`
    ConversationID string `json:"conversation_id,omitempty"`
    Answer         string `json:"answer,omitempty"`
    TaskID         string `json:"task_id,omitempty"`
    CreatedAt      int64  `json:"created_at,omitempty"`
}
```

#### Conversation - 对话
```go
type Conversation struct {
    ID           string                 `json:"id"`
    Name         string                 `json:"name"`
    Inputs       map[string]interface{} `json:"inputs"`
    Introduction string                 `json:"introduction"`
    CreatedAt    time.Time              `json:"created_at"`
    UpdatedAt    time.Time              `json:"updated_at"`
}
```

#### Message - 消息
```go
type Message struct {
    ID                string                 `json:"id"`
    ConversationID    string                 `json:"conversation_id"`
    Query             string                 `json:"query"`
    Answer            string                 `json:"answer"`
    MessageFiles      []MessageFile          `json:"message_files"`
    Feedback          *MessageFeedback       `json:"feedback"`
    CreatedAt         time.Time              `json:"created_at"`
    Status            string                 `json:"status"`
}
```

### 枚举类型

#### AppMode - 应用模式
- `AppModeCompletion` - 文本补全
- `AppModeChat` - 聊天对话
- `AppModeAgentChat` - Agent聊天
- `AppModeAdvancedChat` - 高级聊天
- `AppModeWorkflow` - 工作流

#### ResponseMode - 响应模式
- `ResponseModeBlocking` - 阻塞模式
- `ResponseModeStreaming` - 流式模式

#### AccessMode - 访问模式
- `AccessModePublic` - 公开访问
- `AccessModePrivate` - 私有访问