package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"
	"time"

	"gorm.io/gorm"
)

type DoctorAvailabilityRepository struct {
	db *gorm.DB
}

func NewDoctorAvailabilityRepository() *DoctorAvailabilityRepository {
	return &DoctorAvailabilityRepository{db: config.DB}
}

func (r *DoctorAvailabilityRepository) IsDoctorAvailable(doctorID uint, datetime time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&models.DoctorAvailability{}).
		Where("doctor_id = ? AND start_time <= ? AND end_time >= ? AND is_available = true",
			doctorID,
			datetime,
			datetime,
		).
		Count(&count).
		Error

	return count > 0, err
}
