// internal/services/appointment_service.go

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
	if appointment.DateTime.Before(time.Now().Add(15 * time.Minute)) {
		return errors.New("appointment must be scheduled at least 15 minutes in advance")
	}

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

func (s *AppointmentService) GetAppointment(id uint) (*models.Appointment, error) {
	return s.repo.GetByID(id)
}

func (s *AppointmentService) UpdateAppointment(appointment *models.Appointment) error {
	return s.repo.Update(appointment)
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	return s.repo.Delete(id)
}

func (s *AppointmentService) ListAppointments(startDate, endDate time.Time, patientID, doctorID uint) ([]models.Appointment, error) {
	return s.repo.List(startDate, endDate, patientID, doctorID)
}
