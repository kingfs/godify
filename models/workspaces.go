package models

type SetModelProvidersRequest struct {
	Model       string                      `json:"model"`
	ModelType   string                      `json:"model_type"`
	Credentials SetModelProviderCredentials `json:"credentials"`
}

type SetModelProviderCredentials struct {
	AgentThoughSupport      string `json:"agent_though_support"`
	ApiKey                  string `json:"api_key"`
	ContextSize             string `json:"context_size"`
	DisplayName             string `json:"display_name"`
	EndpointModelName       string `json:"endpoint_model_name"`
	EndpointUrl             string `json:"endpoint_url"`
	FunctionCallingType     string `json:"function_calling_type"`
	MaxTokensToSample       string `json:"max_tokens_to_sample"`
	Mode                    string `json:"mode"`
	StreamFunctionCalling   string `json:"stream_function_calling"`
	StreamModeAuth          string `json:"stream_mode_auth"`
	StreamModeDelimiter     string `json:"stream_mode_delimiter"`
	StructuredOutputSupport string `json:"structured_output_support"`
	VisionSupport           string `json:"vision_support"`
	Voices                  string `json:"voices"`
	ModelName               string `json:"__model_name"`
	ModelType               string `json:"__model_type"`
}
