package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"errors"
	"time"
)

type DoctorService struct {
	repo *repositories.DoctorRepository
}

func NewDoctorService(repo *repositories.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) CreateDoctor(doctor *models.Doctor) error {
	return s.repo.Create(doctor)
}

func (s *DoctorService) GetDoctor(id uint) (*models.Doctor, error) {
	return s.repo.GetByID(id)
}

func (s *DoctorService) UpdateDoctor(doctor *models.Doctor) error {
	return s.repo.Update(doctor)
}

func (s *DoctorService) DeleteDoctor(id uint) error {
	return s.repo.Delete(id)
}

func (s *DoctorService) ListDoctors() ([]models.Doctor, error) {
	return s.repo.List()
}

func (s *DoctorService) CreateAvailability(doctorID uint, startTime, endTime time.Time) error {
	if startTime.After(endTime) {
		return errors.New("start time must be before end time")
	}

	availability := &models.DoctorAvailability{
		DoctorID:  doctorID,
		StartTime: startTime,
		EndTime:   endTime,
	}

	return s.repo.CreateAvailability(availability)
}

func (s *DoctorService) GetAvailabilities(doctorID uint) ([]models.DoctorAvailability, error) {
	return s.repo.GetAvailabilities(doctorID)
}

func (s *DoctorService) DeleteAvailability(id uint) error {
	return s.repo.DeleteAvailability(id)
}
