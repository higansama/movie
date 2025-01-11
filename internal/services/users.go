package services

import (
	"movie-app/internal/core/repositories"
	coreRepo "movie-app/internal/core/repositories"
	"movie-app/internal/core/services"
	"movie-app/utils/infra"
)

type UserServiceImpl struct {
	infra       *infra.Infrastructure
	MovieRepo   repositories.MovieRepository
	CastingRepo repositories.CastingRepo
}

func NewUserServices(
	infra *infra.Infrastructure,
	movieRepo coreRepo.MovieRepository,
	castingRepo coreRepo.CastingRepo,
) (services.UserServices, error) {
	u := &AdminServiceImpl{
		infra:       infra,
		MovieRepo:   movieRepo,
		CastingRepo: castingRepo,
	}

	return u, nil
}
