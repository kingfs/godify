# Dify API 路径完整文档

本文档列出了Dify平台中所有API的完整路径结构，用于验证和修复Golang SDK中的API路径。

## API 蓝图结构

Dify使用Flask蓝图来组织API，主要包含以下蓝图：

### 1. Service API (`/v1`)
- **蓝图**: `service_api`
- **前缀**: `/v1`
- **认证**: Bearer Token (App API Token)
- **用途**: 面向开发者的API

### 2. Web API (`/api`)
- **蓝图**: `web`
- **前缀**: `/api`
- **认证**: Passport Token + X-App-Code
- **用途**: 面向最终用户的API

### 3. Console API (`/console/api`)
- **蓝图**: `console`
- **前缀**: `/console/api`
- **认证**: Session Cookie 或 Bearer Token
- **用途**: 面向管理员的控制台API

### 4. Files API (无前缀)
- **蓝图**: `files`
- **前缀**: 无
- **认证**: 根据上下文
- **用途**: 文件处理API

### 5. MCP API (`/mcp`)
- **蓝图**: `mcp`
- **前缀**: `/mcp`
- **认证**: 根据上下文
- **用途**: Model Context Protocol API

### 6. Inner API (`/inner/api`)
- **蓝图**: `inner_api`
- **前缀**: `/inner/api`
- **认证**: 内部认证
- **用途**: 内部服务API

## 完整API路径列表

### Service API (`/v1`)

#### 应用信息
- `GET /v1/parameters` - 获取应用参数
- `GET /v1/meta` - 获取应用元数据
- `GET /v1/info` - 获取应用信息

#### 对话功能
- `POST /v1/completion-messages` - 文本补全
- `POST /v1/completion-messages/{task_id}/stop` - 停止文本补全
- `POST /v1/chat-messages` - 聊天对话
- `POST /v1/chat-messages/{task_id}/stop` - 停止聊天

#### 对话管理
- `GET /v1/conversations` - 获取对话列表
- `GET /v1/conversations/{c_id}` - 获取对话详情
- `POST /v1/conversations/{c_id}/name` - 重命名对话
- `DELETE /v1/conversations/{c_id}` - 删除对话
- `GET /v1/conversations/{c_id}/variables` - 获取对话变量

#### 消息管理
- `GET /v1/messages` - 获取消息列表
- `POST /v1/messages/{message_id}/feedbacks` - 发送消息反馈
- `GET /v1/messages/{message_id}/suggested` - 获取建议问题

#### 文件处理
- `POST /v1/files/upload` - 上传文件

#### 音频处理
- `POST /v1/audio-to-text` - 语音转文字
- `POST /v1/text-to-audio` - 文字转语音

#### 工作流
- `POST /v1/workflows/run` - 运行工作流
- `GET /v1/workflows/run/{workflow_run_id}` - 获取工作流运行详情
- `POST /v1/workflows/tasks/{task_id}/stop` - 停止工作流任务
- `GET /v1/workflows/logs` - 获取工作流日志

### Web API (`/api`)

#### 认证
- `GET /api/passport` - 获取访问令牌

#### 应用信息
- `GET /api/parameters` - 获取应用参数
- `GET /api/meta` - 获取应用元数据
- `GET /api/webapp/access-mode` - 检查访问模式
- `GET /api/webapp/permission` - 检查权限

#### 对话功能
- `POST /api/completion-messages` - 文本补全
- `POST /api/completion-messages/{task_id}/stop` - 停止文本补全
- `POST /api/chat-messages` - 聊天对话
- `POST /api/chat-messages/{task_id}/stop` - 停止聊天

#### 对话管理
- `GET /api/conversations` - 获取对话列表
- `GET /api/conversations/{c_id}` - 获取对话详情
- `POST /api/conversations/{c_id}/name` - 重命名对话
- `DELETE /api/conversations/{c_id}` - 删除对话
- `PATCH /api/conversations/{c_id}/pin` - 置顶对话
- `PATCH /api/conversations/{c_id}/unpin` - 取消置顶对话

#### 消息管理
- `GET /api/messages` - 获取消息列表
- `POST /api/messages/{message_id}/feedbacks` - 发送消息反馈
- `GET /api/messages/{message_id}/more-like-this` - 获取类似内容
- `GET /api/messages/{message_id}/suggested-questions` - 获取建议问题

#### 文件处理
- `POST /api/files/upload` - 上传文件
- `GET /api/remote-files/{url}` - 获取远程文件信息
- `POST /api/remote-files/upload` - 上传远程文件

#### 音频处理
- `POST /api/audio-to-text` - 语音转文字
- `POST /api/text-to-audio` - 文字转语音

#### 工作流
- `POST /api/workflows/run` - 运行工作流
- `POST /api/workflows/tasks/{task_id}/stop` - 停止工作流任务

#### 保存消息
- `GET /api/saved-messages` - 获取保存的消息列表
- `GET /api/saved-messages/{message_id}` - 获取保存的消息详情

### Console API (`/console/api`)

#### 应用管理
- `GET /console/api/apps` - 获取应用列表
- `POST /console/api/apps` - 创建应用
- `GET /console/api/apps/{app_id}` - 获取应用详情
- `PUT /console/api/apps/{app_id}` - 更新应用
- `DELETE /console/api/apps/{app_id}` - 删除应用
- `POST /console/api/apps/{app_id}/copy` - 复制应用
- `GET /console/api/apps/{app_id}/export` - 导出应用
- `PUT /console/api/apps/{app_id}/name` - 更新应用名称
- `PUT /console/api/apps/{app_id}/icon` - 更新应用图标
- `PUT /console/api/apps/{app_id}/site-enable` - 启用/禁用站点
- `PUT /console/api/apps/{app_id}/api-enable` - 启用/禁用API

#### 应用对话功能
- `POST /console/api/apps/{app_id}/completion-messages` - 应用文本补全
- `POST /console/api/apps/{app_id}/completion-messages/{task_id}/stop` - 停止应用文本补全
- `POST /console/api/apps/{app_id}/chat-messages` - 应用聊天对话
- `POST /console/api/apps/{app_id}/chat-messages/{task_id}/stop` - 停止应用聊天

#### 应用对话管理
- `GET /console/api/apps/{app_id}/completion-conversations` - 获取补全对话列表
- `GET /console/api/apps/{app_id}/completion-conversations/{conversation_id}` - 获取补全对话详情
- `GET /console/api/apps/{app_id}/chat-conversations` - 获取聊天对话列表
- `GET /console/api/apps/{app_id}/chat-conversations/{conversation_id}` - 获取聊天对话详情

#### 应用消息管理
- `GET /console/api/apps/{app_id}/messages/{message_id}` - 获取消息详情
- `GET /console/api/apps/{app_id}/chat-messages` - 获取聊天消息列表
- `GET /console/api/apps/{app_id}/chat-messages/{message_id}/suggested-questions` - 获取建议问题

#### 应用音频处理
- `POST /console/api/apps/{app_id}/audio-to-text` - 语音转文字
- `POST /console/api/apps/{app_id}/text-to-audio` - 文字转语音
- `GET /console/api/apps/{app_id}/text-to-audio/voices` - 获取语音模式

#### 数据集管理
- `GET /console/api/datasets` - 获取数据集列表
- `POST /console/api/datasets` - 创建数据集
- `GET /console/api/datasets/{dataset_id}` - 获取数据集详情
- `PUT /console/api/datasets/{dataset_id}` - 更新数据集
- `DELETE /console/api/datasets/{dataset_id}` - 删除数据集
- `GET /console/api/datasets/{dataset_id}/use-check` - 检查数据集使用情况
- `GET /console/api/datasets/{dataset_id}/queries` - 获取数据集查询
- `GET /console/api/datasets/{dataset_id}/error-docs` - 获取错误文档
- `GET /console/api/datasets/indexing-estimate` - 获取索引估算
- `GET /console/api/datasets/{dataset_id}/related-apps` - 获取相关应用
- `GET /console/api/datasets/{dataset_id}/indexing-status` - 获取索引状态
- `GET /console/api/datasets/api-keys` - 获取数据集API密钥
- `DELETE /console/api/datasets/api-keys/{api_key_id}` - 删除数据集API密钥
- `GET /console/api/datasets/api-base-info` - 获取API基础信息
- `GET /console/api/datasets/retrieval-setting` - 获取检索设置
- `GET /console/api/datasets/retrieval-setting/{vector_type}` - 获取检索设置模拟
- `GET /console/api/datasets/{dataset_id}/permission-part-users` - 获取权限用户列表
- `GET /console/api/datasets/{dataset_id}/auto-disable-logs` - 获取自动禁用日志

#### 数据集文档管理
- `GET /console/api/datasets/{dataset_id}/documents` - 获取文档列表
- `POST /console/api/datasets/{dataset_id}/documents` - 创建文档
- `GET /console/api/datasets/{dataset_id}/documents/{document_id}` - 获取文档详情
- `PUT /console/api/datasets/{dataset_id}/documents/{document_id}` - 更新文档
- `DELETE /console/api/datasets/{dataset_id}/documents/{document_id}` - 删除文档
- `POST /console/api/datasets/{dataset_id}/documents/{document_id}/indexing` - 索引文档
- `POST /console/api/datasets/{dataset_id}/documents/{document_id}/indexing/stop` - 停止索引文档
- `GET /console/api/datasets/{dataset_id}/documents/{document_id}/indexing-status` - 获取文档索引状态
- `POST /console/api/datasets/{dataset_id}/documents/{document_id}/indexing/retry` - 重试索引文档

#### 数据集分段管理
- `GET /console/api/datasets/{dataset_id}/segments` - 获取分段列表
- `GET /console/api/datasets/{dataset_id}/segments/{segment_id}` - 获取分段详情
- `PUT /console/api/datasets/{dataset_id}/segments/{segment_id}` - 更新分段
- `DELETE /console/api/datasets/{dataset_id}/segments/{segment_id}` - 删除分段
- `POST /console/api/datasets/{dataset_id}/segments/{segment_id}/indexing` - 索引分段
- `POST /console/api/datasets/{dataset_id}/segments/{segment_id}/indexing/stop` - 停止索引分段
- `GET /console/api/datasets/{dataset_id}/segments/{segment_id}/indexing-status` - 获取分段索引状态
- `POST /console/api/datasets/{dataset_id}/segments/{segment_id}/indexing/retry` - 重试索引分段

#### 数据集元数据管理
- `GET /console/api/datasets/{dataset_id}/metadata` - 获取元数据列表
- `POST /console/api/datasets/{dataset_id}/metadata` - 创建元数据
- `GET /console/api/datasets/{dataset_id}/metadata/{metadata_id}` - 获取元数据详情
- `PUT /console/api/datasets/{dataset_id}/metadata/{metadata_id}` - 更新元数据
- `DELETE /console/api/datasets/{dataset_id}/metadata/{metadata_id}` - 删除元数据

#### 数据集命中测试
- `POST /console/api/datasets/{dataset_id}/hit-testing` - 命中测试
- `POST /console/api/datasets/{dataset_id}/external-hit-testing` - 外部命中测试

#### 外部数据集
- `POST /console/api/datasets/external` - 创建外部数据集
- `GET /console/api/datasets/external-knowledge-api` - 获取外部知识API列表
- `GET /console/api/datasets/external-knowledge-api/{external_knowledge_api_id}` - 获取外部知识API详情
- `GET /console/api/datasets/external-knowledge-api/{external_knowledge_api_id}/use-check` - 检查外部知识API使用情况

#### API密钥管理
- `GET /console/api/apps/{app_id}/api-keys` - 获取应用API密钥列表
- `POST /console/api/apps/{app_id}/api-keys` - 创建应用API密钥
- `DELETE /console/api/apps/{app_id}/api-keys/{api_key_id}` - 删除应用API密钥

#### 工作区管理
- `GET /console/api/workspaces` - 获取工作区列表
- `GET /console/api/all-workspaces` - 获取所有工作区
- `GET /console/api/workspaces/current` - 获取当前工作区
- `GET /console/api/info` - 获取信息 (已弃用)
- `POST /console/api/workspaces/switch` - 切换工作区
- `PUT /console/api/workspaces/custom-config` - 更新自定义配置
- `POST /console/api/workspaces/custom-config/webapp-logo/upload` - 上传Web应用Logo
- `GET /console/api/workspaces/info` - 获取工作区信息

#### 成员管理
- `GET /console/api/workspaces/current/members` - 获取成员列表
- `POST /console/api/workspaces/current/members` - 邀请成员
- `PUT /console/api/workspaces/current/members/{member_id}` - 更新成员
- `DELETE /console/api/workspaces/current/members/{member_id}` - 删除成员
- `POST /console/api/workspaces/current/members/{member_id}/role` - 更新成员角色

#### 模型提供商管理
- `GET /console/api/workspaces/current/model-providers` - 获取模型提供商列表
- `GET /console/api/workspaces/current/model-providers/{provider}/credentials` - 获取提供商凭据
- `POST /console/api/workspaces/current/model-providers/{provider}/credentials/validate` - 验证提供商凭据
- `GET /console/api/workspaces/current/model-providers/{provider}` - 获取提供商详情
- `PUT /console/api/workspaces/current/model-providers/{provider}/preferred-provider-type` - 更新首选提供商类型
- `GET /console/api/workspaces/current/model-providers/{provider}/checkout-url` - 获取结账URL
- `GET /console/api/workspaces/{tenant_id}/model-providers/{provider}/{icon_type}/{lang}` - 获取提供商图标

#### 工具提供商管理
- `GET /console/api/workspaces/current/tool-providers` - 获取工具提供商列表
- `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/tools` - 获取内置工具提供商工具
- `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/info` - 获取内置工具提供商信息
- `DELETE /console/api/workspaces/current/tool-provider/builtin/{provider}/delete` - 删除内置工具提供商
- `PUT /console/api/workspaces/current/tool-provider/builtin/{provider}/update` - 更新内置工具提供商
- `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/credentials` - 获取内置工具提供商凭据
- `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/credentials_schema` - 获取内置工具提供商凭据模式
- `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/icon` - 获取内置工具提供商图标

#### API工具提供商
- `POST /console/api/workspaces/current/tool-provider/api/add` - 添加API工具提供商
- `POST /console/api/workspaces/current/tool-provider/api/remote` - 获取远程模式
- `GET /console/api/workspaces/current/tool-provider/api/tools` - 获取API工具提供商工具
- `PUT /console/api/workspaces/current/tool-provider/api/update` - 更新API工具提供商
- `DELETE /console/api/workspaces/current/tool-provider/api/delete` - 删除API工具提供商
- `GET /console/api/workspaces/current/tool-provider/api/get` - 获取API工具提供商
- `GET /console/api/workspaces/current/tool-provider/api/schema` - 获取API工具提供商模式
- `POST /console/api/workspaces/current/tool-provider/api/test/pre` - API工具提供商预测试

#### 工作流工具提供商
- `POST /console/api/workspaces/current/tool-provider/workflow/create` - 创建工作流工具提供商
- `PUT /console/api/workspaces/current/tool-provider/workflow/update` - 更新工作流工具提供商

#### 插件管理
- `GET /console/api/workspaces/current/plugin/debugging-key` - 获取调试密钥
- `GET /console/api/workspaces/current/plugin/list` - 获取插件列表
- `GET /console/api/workspaces/current/plugin/list/latest-versions` - 获取最新版本插件列表
- `GET /console/api/workspaces/current/plugin/list/installations/ids` - 获取安装ID列表
- `GET /console/api/workspaces/current/plugin/icon` - 获取插件图标
- `POST /console/api/workspaces/current/plugin/upload/pkg` - 从包上传插件
- `POST /console/api/workspaces/current/plugin/upload/github` - 从GitHub上传插件
- `POST /console/api/workspaces/current/plugin/upload/bundle` - 从包上传插件
- `POST /console/api/workspaces/current/plugin/install/pkg` - 从包安装插件
- `POST /console/api/workspaces/current/plugin/install/github` - 从GitHub安装插件
- `POST /console/api/workspaces/current/plugin/upgrade/marketplace` - 从市场升级插件
- `POST /console/api/workspaces/current/plugin/upgrade/github` - 从GitHub升级插件
- `POST /console/api/workspaces/current/plugin/install/marketplace` - 从市场安装插件
- `GET /console/api/workspaces/current/plugin/fetch-manifest` - 获取清单
- `GET /console/api/workspaces/current/plugin/tasks` - 获取任务
- `GET /console/api/workspaces/current/plugin/tasks/{task_id}` - 获取任务详情
- `DELETE /console/api/workspaces/current/plugin/tasks/{task_id}/delete` - 删除任务
- `DELETE /console/api/workspaces/current/plugin/tasks/delete_all` - 删除所有任务项
- `DELETE /console/api/workspaces/current/plugin/tasks/{task_id}/delete/{identifier}` - 删除任务项
- `POST /console/api/workspaces/current/plugin/uninstall` - 卸载插件
- `GET /console/api/workspaces/current/plugin/marketplace/pkg` - 获取市场包
- `POST /console/api/workspaces/current/plugin/permission/change` - 更改权限

#### 文件管理
- `POST /console/api/files/upload` - 上传文件
- `GET /console/api/files/{file_id}/preview` - 预览文件
- `GET /console/api/files/support-type` - 获取支持类型

#### 远程文件
- `GET /console/api/remote-files/{url}` - 获取远程文件信息
- `POST /console/api/remote-files/upload` - 上传远程文件

#### 应用导入
- `POST /console/api/apps/imports` - 导入应用
- `POST /console/api/apps/imports/{import_id}/confirm` - 确认导入
- `GET /console/api/apps/imports/{app_id}/check-dependencies` - 检查依赖

#### 已安装应用探索
- `GET /console/api/installed-apps/{installed_app_id}/audio-to-text` - 语音转文字
- `GET /console/api/installed-apps/{installed_app_id}/text-to-audio` - 文字转语音
- `POST /console/api/installed-apps/{installed_app_id}/completion-messages` - 文本补全
- `POST /console/api/installed-apps/{installed_app_id}/completion-messages/{task_id}/stop` - 停止文本补全
- `POST /console/api/installed-apps/{installed_app_id}/chat-messages` - 聊天对话
- `POST /console/api/installed-apps/{installed_app_id}/chat-messages/{task_id}/stop` - 停止聊天对话
- `POST /console/api/installed-apps/{installed_app_id}/conversations/{c_id}/name` - 重命名对话
- `GET /console/api/installed-apps/{installed_app_id}/conversations` - 获取对话列表
- `GET /console/api/installed-apps/{installed_app_id}/conversations/{c_id}` - 获取对话详情
- `PATCH /console/api/installed-apps/{installed_app_id}/conversations/{c_id}/pin` - 置顶对话
- `PATCH /console/api/installed-apps/{installed_app_id}/conversations/{c_id}/unpin` - 取消置顶对话
- `GET /console/api/installed-apps/{installed_app_id}/messages` - 获取消息列表
- `POST /console/api/installed-apps/{installed_app_id}/messages/{message_id}/feedbacks` - 发送消息反馈
- `GET /console/api/installed-apps/{installed_app_id}/messages/{message_id}/more-like-this` - 获取类似内容
- `GET /console/api/installed-apps/{installed_app_id}/messages/{message_id}/suggested-questions` - 获取建议问题
- `POST /console/api/installed-apps/{installed_app_id}/workflows/run` - 运行工作流
- `POST /console/api/installed-apps/{installed_app_id}/workflows/tasks/{task_id}/stop` - 停止工作流任务

### Files API (无前缀)

#### 文件上传
- `POST /files/upload/for-plugin` - 为插件上传文件

#### 文件预览
- `GET /files/{file_id}/image-preview` - 图片预览
- `GET /files/{file_id}/file-preview` - 文件预览
- `GET /files/workspaces/{workspace_id}/webapp-logo` - 工作区Web应用Logo

#### 工具文件
- `GET /files/tools/{file_id}.{extension}` - 工具文件预览

### MCP API (`/mcp`)

#### MCP服务
- `POST /mcp/` - MCP请求处理

### Inner API (`/inner/api`)

#### 邮件
- `POST /inner/api/mail` - 发送邮件

#### 插件
- `POST /inner/api/invoke/llm` - 调用LLM
- `POST /inner/api/invoke/llm/structured-output` - 调用结构化输出LLM
- `POST /inner/api/invoke/text-embedding` - 调用文本嵌入
- `POST /inner/api/invoke/rerank` - 调用重排序
- `POST /inner/api/invoke/tts` - 调用TTS
- `POST /inner/api/invoke/speech2text` - 调用语音转文字
- `POST /inner/api/invoke/moderation` - 调用内容审核
- `POST /inner/api/invoke/tool` - 调用工具
- `POST /inner/api/invoke/parameter-extractor` - 调用参数提取器
- `POST /inner/api/invoke/question-classifier` - 调用问题分类器
- `POST /inner/api/invoke/app` - 调用应用
- `POST /inner/api/invoke/encrypt` - 调用加密
- `POST /inner/api/invoke/summary` - 调用摘要
- `POST /inner/api/upload/file/request` - 上传文件请求
- `POST /inner/api/fetch/app/info` - 获取应用信息

#### 工作区
- `GET /inner/api/workspace` - 获取工作区信息

## 路径验证结果

通过对比Dify源码和Golang SDK，发现以下路径问题：

### ✅ 正确的路径
- Service API: `/v1/*` ✅
- Web API: `/api/*` ✅
- Console API: `/console/api/*` ✅
- Files API: 无前缀 ✅
- MCP API: `/mcp/*` ✅
- Inner API: `/inner/api/*` ✅

### 🔧 需要修复的问题
1. **Web API认证流程**: 已修复，使用Passport认证
2. **路径前缀**: 已确认正确
3. **路由注册**: 已确认正确

## 总结

经过深入调查，Dify Golang SDK中的API路径基本正确，主要问题在于Web API的认证流程，这已经在之前的修复中解决了。所有API路径都与Dify源码中的路由注册一致。 