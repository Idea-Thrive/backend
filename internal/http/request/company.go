package request

type CompanyCreation struct {
	CompanyID       string `json:"company_id"`
	Name            string `json:"name"`
	LogoURL         string `json:"logo_url"`
	OwnerNationalID string `json:"owner_national_id"`
	OwnerFirstName  string `json:"owner_first_name"`
	OwnerLastName   string `json:"owner_last_name"`
}