package services

import (
	"movie-app/internal/core/reqres"
	"movie-app/utils/pagination"
)

type UserServices interface {
	ListMovies(page pagination.Pagination) ([]reqres.MovieResponse, error)
	SearchMovies()
	History()
	VoteMovie()
}
