package request

type CategoryCreation struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
