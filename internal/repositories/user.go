package repositories

import (
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/models"

	"gorm.io/gorm"
)

type UserRepoImplementation struct {
	db *gorm.DB
}

// FindUser implements repositories.UserRepositories.
func (u *UserRepoImplementation) FindUser(username string) (result models.User, err error) {
	err = u.db.Find(&result).Error
	return result, err
}

// Register implements repositories.UserRepositories.
func (u *UserRepoImplementation) Register(data models.User) error {
	return u.db.Create(&data).Error
}

func NewUserRepo(db *gorm.DB) coreRepo.UserRepositories {
	return &UserRepoImplementation{db: db}
}
