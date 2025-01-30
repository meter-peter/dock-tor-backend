package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"clinic-management/internal/utils"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				utils.LogError(fmt.Sprintf("Panic recovered: %v\n%s", err, debug.Stack()))

				// Respond with error
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Internal server error",
				})
			}
		}()

		c.Next()
	}
}
