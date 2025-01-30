package handlers

import (
	"encoding/csv"
	"net/http"

	"clinic-management/internal/models"
	"clinic-management/internal/services"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service services.PatientService
}

func NewPatientHandler(service services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

// CreatePatient godoc
// @Summary Create a new patient
// @Description Add a new patient to the system
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} utils.ErrorResponse
// @Failure 409 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /patients [post]

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatePatient(&patient); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

// @Summary Bulk upload patients
// @Tags Patients
// @Security BearerAuth
// @Param file formData file true "CSV file"
// @Success 202 {object} map[string]interface{}
// @Router /patients/bulk [post]
func (h *PatientHandler) BulkUploadPatients(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CSV format"})
		return
	}

	successCount, errors := h.service.BulkCreatePatients(records[1:]) // Skip header

	response := gin.H{
		"message":       "Bulk upload processed",
		"success_count": successCount,
		"error_count":   len(errors),
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	c.JSON(http.StatusAccepted, response)
}

// Implement remaining interface methods
func (h *PatientHandler) GetPatient(c *gin.Context) {
	// Implementation
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	// Implementation
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	// Implementation
}

func (h *PatientHandler) ListPatients(c *gin.Context) {
	// Implementation
}
