package services

import (
	"movie-app/internal/core/reqres"
	"movie-app/utils/pagination"

	"github.com/google/uuid"
)

type UserServices interface {
	ListMovies(page pagination.Pagination) ([]reqres.MovieResponse, error)
	SearchMovies(q string) ([]reqres.MovieResponse, error)
	History()
	VoteMovie()
	WatchMovie(movieid uuid.UUID, payload reqres.WatchMovieReq) (reqres.WatchMovie, error)
}
