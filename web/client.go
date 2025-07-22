package web

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// Client Web API客户端
type Client struct {
	baseClient  *client.BaseClient
	appCode     string
	accessToken string
}

// NewClient 创建Web API客户端 (需要app code而不是api key)
func NewClient(appCode, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/api",
		AuthType:   client.AuthTypeBearer,
		Token:      "", // 初始为空，通过passport获取
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
		appCode:    appCode,
	}
}

// PassportResponse Passport响应
type PassportResponse struct {
	AccessToken string `json:"access_token"`
}

// GetPassport 获取访问令牌
func (c *Client) GetPassport(ctx context.Context, userID string) error {
	req := &client.Request{
		Method: "GET",
		Path:   "/passport",
		Headers: map[string]string{
			"X-App-Code": c.appCode,
		},
		Query: map[string]string{},
	}

	if userID != "" {
		req.Query["user_id"] = userID
	}

	var result PassportResponse
	if err := c.baseClient.DoJSON(ctx, req, &result); err != nil {
		return fmt.Errorf("failed to get passport: %w", err)
	}

	c.accessToken = result.AccessToken
	c.baseClient.WithToken(result.AccessToken)
	return nil
}

// ensureAuthenticated 确保已认证
func (c *Client) ensureAuthenticated(ctx context.Context) error {
	if c.accessToken == "" {
		return c.GetPassport(ctx, "")
	}
	return nil
}

// doAuthenticatedRequest 执行认证请求
func (c *Client) doAuthenticatedRequest(ctx context.Context, req *client.Request, result interface{}) error {
	if err := c.ensureAuthenticated(ctx); err != nil {
		return err
	}

	if req.Headers == nil {
		req.Headers = make(map[string]string)
	}
	req.Headers["X-App-Code"] = c.appCode

	return c.baseClient.DoJSON(ctx, req, result)
}

// GetAppParameters 获取应用参数
func (c *Client) GetAppParameters(ctx context.Context) (*models.AppParameters, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/parameters",
	}

	var result models.AppParameters
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// GetAppMeta 获取应用元数据
func (c *Client) GetAppMeta(ctx context.Context) (*models.AppMeta, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/meta",
	}

	var result models.AppMeta
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// GetWebAppAccessMode 检查Web应用访问模式
func (c *Client) GetWebAppAccessMode(ctx context.Context, appID string, appCode string) (*models.WebAppAccessMode, error) {
	query := make(map[string]string)
	if appID != "" {
		query["appId"] = appID
	}
	if appCode != "" {
		query["appCode"] = appCode
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/webapp/access-mode",
		Query:  query,
	}

	var result models.WebAppAccessMode
	// 这个端点不需要认证
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CheckWebAppPermission 检查Web应用权限
func (c *Client) CheckWebAppPermission(ctx context.Context, appID string) (*models.WebAppPermission, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/webapp/permission",
		Query:  map[string]string{"appId": appID},
	}

	var result models.WebAppPermission
	// 这个端点不需要认证
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CompletionRequest 文本补全请求
type CompletionRequest struct {
	Inputs        map[string]interface{}   `json:"inputs"`
	Query         string                   `json:"query,omitempty"`
	Files         []map[string]interface{} `json:"files,omitempty"`
	ResponseMode  models.ResponseMode      `json:"response_mode,omitempty"`
	RetrieverFrom string                   `json:"retriever_from,omitempty"`
}

// Completion 文本补全
func (c *Client) Completion(ctx context.Context, req *CompletionRequest) (*models.GenerateResponse, error) {
	if req.ResponseMode == "" {
		req.ResponseMode = models.ResponseModeBlocking
	}
	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "web_app"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages",
		Body:   req,
	}

	var result models.GenerateResponse
	err := c.doAuthenticatedRequest(ctx, httpReq, &result)
	return &result, err
}

// CompletionStop 停止文本补全
func (c *Client) CompletionStop(ctx context.Context, taskID string) error {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages/" + taskID + "/stop",
	}

	var result map[string]string
	return c.doAuthenticatedRequest(ctx, httpReq, &result)
}

// ChatRequest 聊天请求
type ChatRequest struct {
	Inputs          map[string]interface{}   `json:"inputs"`
	Query           string                   `json:"query"`
	Files           []map[string]interface{} `json:"files,omitempty"`
	ResponseMode    models.ResponseMode      `json:"response_mode,omitempty"`
	ConversationID  string                   `json:"conversation_id,omitempty"`
	ParentMessageID string                   `json:"parent_message_id,omitempty"`
	RetrieverFrom   string                   `json:"retriever_from,omitempty"`
}

// Chat 聊天对话
func (c *Client) Chat(ctx context.Context, req *ChatRequest) (*models.GenerateResponse, error) {
	if req.ResponseMode == "" {
		req.ResponseMode = models.ResponseModeBlocking
	}
	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "web_app"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages",
		Body:   req,
	}

	var result models.GenerateResponse
	err := c.doAuthenticatedRequest(ctx, httpReq, &result)
	return &result, err
}

// ChatStop 停止聊天
func (c *Client) ChatStop(ctx context.Context, taskID string) error {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages/" + taskID + "/stop",
	}

	var result map[string]string
	return c.doAuthenticatedRequest(ctx, httpReq, &result)
}

// ChatStream 流式聊天对话
func (c *Client) ChatStream(ctx context.Context, req *ChatRequest, handler client.SSEHandler) error {
	// 强制设置为流式模式
	req.ResponseMode = models.ResponseModeStreaming

	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "web_app"
	}

	// 确保已认证
	if err := c.ensureAuthenticated(ctx); err != nil {
		return err
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages",
		Body:   req,
		Headers: map[string]string{
			"X-App-Code": c.appCode,
		},
	}

	return c.baseClient.StreamResponse(ctx, httpReq, handler)
}

// CompletionStream 流式文本补全
func (c *Client) CompletionStream(ctx context.Context, req *CompletionRequest, handler client.SSEHandler) error {
	// 强制设置为流式模式
	req.ResponseMode = models.ResponseModeStreaming

	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "web_app"
	}

	// 确保已认证
	if err := c.ensureAuthenticated(ctx); err != nil {
		return err
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages",
		Body:   req,
		Headers: map[string]string{
			"X-App-Code": c.appCode,
		},
	}

	return c.baseClient.StreamResponse(ctx, httpReq, handler)
}

// GetConversations 获取对话列表
func (c *Client) GetConversations(ctx context.Context, lastID string, limit int, pinned *bool, sortBy string) (*models.ConversationListResponse, error) {
	query := make(map[string]string)
	if lastID != "" {
		query["last_id"] = lastID
	}
	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}
	if pinned != nil {
		query["pinned"] = strconv.FormatBool(*pinned)
	}
	if sortBy != "" {
		query["sort_by"] = sortBy
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/conversations",
		Query:  query,
	}

	var result models.ConversationListResponse
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// DeleteConversation 删除对话
func (c *Client) DeleteConversation(ctx context.Context, conversationID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/conversations/" + conversationID,
	}

	var result models.ConversationOperationResponse
	return c.doAuthenticatedRequest(ctx, req, &result)
}

// RenameConversation 重命名对话
func (c *Client) RenameConversation(ctx context.Context, conversationID string, request *models.ConversationRenameRequest) (*models.SimpleConversation, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/conversations/" + conversationID + "/name",
		Body:   request,
	}

	var result models.SimpleConversation
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// PinConversation 置顶对话
func (c *Client) PinConversation(ctx context.Context, conversationID string) error {
	req := &client.Request{
		Method: "PATCH",
		Path:   "/conversations/" + conversationID + "/pin",
	}

	var result models.ConversationOperationResponse
	return c.doAuthenticatedRequest(ctx, req, &result)
}

// UnpinConversation 取消置顶对话
func (c *Client) UnpinConversation(ctx context.Context, conversationID string) error {
	req := &client.Request{
		Method: "PATCH",
		Path:   "/conversations/" + conversationID + "/unpin",
	}

	var result models.ConversationOperationResponse
	return c.doAuthenticatedRequest(ctx, req, &result)
}

// GetMessages 获取消息列表
func (c *Client) GetMessages(ctx context.Context, conversationID string, firstID string, limit int) (*models.MessageListResponse, error) {
	query := make(map[string]string)
	if firstID != "" {
		query["first_id"] = firstID
	}
	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/messages",
		Query:  query,
	}

	var result models.MessageListResponse
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// SendMessageFeedback 发送消息反馈
func (c *Client) SendMessageFeedback(ctx context.Context, messageID string, feedback *models.MessageFeedbackRequest) error {
	req := &client.Request{
		Method: "POST",
		Path:   "/messages/" + messageID + "/feedbacks",
		Body:   feedback,
	}

	var result map[string]string
	return c.doAuthenticatedRequest(ctx, req, &result)
}

// GetMessageMoreLikeThis 获取类似消息
func (c *Client) GetMessageMoreLikeThis(ctx context.Context, messageID string, responseMode models.ResponseMode) (*models.GenerateResponse, error) {
	query := map[string]string{
		"response_mode": string(responseMode),
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/messages/" + messageID + "/more-like-this",
		Query:  query,
	}

	var result models.GenerateResponse
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// GetSuggestedQuestions 获取建议问题
func (c *Client) GetSuggestedQuestions(ctx context.Context, messageID string) (*models.SuggestedQuestionsResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/messages/" + messageID + "/suggested-questions",
	}

	var result models.SuggestedQuestionsResponse
	err := c.doAuthenticatedRequest(ctx, req, &result)
	return &result, err
}

// UploadFile 上传文件
func (c *Client) UploadFile(ctx context.Context, filename string, fileData []byte, source string) (*models.FileUpload, error) {
	// 确保已认证
	if err := c.ensureAuthenticated(ctx); err != nil {
		return nil, err
	}

	extraFields := make(map[string]string)
	if source != "" {
		extraFields["source"] = source
	}

	_, err := c.baseClient.UploadFile(ctx, "/files", "file", filename, fileData, extraFields)
	if err != nil {
		return nil, err
	}

	var result models.FileUpload
	if err := c.baseClient.DoJSON(ctx, &client.Request{}, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AudioToText 语音转文字
func (c *Client) AudioToText(ctx context.Context, audioData []byte, filename string) (map[string]interface{}, error) {
	// 确保已认证
	if err := c.ensureAuthenticated(ctx); err != nil {
		return nil, err
	}

	_, err := c.baseClient.UploadFile(ctx, "/audio-to-text", "file", filename, audioData, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := c.baseClient.DoJSON(ctx, &client.Request{}, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TextToAudioRequest 文字转语音请求
type TextToAudioRequest struct {
	MessageID string `json:"message_id,omitempty"`
	Voice     string `json:"voice,omitempty"`
	Text      string `json:"text,omitempty"`
	Streaming bool   `json:"streaming,omitempty"`
}

// TextToAudio 文字转语音
func (c *Client) TextToAudio(ctx context.Context, req *TextToAudioRequest) (map[string]interface{}, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/text-to-audio",
		Body:   req,
	}

	var result map[string]interface{}
	err := c.doAuthenticatedRequest(ctx, httpReq, &result)
	return result, err
}

// ============ 工作流相关 ============

// RunWorkflow 运行工作流
func (c *Client) RunWorkflow(ctx context.Context, req *models.WorkflowRunRequest) (*models.WorkflowRunResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/workflows/run",
		Body:   req,
	}

	var result models.WorkflowRunResponse
	err := c.doAuthenticatedRequest(ctx, httpReq, &result)
	return &result, err
}

// StopWorkflowTask 停止工作流任务
func (c *Client) StopWorkflowTask(ctx context.Context, taskID string) (*models.WorkflowStopResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/workflows/tasks/" + taskID + "/stop",
	}

	var result models.WorkflowStopResponse
	err := c.doAuthenticatedRequest(ctx, httpReq, &result)
	return &result, err
}
