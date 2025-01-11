package repositories

import "github.com/google/uuid"

type GenreRepo interface {
	AddCountUp(movieID uuid.UUID) error
}
