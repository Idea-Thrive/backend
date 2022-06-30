package model

// User struct.
type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	PhotoURL    string `json:"photo_url"`
	CompanyID   string `json:"company_id"`
	PersonnelID string `json:"personnel_id"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
}
