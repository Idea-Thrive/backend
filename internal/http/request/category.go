package request

// CategoryCreation struct.
type CategoryCreation struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}
