package model

// Company struct.
type Company struct {
	CompanyID       string `json:"company_id,omitempty"`
	Name            string `json:"name,omitempty"`
	LogoURL         string `json:"logo_url,omitempty"`
	OwnerNationalID string `json:"owner_national_id,omitempty"`
	OwnerFirstName  string `json:"owner_first_name,omitempty"`
	OwnerLastName   string `json:"owner_last_name,omitempty"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
