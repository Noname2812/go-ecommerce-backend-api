package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusCodeMap = map[int]int{
	200: http.StatusOK,
	400: http.StatusBadRequest,
	401: http.StatusUnauthorized,
	404: http.StatusNotFound,
	408: http.StatusRequestTimeout,
	409: http.StatusConflict,
	422: http.StatusUnprocessableEntity,
	429: http.StatusTooManyRequests,
	500: http.StatusInternalServerError,
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type ErrorResponseData struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	status, exists := statusCodeMap[code/100]
	if !exists {
		status = http.StatusOK
	}
	c.JSON(status, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int, message string, errors map[string]string) {
	status, exists := statusCodeMap[code/100]
	if !exists {
		status = http.StatusInternalServerError
	}
	mess, ok := msg[code]
	if message == "" || (ok && mess != "") {
		message = msg[code]
	}
	c.JSON(status, ErrorResponseData{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
