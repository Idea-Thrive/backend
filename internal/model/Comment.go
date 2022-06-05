package model

import "time"

// Comment struct.
type Comment struct {
	Score       int       `json:"score,omitempty"`
	Description string    `json:"description,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	CompanyID   int       `json:"company_id,omitempty"`
	IdeaID      int       `json:"idea_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
