package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	IdentityID string `gorm:"uniqueIndex;not null"`
	Role       string `gorm:"not null;check:role IN ('patient','secretary','doctor')"`
}

type Patient struct {
	gorm.Model
	UserID       uint      `gorm:"uniqueIndex;not null"`
	AMKA         string    `gorm:"uniqueIndex;not null"`
	Registration time.Time `gorm:"not null"`
}

type Appointment struct {
	gorm.Model
	PatientID uint      `gorm:"not null"`
	DoctorID  uint      `gorm:"not null"`
	DateTime  time.Time `gorm:"not null"`
	Reason    string    `gorm:"not null"`
	Status    string    `gorm:"not null;check:status IN ('Created','Completed','Canceled')"`
}

type MedicalHistory struct {
	gorm.Model
	PatientID uint `gorm:"uniqueIndex;not null"`
}

type HistoryEntry struct {
	gorm.Model
	MedicalHistoryID uint      `gorm:"not null"`
	Date             time.Time `gorm:"not null"`
	HealthIssues     string    `gorm:"not null"`
	Treatment        string    `gorm:"not null"`
}

type DoctorAvailability struct {
	gorm.Model
	DoctorID    uint      `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	IsAvailable bool      `gorm:"not null;default:true"`
}
