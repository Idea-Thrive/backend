package model

// Idea struct.
type Idea struct {
	Title         string `json:"title,omitempty"`
	Category      string `json:"category,omitempty"`
	Description   string `json:"description,omitempty"`
	UpVoteCount   int    `json:"up_vote_count,omitempty"`
	DownVoteCount int    `json:"down_vote_count,omitempty"`
	CreatorID     string `json:"creator_id,omitempty"`
	CompanyID     string `json:"company_id,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
