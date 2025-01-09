package services

import (
	"movie-app/internal/core/reqres"
	"movie-app/internal/models"
)

type AdminServices interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovie(id int) (models.Movie, error)
	CreateMovie(movie reqres.CreateMovieRequest) (models.Movie, error)
}
