package request

// UserCreation struct.
type UserCreation struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PhotoURL    string `json:"photo_url"`
	PersonnelID string `json:"personnel_id"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
}
