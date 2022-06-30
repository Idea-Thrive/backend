package model

// Category struct.
type Category struct {
	ID          string `json:"id,omitempty"`
	CompanyID   string `json:"company_id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
