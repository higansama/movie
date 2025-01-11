package services

import (
	"movie-app/internal/core/repositories"
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/core/reqres"
	"movie-app/internal/core/services"
	"movie-app/internal/models"
	"movie-app/internal/utils"
	"movie-app/utils/exception"
	"movie-app/utils/infra"
	"movie-app/utils/pagination"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AdminServiceImpl struct {
	infra       *infra.Infrastructure
	MovieRepo   repositories.MovieRepository
	CastingRepo repositories.CastingRepo
}

// History implements services.UserServices.
func (a *AdminServiceImpl) History() {
	panic("unimplemented")
}

// ListMovies implements services.UserServices.
func (a *AdminServiceImpl) ListMovies(page pagination.Pagination) ([]reqres.MovieResponse, error) {
	panic("unimplemented")
}

// SearchMovies implements services.UserServices.
func (a *AdminServiceImpl) SearchMovies() {
	panic("unimplemented")
}

// VoteMovie implements services.UserServices.
func (a *AdminServiceImpl) VoteMovie() {
	panic("unimplemented")
}

// UploadMovie implements services.AdminServices.
func (a *AdminServiceImpl) UploadMovie(path string, movieID string) error {
	uid, _ := utils.StringToUUID(movieID)
	movie, err := a.MovieRepo.FindByID(uid)
	if err != nil {
		return exception.NewErrorMovie(500, "error cari data", err)
	}

	movie.Files = path
	return a.MovieRepo.Update(uid, movie)
}

// EditMovie implements services.AdminServices.
func (a *AdminServiceImpl) EditMovie(id uuid.UUID, payload reqres.EditMovieRequest) error {
	movie, _ := a.MovieRepo.FindByID(id)
	if movie == nil {
		return exception.NewErrorMovie(500, "movie not found", nil)
	}

	movie.Title = payload.Title
	movie.Director = payload.Director
	movie.Description = payload.Description
	movie.Duration = payload.Duration
	movie.Year = payload.Year
	movie.UpdatedAt = time.Now()

	return a.MovieRepo.Update(id, movie)
}

// GetAllMovies implements services.AdminServices.
func (a *AdminServiceImpl) GetAllMovies(page pagination.Pagination) ([]reqres.MovieResponse, error) {
	result := make([]reqres.MovieResponse, 0)

	movies, err := a.MovieRepo.FindAll(page)
	if err != nil {
		return result, exception.NewErrorMovie(500, "error get movie", err)
	}

	for _, v := range movies {
		t := reqres.MovieResponse{
			ID:          v.ID.String(),
			Title:       v.Title,
			Slug:        v.Slug,
			Director:    v.Director,
			Description: v.Description,
			Duration:    v.Duration,
			// Genres:      v.Genres,
			Files:      v.Files,
			Year:       v.Year,
			Count:      v.Count,
			UploadedAt: v.UploadedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		result = append(result, t)
	}

	return result, nil
}

// CreateMovie implements services.AdminServices.
func (a *AdminServiceImpl) CreateMovie(payload reqres.CreateMovieRequest) (reqres.MovieResponse, error) {
	MovieID := uuid.New()

	// insert to casting
	casts := make([]models.Casting, 0)
	for _, casting := range payload.Actor {
		ActorID, err := utils.StringToUUID(casting.ID)
		if err != nil {
			return reqres.MovieResponse{}, exception.NewErrorMovie(500, "invalid id actor", err)
		}
		castingToInsert := models.Casting{
			ID:      uuid.New(),
			MovieID: MovieID,
			ActorID: ActorID,
			Role:    casting.Role,
		}
		casts = append(casts, castingToInsert)
	}

	// insert to movie genre
	var genres []models.Genre
	for _, v := range payload.Movie.Genres {
		genreID, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return reqres.MovieResponse{}, exception.NewErrorMovie(500, "invalid genre id", err)
		}
		genres = append(genres, models.Genre{
			ID: uint(genreID),
		})
	}

	dataToInsert := models.Movie{
		ID:          MovieID,
		Title:       payload.Movie.Title,
		Director:    payload.Movie.Director,
		Description: payload.Movie.Description,
		Duration:    payload.Movie.Duration,
		Genre:       genres,
		Year:        payload.Movie.Year,
	}

	err := a.MovieRepo.Create(&dataToInsert)
	if err != nil {
		return reqres.MovieResponse{}, exception.NewErrorMovie(500, "error creating movie", err)
	}

	// err = a.CastingRepo.AddActorsToMovies(casts)
	// if err != nil {
	// 	return reqres.MovieResponse{}, exception.NewErrorMovie(500, "error creating casting", err)
	// }

	response := reqres.MovieResponse{
		ID:          dataToInsert.ID.String(),
		Title:       dataToInsert.Title,
		Slug:        dataToInsert.Slug,
		Director:    dataToInsert.Director,
		Description: dataToInsert.Description,
		Duration:    dataToInsert.Duration,
		Files:       dataToInsert.Files,
		Year:        dataToInsert.Year,
		Count:       dataToInsert.Count,
		UploadedAt:  dataToInsert.UploadedAt,
		UpdatedAt:   dataToInsert.UpdatedAt,
	}

	return response, nil
}

// GetMovie implements services.AdminServices.
func (a *AdminServiceImpl) GetMovie(id uuid.UUID) (models.Movie, error) {
	// movie, err := a.MovieRepo.FindByID(id)
	return models.Movie{}, nil
}

func NewAdminServices(
	infra *infra.Infrastructure,
	movieRepo coreRepo.MovieRepository,
	castingRepo coreRepo.CastingRepo,
) (services.AdminServices, error) {
	u := &AdminServiceImpl{
		infra:       infra,
		MovieRepo:   movieRepo,
		CastingRepo: castingRepo,
	}

	return u, nil
}
