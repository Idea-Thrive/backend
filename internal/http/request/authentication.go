package request

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
