package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryWithLogger middleware catches all panics and logs the error with stack trace.
func RecoveryWithLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				// write stack trace
				var buf bytes.Buffer
				fmt.Fprintf(&buf, "PANIC RECOVERED: %v\n", rec)
				buf.Write(debug.Stack())

				logger.Error("Recovered from panic",
					zap.Any("error", rec),
					zap.String("stack", buf.String()),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				// response error
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"success": false,
					"code":    500,
					"message": "Internal Server Error",
					"errors":  map[string]string{"error": "Internal Server Error"},
					"data":    nil,
				})
			}
		}()

		// continue
		c.Next()
	}
}
