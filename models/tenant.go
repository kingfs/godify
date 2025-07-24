package models

type Tenant struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Plan           string   `json:"plan"`
	Status         string   `json:"status"`
	CreatedAt      UnixTime `json:"created_at"`
	TrialEndReason *string  `json:"trial_end_reason,omitempty"`
	Role           string   `json:"role"`
}
