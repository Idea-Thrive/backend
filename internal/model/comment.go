package model

// Comment struct.
type Comment struct {
	Score       int    `json:"score,omitempty"`
	Description string `json:"description,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	CompanyID   int    `json:"company_id,omitempty"`
	IdeaID      int    `json:"idea_id,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
