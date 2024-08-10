package dto

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"statusCode"`
}

func BuildResponseError(message, error string, status int) *Error {
	return &Error{
		Message: message,
		Error:   error,
		Status:  status,
	}
}
