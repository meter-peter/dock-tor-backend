// internal/repositories/patient_repo.go

package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetByID(id uint) (*models.Patient, error)
	GetByAMKA(amka string) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
	List() ([]models.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository() PatientRepository {
	return &patientRepository{db: config.DB}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	return &patient, err
}

func (r *patientRepository) GetByAMKA(amka string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Where("amka = ?", amka).First(&patient).Error
	return &patient, err
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}

func (r *patientRepository) List() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Find(&patients).Error
	return patients, err
}
