package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// GenerateRuleRequest is the request to generate a rule config.
type GenerateRuleRequest struct {
	Instruction string                 `json:"instruction"`
	ModelConfig map[string]interface{} `json:"model_config"`
	NoVariable  bool                   `json:"no_variable"`
}

// GenerateRuleResponse is the response from generating a rule config.
type GenerateRuleResponse struct {
	Prompt           string   `json:"prompt"`
	Variables        []string `json:"variables"`
	OpeningStatement string   `json:"opening_statement"`
	Error            string   `json:"error"`
}

// GenerateCodeRequest is the request to generate code.
type GenerateCodeRequest struct {
	Instruction  string                 `json:"instruction"`
	ModelConfig  map[string]interface{} `json:"model_config"`
	CodeLanguage string                 `json:"code_language,omitempty"`
}

// GenerateCodeResponse is the response from generating code.
type GenerateCodeResponse struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Error    string `json:"error"`
}

// GenerateStructuredOutputRequest is the request to generate structured output.
type GenerateStructuredOutputRequest struct {
	Instruction string                 `json:"instruction"`
	ModelConfig map[string]interface{} `json:"model_config"`
}

// GenerateStructuredOutputResponse is the response from generating structured output.
type GenerateStructuredOutputResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

// GenerateRule generates a rule config.
func (c *client.Client) GenerateRule(ctx context.Context, req *GenerateRuleRequest) (*GenerateRuleResponse, error) {
	var result GenerateRuleResponse
	err := c.sendRequest(ctx, "POST", "/console/api/rule-generate", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateCode generates code from an instruction.
func (c *client.Client) GenerateCode(ctx context.Context, req *GenerateCodeRequest) (*GenerateCodeResponse, error) {
	var result GenerateCodeResponse
	err := c.sendRequest(ctx, "POST", "/console/api/rule-code-generate", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateStructuredOutput generates a structured output (JSON schema) from an instruction.
func (c *client.Client) GenerateStructuredOutput(ctx context.Context, req *GenerateStructuredOutputRequest) (*GenerateStructuredOutputResponse, error) {
	var result GenerateStructuredOutputResponse
	err := c.sendRequest(ctx, "POST", "/console/api/rule-structured-output-generate", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
