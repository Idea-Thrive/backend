package request

// Login struct.
type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
