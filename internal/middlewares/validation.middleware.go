package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidatorMiddleware(v *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		// set the middleware
		c.Set("validation", v) // context
		c.Set("trace_id", uuid.New().String())

		c.Next()

	}
}
