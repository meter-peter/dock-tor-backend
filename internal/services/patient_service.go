// internal/services/patient_service.go

package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type PatientService interface {
	CreatePatient(patient *models.Patient) error
	GetPatient(id uint) (*models.Patient, error)
	UpdatePatient(patient *models.Patient) (*models.Patient, error)
	DeletePatient(id uint) error
	ListPatients() ([]models.Patient, error)
	BulkCreatePatients(reader io.Reader) (int, []error)
}

type patientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patient *models.Patient) error {
	patient.Registration = time.Now()
	_, err := s.repo.GetByAMKA(patient.AMKA)
	if err == nil {
		return errors.New("patient with this AMKA already exists")
	}
	return s.repo.Create(patient)
}

func (s *patientService) GetPatient(id uint) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *patientService) UpdatePatient(patient *models.Patient) (*models.Patient, error) {
	existing, err := s.repo.GetByID(patient.ID)
	if err != nil {
		return nil, err
	}
	if existing.AMKA != patient.AMKA {
		return nil, errors.New("AMKA cannot be modified")
	}
	err = s.repo.Update(patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}
func (s *patientService) DeletePatient(id uint) error {
	return s.repo.Delete(id)
}

func (s *patientService) ListPatients() ([]models.Patient, error) {
	return s.repo.List()
}

func (s *patientService) BulkCreatePatients(reader io.Reader) (int, []error) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return 0, []error{err}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(records))
	sem := make(chan struct{}, 10) // Limit concurrent goroutines
	var successCount int32

	for _, record := range records[1:] { // Skip header
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
