package repositories

import "movie-app/internal/models"

type MovieRepository interface {
    Create(movie *models.Movie) error
    Update(movie *models.Movie) error
    Delete(id uint) error
    Hide(id uint) error
    FindByID(id uint) (*models.Movie, error)
    FindAll() ([]models.Movie, error)
}