# 贡献指南

感谢您对 Dify Golang SDK 的关注！我们欢迎各种形式的贡献。

## 🚀 如何贡献

### 报告 Bug

1. 在 [Issues](https://github.com/kingfs/godify/issues) 页面搜索是否已有相关问题
2. 如果没有，请创建新的 Issue，包含：
   - 清晰的问题描述
   - 复现步骤
   - 期望的行为
   - 实际的行为
   - 环境信息（Go版本、操作系统等）

### 功能建议

1. 在 Issues 页面创建 Feature Request
2. 详细描述您的需求和使用场景
3. 如果可能，提供 API 设计建议

### 代码贡献

#### 开发环境准备

```bash
# 1. Fork 项目到您的 GitHub 账户
# 2. 克隆您的 Fork
git clone https://github.com/YOUR_USERNAME/dify-golang-sdk.git
cd dify-golang-sdk

# 3. 添加上游仓库
git remote add upstream https://github.com/kingfs/godify.git

# 4. 安装依赖
make deps

# 5. 创建开发分支
git checkout -b feature/your-feature-name
```

#### 开发流程

1. **编写代码**
   - 遵循 Go 语言规范和项目编码风格
   - 为新功能添加相应的测试
   - 更新相关文档

2. **运行测试**
   ```bash
   # 运行所有测试
   make test
   
   # 生成覆盖率报告
   make test-coverage
   
   # 代码格式化
   make format
   
   # 代码检查
   make lint
   ```

3. **提交代码**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   git push origin feature/your-feature-name
   ```

4. **创建 Pull Request**
   - 在 GitHub 上创建 PR
   - 填写详细的描述
   - 关联相关的 Issues

#### 代码规范

- **命名**: 使用有意义的变量和函数名
- **注释**: 为导出的函数和类型添加文档注释
- **错误处理**: 适当处理所有可能的错误情况
- **测试**: 为新功能编写单元测试，覆盖率应保持在 80% 以上

#### 提交信息规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 格式：

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

类型：
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式化
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或工具变更

示例：
```
feat(service): add dataset hit testing API

- Implement HitTestDataset function in service client
- Add HitTestingRequest and HitTestingResponse models
- Add comprehensive unit tests

Closes #123
```

## 📋 开发指南

### 项目结构

```
dify-golang-sdk/
├── client/         # 核心HTTP客户端
├── models/         # 数据模型定义
├── service/        # Service API客户端
├── web/           # Web API客户端
├── console/       # Console API客户端
├── files/         # Files API客户端
├── mcp/           # MCP API客户端
├── errors/        # 错误处理
├── examples/      # 使用示例
├── docs/          # 文档
└── tests/         # 测试文件
```

### 添加新的 API 端点

1. **分析 API**：研究 Dify API 文档和实现
2. **添加模型**：在 `models/` 中定义请求和响应结构
3. **实现客户端方法**：在相应的客户端中添加方法
4. **编写测试**：添加单元测试和集成测试
5. **更新文档**：更新 API 文档和示例

### 测试指南

- **单元测试**：测试单个函数的行为
- **集成测试**：测试完整的API调用流程
- **模拟测试**：使用 `httptest` 模拟HTTP服务器

示例测试：

```go
func TestNewClient(t *testing.T) {
    client := NewClient("test-token", "https://api.example.com")
    if client == nil {
        t.Fatal("Expected client to be created")
    }
}
```

### 文档更新

- **API文档**：在 `docs/API.md` 中添加新的API说明
- **README**：更新功能列表和示例
- **代码注释**：为公开的函数和类型添加注释

## 🔍 Code Review

所有代码都需要经过 Code Review：

1. **自动检查**：CI 会运行测试、格式检查和 lint
2. **人工审查**：维护者会审查代码逻辑和设计
3. **反馈处理**：根据反馈修改代码
4. **合并**：通过审查后会被合并到主分支

## 📞 联系我们

- **Issues**: [GitHub Issues](https://github.com/kingfs/godify/issues)
- **Discussions**: [GitHub Discussions](https://github.com/kingfs/godify/discussions)
- **Email**: sdk@dify.ai

## 📜 行为准则

请遵循我们的 [行为准则](CODE_OF_CONDUCT.md)，维护一个友好、包容的社区环境。

---

再次感谢您的贡献！🎉