package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.DB}
}

func (r *UserRepository) GetUserByEmail(user *models.User, email string) error {
	return r.db.Where("email = ?", email).First(user).Error
}
