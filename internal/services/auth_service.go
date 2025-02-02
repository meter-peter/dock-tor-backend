package services

import (
	"errors"

	"clinic-management/internal/models"
	"clinic-management/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (a *AuthService) RegisterUser(req models.UserRegistration) (*models.User, error) {
	existingUser, _ := a.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}

	err = a.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
