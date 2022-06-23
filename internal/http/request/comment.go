package request

type CommentCreation struct {
	UserID      string `json:"user_id"`
	CompanyID   string `json:"company_id"`
	IdeaID      string `json:"idea_id"`
	Score       int    `json:"score"`
	Description string `json:"description"`
}
