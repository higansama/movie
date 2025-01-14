package services

import (
	"movie-app/internal/core/repositories"
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/core/reqres"
	coreService "movie-app/internal/core/services"
	"movie-app/internal/models"
	"movie-app/internal/utils"
	"movie-app/utils/auth"
	"movie-app/utils/exception"
	"movie-app/utils/infra"
	"movie-app/utils/pagination"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	infra        *infra.Infrastructure
	MovieRepo    repositories.MovieRepository
	CastingRepo  repositories.CastingRepo
	GenreRepo    repositories.GenreRepo
	UserRepo     repositories.UserRepositories
	Wathcingrepo repositories.WathcingHistoryRepository
}

// Vote implements services.UserServices.
func (u *UserServiceImpl) Vote(payload reqres.VoteRequest) error {
	// get movie id
	movieID, _ := utils.StringToUUID(payload.MovieID)
	movie, err := u.MovieRepo.FindByID(movieID)
	if err != nil {
		return exception.NewErrorMovie(500, "movie is not found", err)
	}
	prevVote := movie.Vote

	if payload.Vote == 1 {
		movie.IncreaseMovieVote()
	} else {
		movie.DecreaseMovieVote()
	}

	// save vote to vote history
	userid := utils.ConvertStringToPointerUUID(payload.Auth.ID)
	votingHistory := &models.VotingHistory{
		UserID:            userid,
		MovieID:           movieID,
		IsLike:            payload.Vote == 1,
		PreviousVoteMovie: prevVote,
		CurrentViteMovie:  movie.Vote,
		DateCreated:       time.Now(),
	}
	return u.MovieRepo.VoteMovie(movie, votingHistory)
}

// Login implements services.UserServices.
func (u *UserServiceImpl) Login(payload reqres.LoginRequest) (response auth.AuthJWT, err error) {
	user, err := u.UserRepo.FindUser(payload.Username)
	if err != nil {
		return response, exception.NewErrorMovie(405, "user not found", err)
	}

	// validate password
	if !user.ValidPassword(payload.Password) {
		return response, exception.NewErrorMovie(405, "invalid username and password", nil)
	}

	response = auth.AuthJWT{
		ID:           user.ID.String(),
		Name:         user.Username,
		Role:         user.Role,
		IsRegistered: true,
	}
	return response, nil
}

// Register implements services.UserServices.
func (u *UserServiceImpl) Register(payload reqres.UserRegister) error {
	salt := auth.GenerateSalt()
	password := auth.GeneratePassword(salt, payload.Password)
	dataToInput := &models.User{
		ID:       uuid.New(),
		Username: payload.Username,
		Salt:     salt,
		Password: password,
	}
	return u.UserRepo.Register(*dataToInput)
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
	movies, err := u.MovieRepo.FindAll(page)
	if err != nil {
		return nil, exception.NewErrorMovie(500, "error fetching movies", err)
	}

	var response []reqres.MovieResponse
	for _, v := range movies {
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

func NewUserServices(
	infra *infra.Infrastructure,
	movieRepo coreRepo.MovieRepository,
	castingRepo coreRepo.CastingRepo,
	GenreRepo coreRepo.GenreRepo,
	Wathcingrepo coreRepo.WathcingHistoryRepository,
	UserRepo coreRepo.UserRepositories,
) (coreService.UserServices, error) {
	u := &UserServiceImpl{
		infra:        infra,
		MovieRepo:    movieRepo,
		CastingRepo:  castingRepo,
		GenreRepo:    GenreRepo,
		Wathcingrepo: Wathcingrepo,
		UserRepo:     UserRepo,
	}

	return u, nil
}
