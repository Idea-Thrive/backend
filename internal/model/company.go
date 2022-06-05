package model

import "time"

// Company struct.
type Company struct {
	Name            string    `json:"name,omitempty"`
	LogoURL         string    `json:"logo_url,omitempty"`
	CompanyID       string    `json:"company_id,omitempty"`
	OwnerNationalID string    `json:"owner_national_id,omitempty"`
	OwnerFirstName  string    `json:"owner_first_name,omitempty"`
	OwnerLastName   string    `json:"owner_last_name,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
