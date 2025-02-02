package repositories

import (
	"clinic-management/config"
	"clinic-management/internal/models"
	"time"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{db: config.DB}
}

func (r *HistoryRepository) AddHistoryEntry(entry *models.HistoryEntry) error {
	return r.db.Create(entry).Error
}

func (r *HistoryRepository) GetPatientHistory(medicalHistoryID uint) ([]models.HistoryEntry, error) {
	var entries []models.HistoryEntry
	err := r.db.Where("medical_history_id = ?", medicalHistoryID).Order("date DESC").Find(&entries).Error
	return entries, err
}

func (r *HistoryRepository) UpdateLatestEntry(entry *models.HistoryEntry) (*models.HistoryEntry, error) {
	var latestEntry models.HistoryEntry
	err := r.db.Where("medical_history_id = ?", entry.MedicalHistoryID).Order("date DESC").First(&latestEntry).Error
	if err != nil {
		return nil, err
	}

	latestEntry.HealthIssues = entry.HealthIssues
	latestEntry.Treatment = entry.Treatment

	err = r.db.Save(&latestEntry).Error
	return &latestEntry, err
}

func (r *HistoryRepository) DeleteLatestEntry(medicalHistoryID uint) error {
	var latestEntry models.HistoryEntry
	err := r.db.Where("medical_history_id = ?", medicalHistoryID).Order("date DESC").First(&latestEntry).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&latestEntry).Error
}

func (r *HistoryRepository) SearchHistory(medicalHistoryID uint, startDate, endDate time.Time, healthIssue string) ([]models.HistoryEntry, error) {
	var entries []models.HistoryEntry
	query := r.db.Where("medical_history_id = ?", medicalHistoryID)

	if !startDate.IsZero() {
		query = query.Where("date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("date <= ?", endDate)
	}
	if healthIssue != "" {
		query = query.Where("health_issues LIKE ?", "%"+healthIssue+"%")
	}

	err := query.Order("date DESC").Find(&entries).Error
	return entries, err
}
