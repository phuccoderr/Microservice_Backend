package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Message []string `json:"message"`
	Error   string   `json:"error"`
	Status  int      `json:"statusCode"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, message, error string, status int) {
	c.AbortWithStatusJSON(status, ResponseError{
		Message: []string{message},
		Error:   error,
		Status:  status,
	})
}
