#### 系统功能

- [ ] `GET /web/api/system-features` - 获取系统功能

#### 登录与认证

- [ ] `POST /web/api/login` - 登录
- [ ] `POST /web/api/email-code-login` - 发送邮箱验证码
- [ ] `POST /web/api/email-code-login/validity` - 校验邮箱验证码
- [ ] `POST /web/api/forgot-password` - 忘记密码发送邮件
- [ ] `POST /web/api/forgot-password/validity` - 校验忘记密码验证码
- [ ] `POST /web/api/forgot-password/resets` - 重置密码
- [ ] `POST /web/api/passport` - 通行证登录

#### 应用相关

- [ ] `GET /web/api/parameters` - 获取应用参数
- [ ] `GET /web/api/meta` - 获取应用元信息
- [ ] `GET /web/api/webapp/access-mode` - 获取 WebApp 访问模式
- [ ] `GET /web/api/webapp/permission` - 获取 WebApp 权限

#### 文件与远程文件

- [ ] `POST /web/api/files/upload` - 上传文件
- [ ] `GET /web/api/remote-files/{url}` - 获取远程文件信息
- [ ] `POST /web/api/remote-files/upload` - 上传远程文件

#### 消息与对话

- [ ] `GET /web/api/messages` - 获取消息列表
- [ ] `POST /web/api/messages/{message_id}/feedbacks` - 消息反馈
- [ ] `GET /web/api/messages/{message_id}/more-like-this` - 获取更多类似消息
- [ ] `GET /web/api/messages/{message_id}/suggested-questions` - 获取建议问题

#### 会话管理

- [ ] `GET /web/api/conversations` - 获取会话列表
- [ ] `GET /web/api/conversations/{c_id}` - 获取会话详情
- [ ] `POST /web/api/conversations/{c_id}/name` - 修改会话名称
- [ ] `POST /web/api/conversations/{c_id}/pin` - 置顶会话
- [ ] `POST /web/api/conversations/{c_id}/unpin` - 取消置顶

#### 补全与聊天

- [ ] `POST /web/api/completion-messages` - 补全消息
- [ ] `POST /web/api/completion-messages/{task_id}/stop` - 停止补全
- [ ] `POST /web/api/chat-messages` - 聊天消息
- [ ] `POST /web/api/chat-messages/{task_id}/stop` - 停止聊天

#### 音频与文本

- [ ] `POST /web/api/audio-to-text` - 音频转文本
- [ ] `POST /web/api/text-to-audio` - 文本转音频

#### 工作流

- [ ] `POST /web/api/workflows/run` - 运行工作流
- [ ] `POST /web/api/workflows/tasks/{task_id}/stop` - 停止工作流任务

#### 站点

- [ ] `GET /web/api/site` - 获取站点信息

#### 收藏消息

- [ ] `GET /web/api/saved-messages` - 获取收藏消息列表
- [ ] `GET /web/api/saved-messages/{message_id}` - 获取收藏消息详情 