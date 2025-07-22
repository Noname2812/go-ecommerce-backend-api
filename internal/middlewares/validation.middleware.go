package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set the middleware
		c.Set("validation", validator.New()) // context
		c.Set("trace_id", uuid.New().String())

		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				if e.Type == gin.ErrorTypeBind {
					validationErrors := formatValidationError(e.Err)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"success": false,
						"code":    400,
						"message": "Validation failed",
						"errors":  validationErrors,
					})
					return
				}
			}
		}

	}
}

// formatValidationError: Format error binding
func formatValidationError(err error) map[string]string {
	result := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			fieldName := fieldErr.Field()
			var msg string

			switch fieldErr.Tag() {
			case "required":
				msg = "This field is required"
			case "email":
				msg = "Invalid email format"
			case "min":
				msg = fmt.Sprintf("Minimum length is %s", fieldErr.Param())
			case "eqfield":
				msg = fmt.Sprintf("Must match %s", fieldErr.Param())
			case "oneof":
				msg = fmt.Sprintf("Must be one of: %s", fieldErr.Param())
			case "e164":
				msg = "Invalid phone number format"
			default:
				msg = "Invalid value"
			}

			result[fieldName] = msg
		}
	} else {
		result["error"] = err.Error()
	}

	return result
}
