package dto

type CustomerDto struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
