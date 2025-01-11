package repositories

import (
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/models"
	"movie-app/utils/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	db *gorm.DB
}

// Create implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) Create(movie *models.Movie) error {
	if err := m.db.Create(movie).Error; err != nil {
		return err
	}
	if err := m.db.Model(movie).Association("Genres").Append(movie.Genre); err != nil {
		return err
	}
	if err := m.db.Model(movie).Association("Castings").Append(movie.Casting); err != nil {
		return err
	}
	return nil
}

// Delete implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) Delete(id uint) error {
	panic("unimplemented")
}

// FindAll implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) FindAll(page pagination.Pagination) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (page.Page - 1) * page.Limit
	if err := m.db.Limit(page.Limit).Offset(offset).Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

// FindByID implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) FindByID(id uuid.UUID) (*models.Movie, error) {
	var result models.Movie
	err := m.db.Find(&result, "id = ?", id.String()).Error
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// Hide implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) Hide(id uint) error {
	panic("unimplemented")
}

// Update implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) Update(id uuid.UUID, movie *models.Movie) error {
	var t models.Movie
	if err := m.db.First(&t, id).Error; err != nil {
		return err
	}
	return m.db.Save(&movie).Error
}

// Update implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) UpdateRaw(id string, data map[string]interface{}) error {
	return m.db.Model(&models.Movie{}).Where("id = ?", id).Updates(data).Error
}

// AddMovieToGenre implements repositories.MovieRepository.
func (m *MovieRepositoryImpl) AddMovieToGenre(movieID uuid.UUID, movie *models.Movie) error {
	m.db.Model(models.Movie{}).Association("Genres").Append(movie.Genre)
	return nil
}

func NewMovieRepository(db *gorm.DB) coreRepo.MovieRepository {
	return &MovieRepositoryImpl{db: db}
}
