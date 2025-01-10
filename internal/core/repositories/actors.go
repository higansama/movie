package repositories

import (
	"movie-app/internal/models"

	"github.com/google/uuid"
)

type ActorRepo interface {
	Add(models.Actor) error
	GetByID(id uuid.UUID) (models.Actor, error)
	FindByMovieID(id uuid.UUID) ([]models.Actor, error)
}
