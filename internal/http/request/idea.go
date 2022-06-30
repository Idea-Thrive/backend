package request

type IdeaCreation struct {
	CategoryID  string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatorID   string `json:"creator_id"`
	CompanyID   string `json:"company_id"`
}
