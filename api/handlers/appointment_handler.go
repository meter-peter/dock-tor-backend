// api/handlers/appointment_handler.go

package handlers

import (
	"net/http"
	"strconv"
	"time"

	"clinic-management/internal/models"
	"clinic-management/internal/services"
	"clinic-management/internal/utils"

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
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /appointments [post]
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateAppointment(&appointment); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create appointment")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, appointment)
}

// @Summary Get appointment
// @Tags Appointments
// @Security BearerAuth
// @Param id path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /appointments/{id} [get]
func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	appointment, err := h.service.GetAppointment(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Appointment not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, appointment)
}

// @Summary Update appointment
// @Tags Appointments
// @Security BearerAuth
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "Updated appointment data"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /appointments/{id} [put]
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	appointment.ID = uint(id)
	if err := h.service.UpdateAppointment(&appointment); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update appointment")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, appointment)
}

// @Summary Delete appointment
// @Tags Appointments
// @Security BearerAuth
// @Param id path int true "Appointment ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /appointments/{id} [delete]
func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	if err := h.service.DeleteAppointment(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to delete appointment")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

// @Summary List appointments
// @Tags Appointments
// @Security BearerAuth
// @Param startDate query string false "Start date (YYYY-MM-DD)"
// @Param endDate query string false "End date (YYYY-MM-DD)"
// @Param patientID query int false "Patient ID"
// @Param doctorID query int false "Doctor ID"
// @Success 200 {array} models.Appointment
// @Failure 500 {object} utils.APIResponse
// @Router /appointments [get]
func (h *AppointmentHandler) ListAppointments(c *gin.Context) {
	startDate, _ := time.Parse("2006-01-02", c.Query("startDate"))
	endDate, _ := time.Parse("2006-01-02", c.Query("endDate"))
	patientID, _ := strconv.ParseUint(c.Query("patientID"), 10, 32)
	doctorID, _ := strconv.ParseUint(c.Query("doctorID"), 10, 32)

	appointments, err := h.service.ListAppointments(startDate, endDate, uint(patientID), uint(doctorID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list appointments")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, appointments)
}
