// api/handlers/patient_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"clinic-management/internal/models"
	"clinic-management/internal/services"
	"clinic-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service services.PatientService
}

func NewPatientHandler(service services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

// @Summary Create a new patient
// @Description Add a new patient to the system
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} utils.APIResponse
// @Failure 409 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /patients [post]
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreatePatient(&patient); err != nil {
		utils.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, patient)
}

// @Summary Get a patient
// @Description Get a patient's details by their ID
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Failure 404 {object} utils.APIResponse
// @Router /patients/{id} [get]
func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	patient, err := h.service.GetPatient(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Patient not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, patient)
}

// @Summary Update a patient
// @Description Update a patient's details
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Updated patient object"
// @Success 200 {object} models.Patient
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /patients/{id} [put]
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	patient.ID = uint(id)
	updatedPatient, err := h.service.UpdatePatient(&patient)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to update patient")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, updatedPatient)
}

// @Summary Delete a patient
// @Description Delete a patient by their ID
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /patients/{id} [delete]
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	if err := h.service.DeletePatient(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to delete patient")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

// @Summary List patients
// @Description Get a list of all patients
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Patient
// @Failure 500 {object} utils.APIResponse
// @Router /patients [get]
func (h *PatientHandler) ListPatients(c *gin.Context) {
	patients, err := h.service.ListPatients()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list patients")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, patients)
}

// @Summary Bulk upload patients
// @Description Upload multiple patients using a CSV file
// @Tags patients
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "CSV file"
// @Success 202 {object} map[string]interface{}
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /patients/bulk [post]
func (h *PatientHandler) BulkUploadPatients(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "File upload required")
		return
	}

	f, err := file.Open()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer f.Close()

	successCount, errors := h.service.BulkCreatePatients(f)
	response := gin.H{
		"message":       "Bulk upload processed",
		"success_count": successCount,
		"error_count":   len(errors),
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	utils.SuccessResponse(c, http.StatusAccepted, response)
}
