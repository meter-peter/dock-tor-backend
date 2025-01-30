package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetByID(id uint) (*models.Patient, error)
	GetByAMKA(amka string) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
	List() ([]models.Patient, error)
}

type patientRepository struct{}

func NewPatientRepository() PatientRepository {
	return &patientRepository{}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return config.DB.Create(patient).Error
}

func (r *patientRepository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetByAMKA(amka string) (*models.Patient, error) {
	var patient models.Patient
	if err := config.DB.Where("amka = ?", amka).First(&patient).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return config.DB.Save(patient).Error
}

func (r *patientRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Patient{}, id).Error
}

func (r *patientRepository) List() ([]models.Patient, error) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
