package services

import (
	"movie-app/internal/core/repositories"
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/core/reqres"
	coreService "movie-app/internal/core/services"
	"movie-app/internal/models"
	"movie-app/internal/utils"
	"movie-app/utils/exception"
	"movie-app/utils/infra"
	"movie-app/utils/pagination"
	"strings"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	infra        *infra.Infrastructure
	MovieRepo    repositories.MovieRepository
	CastingRepo  repositories.CastingRepo
	GenreRepo    repositories.GenreRepo
	Wathcingrepo repositories.WathcingHistoryRepository
}

// WatchMovie implements services.UserServices.
func (u *UserServiceImpl) WatchMovie(movieID uuid.UUID, payload reqres.WatchMovieReq) (response reqres.WatchMovie, err error) {
	movie, err := u.MovieRepo.FindByID(movieID)
	if err != nil {
		return response, exception.NewErrorMovie(500, "error get movie", err)
	}

	// add to watch history
	err = u.MovieRepo.IncreaseMovieWatcher(movie)
	if err != nil {
		return response, exception.NewErrorMovie(500, "count up movie", err)
	}

	// add genre repo
	err = u.GenreRepo.AddCountUp(movie.ID)
	if err != nil {
		return response, exception.NewErrorMovie(500, "cant count up genre", err)
	}

	// add to watching history
	var userid *uuid.UUID
	if payload.IsRegistered {
		userid = utils.ConvertStringToPointerUUID(payload.ID)
	}
	wHistory := &models.WathcingHistory{
		UserID:  userid,
		MovieID: movieID,
	}
	err = u.Wathcingrepo.AddWatchingHistory(wHistory)
	if err != nil {
		return response, exception.NewErrorMovie(500, "cant add to watching", err)
	}

	link := strings.ReplaceAll(movie.Files, "\\", "/")
	response.Link = "http://" + u.infra.Config.AppAttribute.Host + ":" + u.infra.Config.AppAttribute.Port + "/" + link
	return response, nil
}

// History implements services.UserServices.
func (u *UserServiceImpl) History() {
	panic("unimplemented")
}

// SearchMovies implements services.UserServices.
func (u *UserServiceImpl) SearchMovies(q string) (response []reqres.MovieResponse, err error) {
	r, err := u.MovieRepo.FindByQword(q)
	if err != nil {
		return response, exception.NewErrorMovie(500, "error find data", err)
	}
	for _, v := range r {
		rTempt := reqres.MovieResponse{
			ID:          v.ID.String(),
			Title:       v.Title,
			Slug:        v.Slug,
			Director:    v.Director,
			Description: v.Description,
			Duration:    v.Duration,
			Files:       v.Files,
			Year:        v.Year,
			Count:       v.Count,
			UploadedAt:  v.UploadedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		response = append(response, rTempt)
	}

	return response, nil
}

// VoteMovie implements services.UserServices.
func (u *UserServiceImpl) VoteMovie() {
	panic("unimplemented")
}

// ListMovies implements services.UserServices.
func (u *UserServiceImpl) ListMovies(page pagination.Pagination) ([]reqres.MovieResponse, error) {
	panic("unimplemented")
}

func NewUserServices(
	infra *infra.Infrastructure,
	movieRepo coreRepo.MovieRepository,
	castingRepo coreRepo.CastingRepo,
	GenreRepo coreRepo.GenreRepo,
	Wathcingrepo coreRepo.WathcingHistoryRepository,
) (coreService.UserServices, error) {
	u := &UserServiceImpl{
		infra:        infra,
		MovieRepo:    movieRepo,
		CastingRepo:  castingRepo,
		GenreRepo:    GenreRepo,
		Wathcingrepo: Wathcingrepo,
	}

	return u, nil
}
