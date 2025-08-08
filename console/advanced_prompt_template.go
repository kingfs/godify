package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// GetAdvancedPromptTemplate retrieves an advanced prompt template based on app and model configuration.
func (c *client.Client) GetAdvancedPromptTemplate(ctx context.Context, appMode, modelMode, modelName, hasContext string) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/app/prompt-templates?app_mode=%s&model_mode=%s&has_context=%s&model_name=%s",
		appMode, modelMode, hasContext, modelName)

	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
