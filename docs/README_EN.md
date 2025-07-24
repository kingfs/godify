# Dify Golang SDK

Dify Golang SDK is a complete Go language client library for the Dify AI platform, providing a simple and easy-to-use API for interacting with the Dify platform.

## 🌟 Features

- 🚀 **Complete API Coverage**: 63 API endpoints supporting all Dify features
  - Web API (22 endpoints) - For end users
  - Service API (14 endpoints) - For developers
  - Console API (22 endpoints) - For administrators
  - Files API + MCP API (5 endpoints) - Professional features
- 🔐 **Multiple Authentication Methods**: Bearer Token, API Key, Session Cookie
- 📡 **Streaming Response Support**: Real-time streaming conversations and workflows
- 📎 **File Upload Support**: Multimedia file processing and audio conversion
- 🛠 **100% Type Safety**: Complete Go type definitions with compile-time error checking
- 🔄 **Automatic Retry Mechanism**: Built-in intelligent retry and error recovery
- 🧪 **Complete Test Coverage**: Unit tests and integration tests
- 📝 **Detailed Documentation**: Complete API documentation and rich examples

## 📦 Installation

```bash
go get github.com/kingfs/godify
```

## 🚀 Quick Start

### Service API - Developer Integration

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/kingfs/godify"
    "github.com/kingfs/godify/service"
)

func main() {
    // Create Service API client
    client := dify.NewServiceClient("your-app-api-token", "https://api.dify.ai")
    
    // Chat conversation
    req := &service.ChatRequest{
        Query:  "Hello, how are you?",
        User:   "user123",
        Inputs: map[string]interface{}{},
    }
    
    resp, err := client.Chat(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Response: %s\n", resp.Answer)
}
```

### Web API - End User Interaction

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/kingfs/godify"
    "github.com/kingfs/godify/web"
)

func main() {
    // Create Web API client
    client := dify.NewWebClient("your-api-key", "https://api.dify.ai")
    
    // Get conversation list
    conversations, err := client.GetConversations(context.Background(), "", 20, nil, "-updated_at")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d conversations\n", len(conversations.Data))
}
```

### Console API - Administrator Features

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/kingfs/godify"
    "github.com/kingfs/godify/models"
)

func main() {
    // Create Console API client
    client := dify.NewConsoleClient("your-access-token", "https://api.dify.ai")
    
    // Create new application
    req := &models.CreateAppRequest{
        Name: "My AI App",
        Mode: models.AppModeChat,
        Description: "A test application",
    }
    
    app, err := client.CreateApp(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Created app: %s (ID: %s)\n", app.Name, app.ID)
}
```

### Workflow Execution

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/kingfs/godify"
    "github.com/kingfs/godify/models"
)

func main() {
    client := dify.NewWebClient("your-api-key", "https://api.dify.ai")
    
    // Run workflow
    req := &models.WorkflowRunRequest{
        Inputs: map[string]interface{}{
            "user_input": "Analyze the sentiment of this text",
            "text": "Today is a wonderful day!",
        },
    }
    
    resp, err := client.RunWorkflow(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Workflow started: %s\n", resp.WorkflowRunID)
}
```

## 📚 Complete Feature List

### Web API (22 endpoints)
- ✅ Application information and permission management
- ✅ Text completion and chat conversations
- ✅ Conversation management (CRUD operations)
- ✅ Message management and feedback
- ✅ Workflow execution and control
- ✅ File upload and audio processing

### Service API (14 endpoints)
- ✅ Application information retrieval
- ✅ Conversation features (completion/chat)
- ✅ Dataset management and document processing
- ✅ Hit testing and retrieval

### Console API (22 endpoints)
- ✅ User authentication and login
- ✅ Application management (CRUD/import/export)
- ✅ Dataset management
- ✅ API key management

### Professional Features (5 endpoints)
- ✅ Files API - Plugin file upload
- ✅ MCP API - Model Context Protocol

## 🛠 Development Tools

```bash
# Install dependencies
make deps

# Run tests
make test

# Generate coverage report
make test-coverage

# Format code
make format

# Run examples
make examples

# View all commands
make help
```

## 📖 Documentation

- [Complete API Documentation](./API.md) - Detailed API reference
- [Usage Examples](./examples/) - Rich code examples
- [Changelog](./CHANGELOG.md) - Version change records

## 🔧 Advanced Features

### Error Handling

```go
import "github.com/kingfs/godify/errors"

resp, err := client.Chat(ctx, req)
if err != nil {
    if errors.IsAPIError(err) {
        apiErr := errors.GetAPIError(err)
        switch apiErr.Code {
        case "app_unavailable":
            // Handle app unavailable
        case "quota_exceeded":
            // Handle quota exceeded
        }
    }
}
```

### File Upload

```go
fileData := []byte("file content")
file, err := webClient.UploadFile(ctx, "document.txt", fileData, "datasets")
```

### Streaming Response

```go
req := &service.ChatRequest{
    ResponseMode: models.ResponseModeStreaming,
    // ... other parameters
}
// Note: Streaming responses require handling Server-Sent Events
```

## 🤝 Contributing

We welcome all forms of contributions!

1. 🐛 **Report Bugs** - [Submit Issue](https://github.com/kingfs/godify/issues)
2. 💡 **Feature Suggestions** - [Feature Request](https://github.com/kingfs/godify/issues)
3. 🔧 **Code Contributions** - [Pull Request](https://github.com/kingfs/godify/pulls)
4. 📖 **Documentation Improvements** - Documentation and example updates

## 📄 License

MIT License - See [LICENSE](LICENSE) file for details

## 🔗 Related Links

- [Dify Official Website](https://dify.ai)
- [Dify GitHub](https://github.com/langgenius/dify)
- [API Documentation](https://docs.dify.ai/api)

---

**⭐ If this project helps you, please give us a star!**