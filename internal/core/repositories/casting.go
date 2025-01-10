package repositories

import "movie-app/internal/models"

type CastingRepo interface {
	AddActorsToMovies(data []models.Casting) error
}
