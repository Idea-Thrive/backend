package request

type IdeaCreation struct {
	Category    int    `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UpVote      int    `json:"up_vote"`
	DownVote    int    `json:"down_vote"`
	CreatorID   string `json:"creator_id"`
	CompanyID   string `json:"company_id"`
}
