package repositories

import "movie-app/internal/models"

type UserRepositories interface {
	Register(data models.User) error
	FindUser(username string) (result models.User, err error)
}
