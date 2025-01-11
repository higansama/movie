package repositories

import (
	"movie-app/internal/models"

	"github.com/google/uuid"
)

type WathcingHistoryRepository interface {
	AddWatchingHistory(data *models.WathcingHistory) error
	GetUserHistory(user uuid.UUID) (result []MovieRepository, err error)
}
