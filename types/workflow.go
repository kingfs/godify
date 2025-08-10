package types

// WorkflowRunRequest is the request to run a workflow.
type WorkflowRunRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	Files        []File                 `json:"files,omitempty"`
	ResponseMode string                 `json:"response_mode"`
	User         string                 `json:"user"`
}
