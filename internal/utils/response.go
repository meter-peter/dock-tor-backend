package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func UnauthorizedResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(401, APIResponse{
		Success: false,
		Error:   message,
	})
}

func ForbiddenResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(403, APIResponse{
		Success: false,
		Error:   message,
	})
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{
		Success: false,
		Error:   message,
	})
}
