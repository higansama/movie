package repositories

import (
	"log"
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenreImplementation struct {
	db *gorm.DB
}

// AddCountUp implements repositories.GenreRepo.
func (g *GenreImplementation) AddCountUp(movieID uuid.UUID) error {
	var genres []models.Genre
	err := g.db.Table("genres").
		Select("genres.id, genres.title").
		Joins("JOIN movie_genres ON movie_genres.genre_id = genres.id").
		Where("movie_genres.movie_id = ?", movieID).
		Scan(&genres).Error

	if err != nil {
		log.Fatalf("Error fetching genres: %v", err)
	}
	ids := []uint{}
	for _, v := range genres {
		ids = append(ids, v.ID)
	}

	return g.db.Model(&models.Genre{}).Where("id in ?", ids).Update("count", gorm.Expr("count + ?", 1)).Error
}

func NewGenreRepository(db *gorm.DB) coreRepo.GenreRepo {
	return &GenreImplementation{db: db}
}
