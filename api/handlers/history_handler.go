package handlers

import (
	"net/http"

	"clinic-management/config"
	"clinic-management/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary Add history entry
// @Tags History
// @Security BearerAuth
// @Param entry body models.HistoryEntry true "History entry"
// @Success 201 {object} models.HistoryEntry
// @Router /history [post]
func AddHistoryEntry(c *gin.Context) {
	var entry models.HistoryEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

// Add other history handlers as needed
