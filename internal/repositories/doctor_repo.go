// internal/repositories/doctor_repo.go

package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"

	"gorm.io/gorm"
)

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{db: config.DB}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *DoctorRepository) GetByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Preload("User").First(&doctor, id).Error
	return &doctor, err
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *DoctorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Doctor{}, id).Error
}

func (r *DoctorRepository) List() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Preload("User").Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) CreateAvailability(availability *models.DoctorAvailability) error {
	return r.db.Create(availability).Error
}

func (r *DoctorRepository) GetAvailabilities(doctorID uint) ([]models.DoctorAvailability, error) {
	var availabilities []models.DoctorAvailability
	err := r.db.Where("doctor_id = ?", doctorID).Find(&availabilities).Error
	return availabilities, err
}

func (r *DoctorRepository) DeleteAvailability(id uint) error {
	return r.db.Delete(&models.DoctorAvailability{}, id).Error
}
