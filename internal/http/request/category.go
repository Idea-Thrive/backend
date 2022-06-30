package request

type CategoryCreation struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}
