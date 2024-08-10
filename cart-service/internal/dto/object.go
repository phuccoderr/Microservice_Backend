package dto

type ResponseObject struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func BuildResponseObject(message string, status int, data interface{}) *ResponseObject {
	return &ResponseObject{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
