package handlers

import (
	"clinic-management/internal/models"
	"clinic-management/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	service *services.DoctorService
}

func NewDoctorHandler(service *services.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

// @Summary Create a new doctor
// @Description Create a new doctor
// @Tags Doctors
// @Accept json
// @Produce json
// @Param doctor body models.Doctor true "Doctor object"
// @Success 201 {object} models.Doctor
// @Router /doctors [post]
func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDoctor(&doctor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

// @Summary Get a doctor by ID
// @Description Get a doctor's details by their ID
// @Tags Doctors
// @Produce json
// @Param id path int true "Doctor ID"
// @Success 200 {object} models.Doctor
// @Router /doctors/{id} [get]
func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	doctor, err := h.service.GetDoctor(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

// @Summary Update a doctor
// @Description Update a doctor's details
// @Tags Doctors
// @Accept json
// @Produce json
// @Param id path int true "Doctor ID"
// @Param doctor body models.Doctor true "Updated doctor object"
// @Success 200 {object} models.Doctor
// @Router /doctors/{id} [put]
func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor.ID = uint(id)
	if err := h.service.UpdateDoctor(&doctor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

// @Summary Delete a doctor
// @Description Delete a doctor by their ID
// @Tags Doctors
// @Param id path int true "Doctor ID"
// @Success 200 {object} map[string]string
// @Router /doctors/{id} [delete]
func (h *DoctorHandler) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	if err := h.service.DeleteDoctor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

// @Summary List all doctors
// @Description Get a list of all doctors
// @Tags Doctors
// @Produce json
// @Success 200 {array} models.Doctor
// @Router /doctors [get]
func (h *DoctorHandler) ListDoctors(c *gin.Context) {
	doctors, err := h.service.ListDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

// @Summary Create doctor availability
// @Description Create a new availability slot for a doctor
// @Tags Doctors
// @Accept json
// @Produce json
// @Param id path int true "Doctor ID"
// @Param availability body models.DoctorAvailability true "Availability object"
// @Success 201 {object} map[string]string
// @Router /doctors/{id}/availabilities [post]
func (h *DoctorHandler) CreateAvailability(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var input struct {
		StartTime string `json:"start_time" binding:"required"`
		EndTime   string `json:"end_time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
		return
	}

	if err := h.service.CreateAvailability(uint(doctorID), startTime, endTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create availability"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Availability created successfully"})
}

// @Summary Get doctor availabilities
// @Description Get all availability slots for a doctor
// @Tags Doctors
// @Produce json
// @Param id path int true "Doctor ID"
// @Success 200 {array} models.DoctorAvailability
// @Router /doctors/{id}/availabilities [get]
func (h *DoctorHandler) GetAvailabilities(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	availabilities, err := h.service.GetAvailabilities(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get availabilities"})
		return
	}

	c.JSON(http.StatusOK, availabilities)
}

// @Summary Delete doctor availability
// @Description Delete an availability slot for a doctor
// @Tags Doctors
// @Param id path int true "Availability ID"
// @Success 200 {object} map[string]string
// @Router /doctors/availabilities/{id} [delete]
func (h *DoctorHandler) DeleteAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid availability ID"})
		return
	}

	if err := h.service.DeleteAvailability(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete availability"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Availability deleted successfully"})
}
