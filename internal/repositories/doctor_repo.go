package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"
)

type DoctorRepository struct{}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	return config.DB.Create(doctor).Error
}

func (r *DoctorRepository) GetByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := config.DB.Preload("User").First(&doctor, id).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	return config.DB.Save(doctor).Error
}

func (r *DoctorRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Doctor{}, id).Error
}

func (r *DoctorRepository) List() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := config.DB.Preload("User").Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) CreateAvailability(availability *models.DoctorAvailability) error {
	return config.DB.Create(availability).Error
}

func (r *DoctorRepository) GetAvailabilities(doctorID uint) ([]models.DoctorAvailability, error) {
	var availabilities []models.DoctorAvailability
	err := config.DB.Where("doctor_id = ?", doctorID).Find(&availabilities).Error
	return availabilities, err
}

func (r *DoctorRepository) DeleteAvailability(id uint) error {
	return config.DB.Delete(&models.DoctorAvailability{}, id).Error
}
