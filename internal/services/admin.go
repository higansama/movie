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

	"github.com/google/uuid"
)

type AdminServiceImpl struct {
	infra       *infra.Infrastructure
	MovieRepo   repositories.MovieRepository
	CastingRepo repositories.CastingRepo
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
	movie.Genres = payload.Genres
	movie.Year = payload.Year

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
			ID: v.ID.String(),
		}
		result = append(result, t)
	}

	return result, nil
}

// CreateMovie implements services.AdminServices.
func (a *AdminServiceImpl) CreateMovie(payload reqres.CreateMovieRequest) (reqres.MovieResponse, error) {
	dataToInsert := models.Movie{
		ID:          uuid.New(),
		Title:       payload.Movie.Title,
		Director:    payload.Movie.Director,
		Description: payload.Movie.Description,
		Duration:    payload.Movie.Duration,
		Genres:      payload.Movie.Genres,
		Files:       payload.Movie.Files,
		Year:        payload.Movie.Year,
	}

	err := a.MovieRepo.Create(&dataToInsert)
	if err != nil {
		return reqres.MovieResponse{}, exception.NewErrorMovie(500, "error creating movie", err)
	}

	casts := make([]models.Casting, 0)
	for _, casting := range payload.Actor {
		ActorID, err := utils.StringToUUID(casting.ID)
		if err != nil {
			return reqres.MovieResponse{}, exception.NewErrorMovie(500, "invalid id actor", err)
		}
		castingToInsert := models.Casting{
			ID:      uuid.New(),
			MovieID: dataToInsert.ID,
			ActorID: ActorID,
			Role:    casting.Role,
		}
		casts = append(casts, castingToInsert)
	}
	err = a.CastingRepo.AddActorsToMovies(casts)
	if err != nil {
		return reqres.MovieResponse{}, exception.NewErrorMovie(500, "error creating casting", err)
	}

	response := reqres.MovieResponse{
		ID:          dataToInsert.ID.String(),
		Title:       dataToInsert.Title,
		Slug:        dataToInsert.Slug,
		Director:    dataToInsert.Director,
		Description: dataToInsert.Description,
		Duration:    dataToInsert.Duration,
		Genres:      dataToInsert.Genres,
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
