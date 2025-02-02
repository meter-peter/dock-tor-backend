// internal/services/services.go

package services

type Services struct {
	AuthService        *AuthService
	PatientService     PatientService
	AppointmentService *AppointmentService
	DoctorService      *DoctorService
	HistoryService     *HistoryService
}

func NewServices(
	authService *AuthService,
	patientService PatientService,
	appointmentService *AppointmentService,
	doctorService *DoctorService,
	historyService *HistoryService,
) *Services {
	return &Services{
		AuthService:        authService,
		PatientService:     patientService,
		AppointmentService: appointmentService,
		DoctorService:      doctorService,
		HistoryService:     historyService,
	}
}
