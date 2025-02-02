package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// SwaggerGormModel is a wrapper for gorm.Model that includes Swagger annotations
type SwaggerGormModel struct {
	// @Description The unique identifier for the record
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	// @Description The time when the record was created
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	// @Description The time when the record was last updated
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	// @Description The time when the record was deleted (soft delete)
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" example:"2023-01-01T00:00:00Z"`
}

// User represents a user in the system
// @Description User information
type User struct {
	SwaggerGormModel
	// @Description The user's email address
	Email string `json:"email" gorm:"uniqueIndex;not null" example:"georgios.papadopoulos@example.com"`
	// @Description The user's password (hashed)
	Password string `json:"-" gorm:"not null"`
	// @Description The user's first name
	FirstName string `json:"first_name" gorm:"not null" example:"Γεώργιος"`
	// @Description The user's last name
	LastName string `json:"last_name" gorm:"not null" example:"Παπαδόπουλος"`
	// @Description The user's role in the system
	Role string `json:"role" gorm:"not null" example:"patient"`
}

// UserRegistration represents the data needed to register a new user
// @Description User registration information
type UserRegistration struct {
	// @Description The user's email address
	Email string `json:"email" binding:"required,email" example:"georgios.papadopoulos@example.com"`
	// @Description The user's password
	Password string `json:"password" binding:"required,min=8" example:"strongPassword123!"`
	// @Description The user's first name
	FirstName string `json:"first_name" binding:"required" example:"Γεώργιος"`
	// @Description The user's last name
	LastName string `json:"last_name" binding:"required" example:"Παπαδόπουλος"`
	// @Description The user's role in the system
	Role string `json:"role" binding:"required,oneof=patient doctor receptionist" example:"patient"`
}

// LoginCredentials represents the data needed for user login
// @Description Login credentials
type LoginCredentials struct {
	// @Description The user's email address
	Email string `json:"email" binding:"required,email" example:"georgios.papadopoulos@example.com"`
	// @Description The user's password
	Password string `json:"password" binding:"required" example:"strongPassword123!"`
}

// LoginResponse represents the response after successful login
// @Description Login response
type LoginResponse struct {
	// @Description JWT token for authentication
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// Claims represents the claims in a JWT token
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Patient represents a patient in the system
// @Description Patient information
type Patient struct {
	SwaggerGormModel
	// @Description The ID of the associated user
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null" example:"1"`
	// @Description The patient's AMKA (Social Security Number)
	AMKA string `json:"amka" gorm:"uniqueIndex;not null" example:"12345678901"`
	// @Description The date of patient registration
	Registration time.Time `json:"registration" gorm:"not null" example:"2023-01-01T00:00:00Z"`
}

// Appointment represents a medical appointment
// @Description Appointment information
type Appointment struct {
	SwaggerGormModel
	// @Description The ID of the patient
	PatientID uint `json:"patient_id" gorm:"not null" example:"1"`
	// @Description The ID of the doctor
	DoctorID uint `json:"doctor_id" gorm:"not null" example:"2"`
	// @Description The date and time of the appointment
	DateTime time.Time `json:"date_time" gorm:"not null" example:"2023-01-01T10:00:00Z"`
	// @Description The reason for the appointment
	Reason string `json:"reason" gorm:"not null" example:"Τακτικός έλεγχος"`
	// @Description The status of the appointment
	Status string `json:"status" gorm:"not null;check:status IN ('Created','Completed','Canceled')" example:"Created"`
}

// MedicalHistory represents a patient's medical history
// @Description Medical history information
type MedicalHistory struct {
	SwaggerGormModel
	// @Description The ID of the patient
	PatientID uint `json:"patient_id" gorm:"uniqueIndex;not null" example:"1"`
}

// HistoryEntry represents an entry in a patient's medical history
// @Description History entry information
type HistoryEntry struct {
	SwaggerGormModel
	// @Description The ID of the associated medical history
	MedicalHistoryID uint `json:"medical_history_id" gorm:"not null" example:"1"`
	// @Description The date of the entry
	Date time.Time `json:"date" gorm:"not null" example:"2023-01-01T00:00:00Z"`
	// @Description The health issues recorded
	HealthIssues string `json:"health_issues" gorm:"not null" example:"Πονοκέφαλος και πυρετός"`
	// @Description The treatment prescribed
	Treatment string `json:"treatment" gorm:"not null" example:"Συνταγογράφηση παυσίπονων και ανάπαυση"`
}
