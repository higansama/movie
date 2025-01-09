package repositories

import (
	"movie-app/internal/models"

	"github.com/jinzhu/gorm"
)

type MovieRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{db: db}
}

func (repo *MovieRepositoryImpl) Create(movie *models.Movie) error {
	return repo.db.Create(movie).Error
}

func (repo *MovieRepositoryImpl) Update(movie *models.Movie) error {
	return repo.db.Save(movie).Error
}

func (repo *MovieRepositoryImpl) Delete(id uint) error {
	return repo.db.Delete(&models.Movie{}, id).Error
}

func (repo *MovieRepositoryImpl) Hide(id uint) error {
	return repo.db.Model(&models.Movie{}).Where("id = ?", id).Update("hidden", true).Error
}

func (repo *MovieRepositoryImpl) List() ([]models.Movie, error) {
	var movies []models.Movie
	err := repo.db.Find(&movies).Error
	return movies, err
}

func (repo *MovieRepositoryImpl) FindById(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := repo.db.First(&movie, id).Error
	return &movie, err
}
