package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponseData struct {
	Success bool              `json:"success"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Success: true,
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int, message string, errors map[string]string) {
	mess, ok := msg[code]
	if message == "" || (ok && mess != "") {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ErrorResponseData{
		Success: false,
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
