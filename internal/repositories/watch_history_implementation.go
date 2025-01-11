package repositories

import (
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WathcingHistoryImplementation struct {
	db *gorm.DB
}

// AddWatchingHistory implements repositories.WathcingHistoryRepository.
func (w *WathcingHistoryImplementation) AddWatchingHistory(data *models.WathcingHistory) error {
	return w.db.Create(data).Error
}

// GetUserHistory implements repositories.WathcingHistoryRepository.
func (w *WathcingHistoryImplementation) GetUserHistory(user uuid.UUID) (result []coreRepo.MovieRepository, err error) {
	panic("unimplemented")
}

func NewWathcingRepository(db *gorm.DB) coreRepo.WathcingHistoryRepository {
	return &WathcingHistoryImplementation{db: db}
}
