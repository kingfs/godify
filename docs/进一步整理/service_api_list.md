#### 索引与首页

- [ ] `GET /service-api/` - 首页/健康检查

#### 应用相关

- [ ] `GET /service-api/parameters` - 获取应用参数
- [ ] `GET /service-api/meta` - 获取应用元信息
- [ ] `GET /service-api/info` - 获取应用信息

#### 文件与音频

- [ ] `POST /service-api/files/upload` - 上传文件
- [ ] `POST /service-api/audio-to-text` - 音频转文本
- [ ] `POST /service-api/text-to-audio` - 文本转音频

#### 消息与反馈

- [ ] `GET /service-api/messages` - 获取消息列表
- [ ] `POST /service-api/messages/{message_id}/feedbacks` - 消息反馈
- [ ] `GET /service-api/messages/{message_id}/suggested` - 获取建议问题
- [ ] `GET /service-api/app/feedbacks` - 获取应用反馈

#### 会话管理

- [ ] `GET /service-api/conversations` - 获取会话列表
- [ ] `GET /service-api/conversations/{c_id}` - 获取会话详情
- [ ] `POST /service-api/conversations/{c_id}/name` - 修改会话名称
- [ ] `GET /service-api/conversations/{c_id}/variables` - 获取会话变量

#### 补全与聊天

- [ ] `POST /service-api/completion-messages` - 补全消息
- [ ] `POST /service-api/completion-messages/{task_id}/stop` - 停止补全
- [ ] `POST /service-api/chat-messages` - 聊天消息
- [ ] `POST /service-api/chat-messages/{task_id}/stop` - 停止聊天

#### 工作流

- [ ] `POST /service-api/workflows/run` - 运行工作流
- [ ] `GET /service-api/workflows/run/{workflow_run_id}` - 获取工作流运行详情
- [ ] `POST /service-api/workflows/tasks/{task_id}/stop` - 停止工作流任务
- [ ] `GET /service-api/workflows/logs` - 获取工作流日志

#### 站点

- [ ] `GET /service-api/site` - 获取站点信息

#### 数据集管理

- [ ] `GET /service-api/datasets` - 获取数据集列表
- [ ] `GET /service-api/datasets/{dataset_id}` - 获取数据集详情
- [ ] `POST /service-api/datasets/{dataset_id}/hit-testing` - 数据集命中测试
- [ ] `POST /service-api/datasets/{dataset_id}/retrieve` - 数据集检索
- [ ] `POST /service-api/datasets/{dataset_id}/documents/{document_id}/upload-file` - 上传文档文件
- [ ] `GET /service-api/datasets/{dataset_id}/documents` - 获取文档列表
- [ ] `GET /service-api/datasets/{dataset_id}/documents/{document_id}` - 获取文档详情
- [ ] `DELETE /service-api/datasets/{dataset_id}/documents/{document_id}` - 删除文档
- [ ] `GET /service-api/datasets/{dataset_id}/documents/{batch}/indexing-status` - 获取文档索引状态
- [ ] `GET /service-api/datasets/{dataset_id}/documents/{document_id}/segments` - 获取文档分段
- [ ] `POST /service-api/datasets/{dataset_id}/documents/{document_id}/segments` - 添加文档分段
- [ ] `POST /service-api/datasets/{dataset_id}/documents/status/{action}` - 批量操作文档状态
- [ ] `GET /service-api/datasets/tags` - 获取数据集标签
- [ ] `POST /service-api/datasets/tags/binding` - 绑定数据集标签
- [ ] `POST /service-api/datasets/tags/unbinding` - 解绑数据集标签
- [ ] `GET /service-api/datasets/{dataset_id}/tags` - 获取数据集标签绑定状态
- [ ] `POST /service-api/datasets/{dataset_id}/metadata` - 创建元数据
- [ ] `GET /service-api/datasets/{dataset_id}/metadata/{metadata_id}` - 获取元数据详情
- [ ] `GET /service-api/datasets/metadata/built-in` - 获取内置元数据字段
- [ ] `POST /service-api/datasets/{dataset_id}/documents/metadata` - 获取文档元数据

#### 注解管理

- [ ] `POST /service-api/apps/annotation-reply/{action}` - 注解回复
- [ ] `GET /service-api/apps/annotation-reply/{action}/status/{job_id}` - 注解回复状态
- [ ] `GET /service-api/apps/annotations` - 获取注解列表
- [ ] `PUT /service-api/apps/annotations/{annotation_id}` - 更新注解

#### 模型

- [ ] `GET /service-api/workspaces/current/models/model-types/{model_type}` - 获取指定类型模型 