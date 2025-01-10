package repositories

import (
	"movie-app/internal/models"
	"movie-app/utils/pagination"

	"github.com/google/uuid"
)

type MovieRepository interface {
	Create(movie *models.Movie) error
	Update(id uuid.UUID, movie *models.Movie) error
	Delete(id uint) error
	Hide(id uint) error
	FindByID(id uuid.UUID) (*models.Movie, error)
	FindAll(page pagination.Pagination) ([]models.Movie, error)
}
