package dto

type Feedback struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Message     string `json:"message"`
}
