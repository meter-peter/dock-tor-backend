package handlers

import (
	"net/http"
	"time"

	"clinic-management/internal/models"
	"clinic-management/internal/services"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	service *services.AppointmentService
}

func NewAppointmentHandler(service *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

// @Summary Create appointment
// @Tags Appointments
// @Security BearerAuth
// @Param appointment body models.Appointment true "Appointment data"
// @Success 201 {object} models.Appointment
// @Router /appointments [post]
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate appointment time
	if appointment.DateTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment time must be in the future"})
		return
	}

	if err := h.service.CreateAppointment(&appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

// Implement other required methods
func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	// Implementation
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	// Implementation
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	// Implementation
}

func (h *AppointmentHandler) ListAppointments(c *gin.Context) {
	// Implementation
}
