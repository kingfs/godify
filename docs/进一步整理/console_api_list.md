#### 功能管理

- [ ] `GET /console/api/features` - 获取功能列表
- [ ] `GET /console/api/system-features` - 获取系统功能

#### 管理员相关

- [ ] `POST /console/api/admin/insert-explore-apps` - 插入探索应用列表
- [ ] `POST /console/api/admin/insert-explore-apps/{app_id}` - 插入单个探索应用

#### 初始化与设置

- [ ] `POST /console/api/init` - 初始化校验
- [ ] `POST /console/api/setup` - 系统初始化
- [ ] `GET /console/api/version` - 获取版本信息

#### 文件管理

- [ ] `POST /console/api/files/upload` - 上传文件
- [ ] `GET /console/api/files/{file_id}/preview` - 文件预览
- [ ] `GET /console/api/files/support-type` - 获取支持的文件类型
- [ ] `GET /console/api/remote-files/{url}` - 获取远程文件信息
- [ ] `POST /console/api/remote-files/upload` - 上传远程文件

#### API Key 管理

- [ ] `GET /console/api/apps/{resource_id}/api-keys` - 获取应用 API Key 列表
- [ ] `POST /console/api/apps/{resource_id}/api-keys` - 创建应用 API Key
- [ ] `DELETE /console/api/apps/{resource_id}/api-keys/{api_key_id}` - 删除应用 API Key
- [ ] `GET /console/api/datasets/{resource_id}/api-keys` - 获取数据集 API Key 列表
- [ ] `POST /console/api/datasets/{resource_id}/api-keys` - 创建数据集 API Key
- [ ] `DELETE /console/api/datasets/{resource_id}/api-keys/{api_key_id}` - 删除数据集 API Key

#### 探索与推荐应用

- [ ] `GET /console/api/explore/apps` - 获取推荐应用列表
- [ ] `GET /console/api/explore/apps/{app_id}` - 获取推荐应用详情
- [ ] `GET /console/api/installed-apps` - 获取已安装应用列表
- [ ] `GET /console/api/installed-apps/{installed_app_id}` - 获取已安装应用详情
- [ ] `GET /console/api/installed-apps/{installed_app_id}/meta` - 获取已安装应用元信息
- [ ] `POST /console/api/installed-apps/{installed_app_id}/audio-to-text` - 音频转文本
- [ ] `POST /console/api/installed-apps/{installed_app_id}/text-to-audio` - 文本转音频
- [ ] `GET /console/api/installed-apps/{installed_app_id}/messages` - 获取消息列表
- [ ] `POST /console/api/installed-apps/{installed_app_id}/workflows/run` - 运行工作流

#### 认证与账户

- [ ] `POST /console/api/login` - 登录
- [ ] `POST /console/api/logout` - 登出
- [ ] `POST /console/api/email-code-login` - 邮箱验证码登录
- [ ] `POST /console/api/email-code-login/validity` - 校验邮箱验证码
- [ ] `POST /console/api/reset-password` - 重置密码
- [ ] `POST /console/api/refresh-token` - 刷新 token
- [ ] `POST /console/api/forgot-password` - 忘记密码发送邮件
- [ ] `POST /console/api/forgot-password/validity` - 校验忘记密码验证码
- [ ] `POST /console/api/forgot-password/resets` - 重置密码
- [ ] `POST /console/api/activate/check` - 校验激活
- [ ] `POST /console/api/activate` - 激活账户
- [ ] `POST /console/api/oauth/login/{provider}` - 第三方登录
- [ ] `POST /console/api/oauth/authorize/{provider}` - 第三方授权回调
- [ ] `POST /console/api/api-key-auth/data-source` - API Key 数据源认证
- [ ] `POST /console/api/api-key-auth/data-source/binding` - API Key 数据源绑定
- [ ] `DELETE /console/api/api-key-auth/data-source/{binding_id}` - 删除 API Key 数据源绑定
- [ ] `POST /console/api/oauth/data-source/{provider}` - OAuth 数据源认证
- [ ] `POST /console/api/oauth/data-source/callback/{provider}` - OAuth 数据源回调
- [ ] `POST /console/api/oauth/data-source/binding/{provider}` - OAuth 数据源绑定
- [ ] `POST /console/api/oauth/data-source/{provider}/{binding_id}/sync` - OAuth 数据源同步

#### 数据集管理

- [ ] `GET /console/api/datasets` - 获取数据集列表
- [ ] `GET /console/api/datasets/{dataset_id}` - 获取数据集详情
- [ ] `POST /console/api/datasets` - 创建数据集
- [ ] `PUT /console/api/datasets/{dataset_id}` - 更新数据集
- [ ] `DELETE /console/api/datasets/{dataset_id}` - 删除数据集
- [ ] `GET /console/api/datasets/{dataset_id}/documents` - 获取文档列表
- [ ] `POST /console/api/datasets/{dataset_id}/documents` - 创建文档
- [ ] `GET /console/api/datasets/{dataset_id}/documents/{document_id}` - 获取文档详情
- [ ] `PUT /console/api/datasets/{dataset_id}/documents/{document_id}` - 更新文档
- [ ] `DELETE /console/api/datasets/{dataset_id}/documents/{document_id}` - 删除文档
- [ ] `POST /console/api/datasets/{dataset_id}/hit-testing` - 数据集命中测试
- [ ] `POST /console/api/datasets/{dataset_id}/external-hit-testing` - 外部知识命中测试
- [ ] `POST /console/api/datasets/external` - 创建外部数据集
- [ ] `GET /console/api/datasets/external-knowledge-api` - 获取外部知识 API 模板列表
- [ ] `GET /console/api/datasets/external-knowledge-api/{external_knowledge_api_id}` - 获取外部知识 API 模板详情
- [ ] `POST /console/api/datasets/external-knowledge-api/{external_knowledge_api_id}/use-check` - 检查外部知识 API 使用情况
- [ ] `POST /console/api/datasets/{dataset_id}/notion/sync` - Notion 数据同步
- [ ] `GET /console/api/datasets/{dataset_id}/metadata` - 获取数据集元数据
- [ ] `GET /console/api/datasets/{dataset_id}/metadata/{metadata_id}` - 获取指定元数据
- [ ] `GET /console/api/datasets/metadata/built-in` - 获取内置元数据字段
- [ ] `POST /console/api/datasets/{dataset_id}/metadata/built-in/{action}` - 操作内置元数据字段
- [ ] `POST /console/api/datasets/{dataset_id}/documents/metadata` - 获取文档元数据
- [ ] `GET /console/api/datasets/{dataset_id}/documents/{document_id}/segments` - 获取文档分段
- [ ] `POST /console/api/datasets/{dataset_id}/documents/{document_id}/segment` - 添加文档分段
- [ ] `GET /console/api/datasets/{dataset_id}/documents/{document_id}/indexing-status` - 获取文档索引状态
- [ ] `POST /console/api/datasets/{dataset_id}/documents/{document_id}/website-sync` - 文档网站同步
- [ ] `POST /console/api/datasets/{dataset_id}/documents/status/{action}/batch` - 批量操作文档状态
- [ ] `POST /console/api/datasets/{dataset_id}/documents/{document_id}/processing/pause` - 暂停文档处理
- [ ] `POST /console/api/datasets/{dataset_id}/documents/{document_id}/processing/resume` - 恢复文档处理
- [ ] `POST /console/api/datasets/{dataset_id}/retry` - 文档重试
- [ ] `POST /console/api/datasets/{dataset_id}/documents/{document_id}/rename` - 重命名文档
- [ ] `POST /console/api/datasets/process-rule` - 获取处理规则
- [ ] `POST /console/api/datasets/init` - 初始化数据集
- [ ] `POST /console/api/datasets/{dataset_id}/batch/{batch}/indexing-estimate` - 批量索引估算
- [ ] `POST /console/api/datasets/{dataset_id}/batch/{batch}/indexing-status` - 批量索引状态
- [ ] `POST /console/api/datasets/{dataset_id}/error-docs` - 获取错误文档
- [ ] `POST /console/api/datasets/indexing-estimate` - 索引估算
- [ ] `POST /console/api/datasets/{dataset_id}/related-apps` - 获取相关应用
- [ ] `POST /console/api/datasets/{dataset_id}/indexing-status` - 获取索引状态
- [ ] `POST /console/api/datasets/api-keys` - 获取 API Key
- [ ] `DELETE /console/api/datasets/api-keys/{api_key_id}` - 删除 API Key
- [ ] `POST /console/api/datasets/api-base-info` - 获取 API 基础信息
- [ ] `POST /console/api/datasets/retrieval-setting` - 获取检索设置
- [ ] `POST /console/api/datasets/retrieval-setting/{vector_type}` - 获取指定类型检索设置
- [ ] `POST /console/api/datasets/{dataset_id}/permission-part-users` - 获取部分用户权限
- [ ] `POST /console/api/datasets/{dataset_id}/auto-disable-logs` - 获取自动禁用日志

#### 插件管理

- [ ] `GET /console/api/workspaces/current/plugin/list` - 获取插件列表
- [ ] `POST /console/api/workspaces/current/plugin/upload/pkg` - 上传插件包
- [ ] `POST /console/api/workspaces/current/plugin/install/pkg` - 安装本地插件
- [ ] `POST /console/api/workspaces/current/plugin/uninstall` - 卸载插件
- [ ] `GET /console/api/workspaces/current/plugin/fetch-manifest` - 获取插件 manifest
- [ ] `GET /console/api/workspaces/current/plugin/tasks` - 获取插件安装任务
- [ ] `GET /console/api/workspaces/current/plugin/tasks/{task_id}` - 获取插件安装任务详情
- [ ] `POST /console/api/workspaces/current/plugin/tasks/{task_id}/delete` - 删除插件安装任务
- [ ] `POST /console/api/workspaces/current/plugin/tasks/delete_all` - 删除所有插件安装任务
- [ ] `POST /console/api/workspaces/current/plugin/tasks/{task_id}/delete/{identifier}` - 删除插件安装任务项
- [ ] `POST /console/api/workspaces/current/plugin/upgrade/marketplace` - 市场插件升级
- [ ] `POST /console/api/workspaces/current/plugin/upgrade/github` - github 插件升级
- [ ] `POST /console/api/workspaces/current/plugin/install/marketplace` - 安装市场插件
- [ ] `GET /console/api/workspaces/current/plugin/marketplace/pkg` - 获取市场插件包
- [ ] `POST /console/api/workspaces/current/plugin/permission/change` - 修改插件权限
- [ ] `GET /console/api/workspaces/current/plugin/permission/fetch` - 获取插件权限
- [ ] `GET /console/api/workspaces/current/plugin/parameters/dynamic-options` - 获取插件动态参数选项
- [ ] `GET /console/api/workspaces/current/plugin/list/latest-versions` - 获取插件最新版本
- [ ] `POST /console/api/workspaces/current/plugin/list/installations/ids` - 通过 ID 获取插件安装信息
- [ ] `GET /console/api/workspaces/current/plugin/icon` - 获取插件图标
- [ ] `POST /console/api/workspaces/current/plugin/upload/github` - 上传 github 插件
- [ ] `POST /console/api/workspaces/current/plugin/upload/bundle` - 上传插件 bundle
- [ ] `POST /console/api/workspaces/current/plugin/install/github` - 安装 github 插件
- [ ] `GET /console/api/workspaces/current/plugin/debugging-key` - 获取插件调试 key

#### 工作区与成员

- [ ] `GET /console/api/workspaces` - 获取所有租户
- [ ] `GET /console/api/all-workspaces` - 获取所有工作区
- [ ] `GET /console/api/workspaces/current` - 获取当前租户信息
- [ ] `POST /console/api/workspaces/switch` - 切换租户
- [ ] `POST /console/api/workspaces/custom-config` - 自定义配置
- [ ] `POST /console/api/workspaces/custom-config/webapp-logo/upload` - 上传 webapp logo
- [ ] `POST /console/api/workspaces/info` - 修改工作区信息
- [ ] `GET /console/api/workspaces/info` - 获取工作区信息（已废弃）
- [ ] `GET /console/api/workspaces/current/members` - 获取成员列表
- [ ] `POST /console/api/workspaces/current/members/invite-email` - 邀请成员
- [ ] `DELETE /console/api/workspaces/current/members/{member_id}` - 取消邀请
- [ ] `POST /console/api/workspaces/current/members/{member_id}/update-role` - 更新成员角色
- [ ] `GET /console/api/workspaces/current/dataset-operators` - 获取数据集操作员成员

#### 账户设置

- [ ] `POST /console/api/account/init` - 初始化账户
- [ ] `GET /console/api/account/profile` - 获取账户信息
- [ ] `POST /console/api/account/name` - 修改账户名
- [ ] `POST /console/api/account/avatar` - 修改头像
- [ ] `POST /console/api/account/interface-language` - 修改界面语言
- [ ] `POST /console/api/account/interface-theme` - 修改界面主题
- [ ] `POST /console/api/account/timezone` - 修改时区
- [ ] `POST /console/api/account/password` - 修改密码
- [ ] `POST /console/api/account/integrates` - 集成设置
- [ ] `POST /console/api/account/delete/verify` - 删除账户校验
- [ ] `POST /console/api/account/delete` - 删除账户
- [ ] `POST /console/api/account/delete/feedback` - 删除账户反馈
- [ ] `POST /console/api/account/education/verify` - 教育邮箱校验
- [ ] `POST /console/api/account/education` - 教育信息
- [ ] `GET /console/api/account/education/autocomplete` - 教育信息自动补全

#### 模型与工具

- [ ] `GET /console/api/workspaces/current/model-providers` - 获取模型提供商列表
- [ ] `GET /console/api/workspaces/current/model-providers/{provider}` - 获取指定模型提供商
- [ ] `POST /console/api/workspaces/current/model-providers/{provider}/credentials` - 设置模型凭证
- [ ] `POST /console/api/workspaces/current/model-providers/{provider}/credentials/validate` - 校验模型凭证
- [ ] `GET /console/api/workspaces/current/model-providers/{provider}/models` - 获取模型列表
- [ ] `GET /console/api/workspaces/current/models/model-types/{model_type}` - 获取指定类型模型
- [ ] `GET /console/api/workspaces/current/default-model` - 获取默认模型
- [ ] `GET /console/api/workspaces/current/tool-providers` - 获取工具提供商列表
- [ ] `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/tools` - 获取内置工具列表
- [ ] `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/info` - 获取内置工具信息
- [ ] `DELETE /console/api/workspaces/current/tool-provider/builtin/{provider}/delete` - 删除内置工具
- [ ] `POST /console/api/workspaces/current/tool-provider/builtin/{provider}/update` - 更新内置工具
- [ ] `GET /console/api/workspaces/current/tool-provider/builtin/{provider}/icon` - 获取内置工具图标
- [ ] `POST /console/api/workspaces/current/tool-provider/api/add` - 添加 API 工具
- [ ] `GET /console/api/workspaces/current/tool-provider/api/remote` - 获取远程 API 工具 schema
- [ ] `GET /console/api/workspaces/current/tool-provider/api/tools` - 获取 API 工具列表
- [ ] `POST /console/api/workspaces/current/tool-provider/api/update` - 更新 API 工具
- [ ] `DELETE /console/api/workspaces/current/tool-provider/api/delete` - 删除 API 工具
- [ ] `GET /console/api/workspaces/current/tool-provider/api/get` - 获取 API 工具
- [ ] `GET /console/api/workspaces/current/tool-provider/api/schema` - 获取 API 工具 schema
- [ ] `POST /console/api/workspaces/current/tool-provider/api/test/pre` - 预测试 API 工具
- [ ] `POST /console/api/workspaces/current/tool-provider/workflow/create` - 创建工作流工具
- [ ] `POST /console/api/workspaces/current/tool-provider/workflow/update` - 更新工作流工具
- [ ] `DELETE /console/api/workspaces/current/tool-provider/workflow/delete` - 删除工作流工具
- [ ] `GET /console/api/workspaces/current/tool-provider/workflow/get` - 获取工作流工具

#### 端点管理

- [ ] `POST /console/api/workspaces/current/endpoints/create` - 创建端点
- [ ] `GET /console/api/workspaces/current/endpoints/list` - 获取端点列表
- [ ] `GET /console/api/workspaces/current/endpoints/list/plugin` - 获取插件端点列表
- [ ] `DELETE /console/api/workspaces/current/endpoints/delete` - 删除端点
- [ ] `POST /console/api/workspaces/current/endpoints/update` - 更新端点
- [ ] `POST /console/api/workspaces/current/endpoints/enable` - 启用端点
- [ ] `POST /console/api/workspaces/current/endpoints/disable` - 禁用端点

#### 扩展管理

- [ ] `POST /console/api/code-based-extension` - 代码扩展
- [ ] `POST /console/api/api-based-extension` - API 扩展
- [ ] `GET /console/api/api-based-extension/{id}` - 获取 API 扩展详情

#### 标签管理

- [ ] `GET /console/api/tags` - 获取标签列表
- [ ] `PUT /console/api/tags/{tag_id}` - 更新/删除标签
- [ ] `POST /console/api/tag-bindings/create` - 创建标签绑定
- [ ] `POST /console/api/tag-bindings/remove` - 移除标签绑定

#### 合规与账单

- [ ] `GET /console/api/compliance/download` - 下载合规文件
- [ ] `GET /console/api/billing/subscription` - 获取订阅信息
- [ ] `GET /console/api/billing/invoices` - 获取发票信息

#### 应用导入与工作流

- [ ] `POST /console/api/apps/imports` - 导入应用
- [ ] `POST /console/api/apps/imports/{import_id}/confirm` - 确认导入
- [ ] `POST /console/api/apps/imports/{app_id}/check-dependencies` - 检查依赖
- [ ] `GET /console/api/apps/{app_id}/server` - 获取 MCP server 信息
- [ ] `POST /console/api/apps/{server_id}/server/refresh` - 刷新 MCP server
- [ ] `GET /console/api/apps/{app_id}/workflow-app-logs` - 获取工作流日志

#### 其它补充

- [ ] `POST /console/api/test/retrieval` - Bedrock 检索测试 

#### 会话管理（Conversations）

- [ ] `GET /console/api/apps/{app_id}/completion-conversations` - 获取补全会话列表
- [ ] `GET /console/api/apps/{app_id}/completion-conversations/{conversation_id}` - 获取补全会话详情
- [ ] `GET /console/api/apps/{app_id}/chat-conversations` - 获取聊天会话列表
- [ ] `GET /console/api/apps/{app_id}/chat-conversations/{conversation_id}` - 获取聊天会话详情
- [ ] `GET /console/api/apps/{app_id}/conversation-variables` - 获取会话变量 