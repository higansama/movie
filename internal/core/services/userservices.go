package services

import (
	"movie-app/internal/core/reqres"
	"movie-app/utils/auth"
	"movie-app/utils/pagination"

	"github.com/google/uuid"
)

type UserServices interface {
	Register(payload reqres.UserRegister) error
	Login(payload reqres.LoginRequest) (response auth.AuthJWT, err error)
	ListMovies(page pagination.Pagination) ([]reqres.MovieResponse, error)
	SearchMovies(q string) ([]reqres.MovieResponse, error)
	History()
	VoteMovie()
	WatchMovie(movieid uuid.UUID, payload reqres.WatchMovieReq) (reqres.WatchMovie, error)
	Vote(payload reqres.VoteRequest) error
}
