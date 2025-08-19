package models

type SetModelProvidersRequest struct {
	Model         string                        `json:"model"`
	ModelType     string                        `json:"model_type"`
	Credentials   SetModelProviderCredentials   `json:"credentials"`
	LoadBalancing SetModelProviderLoadBalancing `json:"load_balancing"`
}

type SetModelProviderCredentials struct {
	Mode                    string `json:"mode"`
	ContextSize             string `json:"context_size"`
	MaxTokensToSample       string `json:"max_tokens_to_sample"`
	AgentThoughSupport      string `json:"agent_though_support"`
	FunctionCallingType     string `json:"function_calling_type"`
	StreamFunctionCalling   string `json:"stream_function_calling"`
	VisionSupport           string `json:"vision_support"`
	StructuredOutputSupport string `json:"structured_output_support"`
	StreamModeAuth          string `json:"stream_mode_auth"`
	StreamModeDelimiter     string `json:"stream_mode_delimiter"`
	Voices                  string `json:"voices"`
	ApiKey                  string `json:"api_key"`
	EndpointUrl             string `json:"endpoint_url"`
}

type SetModelProviderLoadBalancing struct {
	Enabled bool     `json:"enabled"`
	Configs []string `json:"configs"`
}
