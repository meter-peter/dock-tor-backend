package middleware

import (
	"clinic-management/config"
	"clinic-management/internal/models"
	"clinic-management/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Missing authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.UnauthorizedResponse(c, "Invalid authorization header format")
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JWTSecret), nil
		})

		if err != nil || !token.Valid {
			utils.UnauthorizedResponse(c, "Invalid token")
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RoleAuth(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			utils.UnauthorizedResponse(c, "Missing role information")
			return
		}

		// Convert interface{} to string
		role, ok := userRole.(string)
		if !ok {
			utils.UnauthorizedResponse(c, "Invalid role format")
			return
		}

		// Check if user has any of the allowed roles
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		utils.ForbiddenResponse(c, "Insufficient permissions")
		c.Abort()
	}
}
