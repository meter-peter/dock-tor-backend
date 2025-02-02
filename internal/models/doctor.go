package models

import (
	"time"
)

// Doctor represents a doctor in the clinic
// @Description Doctor information
type Doctor struct {
	SwaggerGormModel
	// @Description The ID of the associated user
	UserID uint `json:"user_id" gorm:"uniqueIndex" example:"1"`
	// @Description The associated user
	User User `json:"user" gorm:"foreignKey:UserID"`
	// @Description The doctor's specialty
	Specialty string `json:"specialty" gorm:"not null" example:"Cardiology"`
}

// DoctorAvailability represents a doctor's availability
// @Description Doctor availability information
type DoctorAvailability struct {
	SwaggerGormModel
	// @Description The ID of the doctor
	DoctorID uint `json:"doctor_id" gorm:"not null" example:"1"`
	// @Description The associated doctor
	Doctor Doctor `json:"doctor" gorm:"foreignKey:DoctorID"`
	// @Description The start time of availability
	StartTime time.Time `json:"start_time" gorm:"not null" example:"2023-01-01T09:00:00Z"`
	// @Description The end time of availability
	EndTime time.Time `json:"end_time" gorm:"not null" example:"2023-01-01T17:00:00Z"`
	// @Description Whether the doctor is available during this time
	IsAvailable bool `json:"is_available" gorm:"not null;default:true" example:"true"`
}
