// api/handlers/doctor_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"clinic-management/internal/models"
	"clinic-management/internal/services"
	"clinic-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	service *services.DoctorService
}

func NewDoctorHandler(service *services.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

// @Summary Create a new doctor
// @Tags Doctors
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param doctor body models.Doctor true "Doctor object"
// @Success 201 {object} models.Doctor
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors [post]
func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateDoctor(&doctor); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create doctor")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, doctor)
}

// @Summary Get a doctor by ID
// @Tags Doctors
// @Security BearerAuth
// @Produce json
// @Param id path int true "Doctor ID"
// @Success 200 {object} models.Doctor
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /doctors/{id} [get]
func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}

	doctor, err := h.service.GetDoctor(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Doctor not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, doctor)
}

// @Summary Update a doctor
// @Tags Doctors
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Doctor ID"
// @Param doctor body models.Doctor true "Updated doctor object"
// @Success 200 {object} models.Doctor
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors/{id} [put]
func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}

	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	doctor.ID = uint(id)
	if err := h.service.UpdateDoctor(&doctor); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update doctor")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, doctor)
}

// @Summary Delete a doctor
// @Tags Doctors
// @Security BearerAuth
// @Param id path int true "Doctor ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors/{id} [delete]
func (h *DoctorHandler) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}

	if err := h.service.DeleteDoctor(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete doctor")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

// @Summary List all doctors
// @Tags Doctors
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Doctor
// @Failure 500 {object} utils.APIResponse
// @Router /doctors [get]
func (h *DoctorHandler) ListDoctors(c *gin.Context) {
	doctors, err := h.service.ListDoctors()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list doctors")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, doctors)
}

// @Summary Create doctor availability
// @Tags Doctors
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Doctor ID"
// @Param availability body models.DoctorAvailability true "Availability object"
// @Success 201 {object} models.DoctorAvailability
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors/{id}/availabilities [post]
func (h *DoctorHandler) CreateAvailability(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}

	var availability models.DoctorAvailability
	if err := c.ShouldBindJSON(&availability); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	availability.DoctorID = uint(doctorID)
	if err := h.service.CreateAvailability(&availability); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create availability")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, availability)
}

// @Summary Get doctor availabilities
// @Tags Doctors
// @Security BearerAuth
// @Produce json
// @Param id path int true "Doctor ID"
// @Success 200 {array} models.DoctorAvailability
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors/{id}/availabilities [get]
func (h *DoctorHandler) GetAvailabilities(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}

	availabilities, err := h.service.GetAvailabilities(uint(doctorID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get availabilities")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, availabilities)
}

// @Summary Delete doctor availability
// @Tags Doctors
// @Security BearerAuth
// @Param id path int true "Availability ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /doctors/availabilities/{id} [delete]
func (h *DoctorHandler) DeleteAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid availability ID")
		return
	}

	if err := h.service.DeleteAvailability(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete availability")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Availability deleted successfully"})
}
