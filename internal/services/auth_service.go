package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (a *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	err := a.userRepo.GetUserByEmail(&user, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// TODO: Implement proper password hashing
	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	return &user, nil
}
