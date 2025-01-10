package repositories

import (
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/models"

	"gorm.io/gorm"
)

type CastsImplementation struct {
	db *gorm.DB
}

// AddActorsToMovies implements repositories.CastingRepo.
func (c *CastsImplementation) AddActorsToMovies(data []models.Casting) error {
	if len(data) == 0 {
		return nil
	}
	return c.db.CreateInBatches(data, len(data)).Error
}

func NewCastingImplementation(db *gorm.DB) coreRepo.CastingRepo {
	return &CastsImplementation{db: db}
}
