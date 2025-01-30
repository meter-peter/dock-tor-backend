package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"errors"
	"time"
)

type AppointmentService struct {
	repo             *repositories.AppointmentRepository
	availabilityRepo *repositories.DoctorAvailabilityRepository
}

func NewAppointmentService(
	repo *repositories.AppointmentRepository,
	availabilityRepo *repositories.DoctorAvailabilityRepository,
) *AppointmentService {
	return &AppointmentService{
		repo:             repo,
		availabilityRepo: availabilityRepo,
	}
}

func (s *AppointmentService) CreateAppointment(appointment *models.Appointment) error {
	// Validate appointment time
	if appointment.DateTime.Before(time.Now().Add(15 * time.Minute)) {
		return errors.New("appointment must be scheduled at least 15 minutes in advance")
	}

	// Check doctor availability
	isAvailable, err := s.availabilityRepo.IsDoctorAvailable(
		appointment.DoctorID,
		appointment.DateTime,
	)
	if err != nil {
		return err
	}
	if !isAvailable {
		return errors.New("doctor is not available at this time")
	}

	// Check patient has no overlapping appointments
	hasConflict, err := s.repo.HasPatientConflict(
		appointment.PatientID,
		appointment.DateTime,
	)
	if err != nil {
		return err
	}
	if hasConflict {
		return errors.New("patient has conflicting appointment")
	}

	return s.repo.Create(appointment)
}

// Add other service methods
