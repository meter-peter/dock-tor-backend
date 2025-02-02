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

type HistoryHandler struct {
	service *services.HistoryService
}

func NewHistoryHandler(service *services.HistoryService) *HistoryHandler {
	return &HistoryHandler{service: service}
}

// @Summary Add history entry
// @Description Add a new entry to a patient's medical history
// @Tags History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param entry body models.HistoryEntry true "History entry"
// @Success 201 {object} models.HistoryEntry
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /history [post]
func (h *HistoryHandler) AddHistoryEntry(c *gin.Context) {
	var entry models.HistoryEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.AddHistoryEntry(&entry); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add history entry")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, entry)
}

// @Summary Get patient history
// @Description Retrieve the complete medical history for a patient
// @Tags History
// @Produce json
// @Security BearerAuth
// @Param medicalHistoryID path int true "Medical History ID"
// @Success 200 {array} models.HistoryEntry
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /history/{medicalHistoryID} [get]
func (h *HistoryHandler) GetPatientHistory(c *gin.Context) {
	medicalHistoryID, err := strconv.ParseUint(c.Param("medicalHistoryID"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID")
		return
	}

	history, err := h.service.GetPatientHistory(uint(medicalHistoryID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get patient history")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, history)
}

// @Summary Update latest history entry
// @Description Update the most recent entry in a patient's medical history
// @Tags History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param entry body models.HistoryEntry true "Updated history entry"
// @Success 200 {object} models.HistoryEntry
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /history/latest [put]
func (h *HistoryHandler) UpdateLatestEntry(c *gin.Context) {
	var entry models.HistoryEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedEntry, err := h.service.UpdateLatestEntry(&entry)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update latest entry")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, updatedEntry)
}

// @Summary Delete latest history entry
// @Description Delete the most recent entry in a patient's medical history
// @Tags History
// @Produce json
// @Security BearerAuth
// @Param medicalHistoryID path int true "Medical History ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /history/{medicalHistoryID}/latest [delete]
func (h *HistoryHandler) DeleteLatestEntry(c *gin.Context) {
	medicalHistoryID, err := strconv.ParseUint(c.Param("medicalHistoryID"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID")
		return
	}

	if err := h.service.DeleteLatestEntry(uint(medicalHistoryID)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete latest entry")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Latest history entry deleted successfully"})
}

// @Summary Search patient history
// @Description Search for specific entries in a patient's medical history
// @Tags History
// @Produce json
// @Security BearerAuth
// @Param medicalHistoryID query int true "Medical History ID"
// @Param startDate query string false "Start date (YYYY-MM-DD)"
// @Param endDate query string false "End date (YYYY-MM-DD)"
// @Param healthIssue query string false "Health issue"
// @Success 200 {array} models.HistoryEntry
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /history/search [get]
func (h *HistoryHandler) SearchHistory(c *gin.Context) {
	medicalHistoryID, err := strconv.ParseUint(c.Query("medicalHistoryID"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID")
		return
	}

	startDate, _ := time.Parse("2006-01-02", c.Query("startDate"))
	endDate, _ := time.Parse("2006-01-02", c.Query("endDate"))
	healthIssue := c.Query("healthIssue")

	results, err := h.service.SearchHistory(uint(medicalHistoryID), startDate, endDate, healthIssue)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search history")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, results)
}
