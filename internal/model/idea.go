package model

// Idea struct.
type Idea struct {
	IdeaID        string `json:"idea_id,omitempty"`
	Category      int    `json:"category,omitempty"`
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	UpVoteCount   int    `json:"up_vote_count,omitempty"`
	DownVoteCount int    `json:"down_vote_count,omitempty"`
	CreatorID     string `json:"creator_id,omitempty"`
	CompanyID     string `json:"company_id,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
