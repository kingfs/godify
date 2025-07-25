package console

// test pr
import (
	"context"
	"errors"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

func (c *Client) GetWorkspaces(ctx context.Context) (*models.WorkspacesApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces",
	}
	var resp models.WorkspacesApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetWorkspacesCurrent(ctx context.Context) (*models.Workspace, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current",
	}
	var resp models.Workspace
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetWorkspacesCurrentMembers(ctx context.Context) (*models.WorkspacesCurrentMembersApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/members",
	}
	var resp models.WorkspacesCurrentMembersApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 邀请成员邮件接口
func (c *Client) CreateWorkspacesCurrentMembersInviteEmail(ctx context.Context, emails []string, role, language string) (*models.WorkspaceInviteEmailApiResponse, error) {
	role_types := []string{"normal", "editor", "admin"}
	// 如果role不在role_types里，报错
	valid := false
	for _, r := range role_types {
		if role == r {
			valid = true
			break
		}
	}
	if !valid {
		err := errors.New("参数错误：role 必须为 normal、editor 或 admin")
		return nil, err
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/members/invite-email",
		Body: map[string]interface{}{
			"emails":   emails,
			"role":     role,
			"language": language,
		},
	}
	var resp models.WorkspaceInviteEmailApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 取消成员邀请/移除成员接口
func (c *Client) DeleteWorkspacesCurrentMembers(ctx context.Context, memberID string) (*models.WorkspaceOperationResponse, error) {
	path := "/workspaces/current/members/" + memberID
	req := &client.Request{
		Method: "DELETE",
		Path:   path,
	}
	var resp models.WorkspaceOperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 更新成员角色接口
func (c *Client) UpdateWorkspacesCurrentMembersRole(ctx context.Context, memberID, role string) (*models.WorkspaceUpdateRoleResponse, error) {
	path := "/workspaces/current/members/" + memberID + "/update-role"
	req := &client.Request{
		Method: "PUT",
		Path:   path,
		Body: map[string]interface{}{
			"role": role,
		},
	}
	var resp models.WorkspaceUpdateRoleResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取数据集操作员成员列表接口
func (c *Client) GetWorkspacesCurrentDatasetOperators(ctx context.Context) (*models.WorkspacesCurrentDatasetOperatorsApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/dataset-operators",
	}
	var resp models.WorkspacesCurrentDatasetOperatorsApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

