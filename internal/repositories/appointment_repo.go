// internal/repositories/appointment_repo.go

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

func (r *AppointmentRepository) GetByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.First(&appointment, id).Error
	return &appointment, err
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Appointment{}, id).Error
}

func (r *AppointmentRepository) List(startDate, endDate time.Time, patientID, doctorID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	query := r.db.Where("date_time BETWEEN ? AND ?", startDate, endDate)

	if patientID != 0 {
		query = query.Where("patient_id = ?", patientID)
	}
	if doctorID != 0 {
		query = query.Where("doctor_id = ?", doctorID)
	}

	err := query.Find(&appointments).Error
	return appointments, err
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
