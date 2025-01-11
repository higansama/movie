package services

import (
	"movie-app/internal/core/reqres"
	"movie-app/internal/models"
	"movie-app/utils/pagination"

	"github.com/google/uuid"
)

type AdminServices interface {
	GetAllMovies(page pagination.Pagination) ([]reqres.MovieResponse, error)
	GetMovie(id uuid.UUID) (models.Movie, error)
	EditMovie(id uuid.UUID, payload reqres.EditMovieRequest) error
	CreateMovie(movie reqres.CreateMovieRequest) (reqres.MovieResponse, error)
	UploadMovie(path, movieID string) error
}
