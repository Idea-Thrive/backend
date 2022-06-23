package model

// Comment struct.
type Comment struct {
	Score       int    `json:"score,omitempty"`
	Description string `json:"description,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	CompanyID   string `json:"company_id,omitempty"`
	IdeaID      string `json:"idea_id,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
