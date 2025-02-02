package handlers

import (
	"net/http"
	"time"

	"clinic-management/config"
	"clinic-management/internal/models"
	"clinic-management/internal/services"
	"clinic-management/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.UserRegistration true "User Registration Info"
// @Success 201 {object} models.User
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.UserRegistration

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, err := h.authService.RegisterUser(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, user)
}

// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginCredentials true "Login Credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginCredentials

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, err := h.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		utils.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	token, err := generateJWTToken(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, models.LoginResponse{Token: token})
}

// @Summary User logout
// @Description Logout the current user (invalidate the token)
// @Tags Authentication
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// In a stateless JWT system, we don't need to do anything server-side
	// The client should discard the token
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func generateJWTToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().JWTSecret))
}
