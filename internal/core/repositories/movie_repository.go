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
	AddMovieToGenre(movieID uuid.UUID, movie *models.Movie) error
	IncreaseMovieWatcher(move *models.Movie) error
	FindByQword(word string) ([]models.Movie, error)
	VoteMovie(movie *models.Movie, vote *models.VotingHistory) error
}
