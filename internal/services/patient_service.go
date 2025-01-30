package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type PatientService interface {
	CreatePatient(patient *models.Patient) error
	GetPatient(id uint) (*models.Patient, error)
	UpdatePatient(patient *models.Patient) error
	DeletePatient(id uint) error
	ListPatients() ([]models.Patient, error)
	BulkCreatePatients(records [][]string) (int, []error)
}

type patientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patient *models.Patient) error {
	// Set registration time
	patient.Registration = time.Now()

	// Check AMKA uniqueness
	_, err := s.repo.GetByAMKA(patient.AMKA)
	if err == nil {
		return errors.New("patient with this AMKA already exists")
	}

	return s.repo.Create(patient)
}

func (s *patientService) GetPatient(id uint) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *patientService) UpdatePatient(patient *models.Patient) error {
	existing, err := s.repo.GetByID(patient.ID)
	if err != nil {
		return err
	}

	// Prevent AMKA modification
	if existing.AMKA != patient.AMKA {
		return errors.New("AMKA cannot be modified")
	}

	return s.repo.Update(patient)
}

func (s *patientService) DeletePatient(id uint) error {
	return s.repo.Delete(id)
}

func (s *patientService) ListPatients() ([]models.Patient, error) {
	return s.repo.List()
}
func (s *patientService) BulkCreatePatients(records [][]string) (int, []error) {
	var wg sync.WaitGroup
	errChan := make(chan error, len(records))
	sem := make(chan struct{}, 10) // Limit concurrent goroutines
	var successCount int32

	for _, record := range records {
		wg.Add(1)
		go func(r []string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			if len(r) < 3 {
				errChan <- errors.New("invalid record format")
				return
			}

			userID, err := strconv.ParseUint(r[1], 10, 64)
			if err != nil {
				errChan <- errors.New("invalid user ID format")
				return
			}

			patient := models.Patient{
				AMKA:         r[0],
				UserID:       uint(userID),
				Registration: time.Now(),
			}

			if err := s.CreatePatient(&patient); err != nil {
				errChan <- err
			} else {
				atomic.AddInt32(&successCount, 1)
			}
		}(record)
	}

	wg.Wait()
	close(errChan)
	close(sem)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	return int(successCount), errors
}
