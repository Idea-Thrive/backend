package model

// Idea struct.
type Idea struct {
	ID            string  `json:"id,omitempty"`
	CategoryID    string  `json:"category_id,omitempty"`
	CategoryName  string  `json:"category_name"`
	CategoryColor string  `json:"category_color"`
	Title         string  `json:"title,omitempty"`
	Description   string  `json:"description,omitempty"`
	Score         float64 `json:"score"`
	CreatorID     string  `json:"creator_id,omitempty"`
	CompanyID     string  `json:"company_id,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
