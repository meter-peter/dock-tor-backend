package services

import (
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"time"
)

type HistoryService struct {
	repo *repositories.HistoryRepository
}

func NewHistoryService(repo *repositories.HistoryRepository) *HistoryService {
	return &HistoryService{repo: repo}
}

func (s *HistoryService) AddHistoryEntry(entry *models.HistoryEntry) error {
	entry.Date = time.Now()
	return s.repo.AddHistoryEntry(entry)
}

func (s *HistoryService) GetPatientHistory(medicalHistoryID uint) ([]models.HistoryEntry, error) {
	return s.repo.GetPatientHistory(medicalHistoryID)
}

func (s *HistoryService) UpdateLatestEntry(entry *models.HistoryEntry) (*models.HistoryEntry, error) {
	return s.repo.UpdateLatestEntry(entry)
}

func (s *HistoryService) DeleteLatestEntry(medicalHistoryID uint) error {
	return s.repo.DeleteLatestEntry(medicalHistoryID)
}

func (s *HistoryService) SearchHistory(medicalHistoryID uint, startDate, endDate time.Time, healthIssue string) ([]models.HistoryEntry, error) {
	return s.repo.SearchHistory(medicalHistoryID, startDate, endDate, healthIssue)
}
