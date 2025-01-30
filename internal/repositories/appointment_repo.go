package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"
	"time"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{db: config.DB}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *AppointmentRepository) HasPatientConflict(patientID uint, datetime time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&models.Appointment{}).
		Where("patient_id = ? AND date_time BETWEEN ? AND ?",
			patientID,
			datetime.Add(-30*time.Minute),
			datetime.Add(30*time.Minute),
		).
		Count(&count).
		Error

	return count > 0, err
}
