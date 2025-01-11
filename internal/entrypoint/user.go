package entrypoint

import (
	"movie-app/internal/config"
	"movie-app/internal/controllers"
	"movie-app/internal/repositories"
	"movie-app/internal/services"
	"movie-app/utils/infra"

	"github.com/gin-gonic/gin"
)

func NewUserModule(engine *gin.Engine, config config.Config, infra infra.Infrastructure) error {
	movieRepo := repositories.NewMovieRepository(infra.GormConnection)
	castRepo := repositories.NewCastingImplementation(infra.GormConnection)
	wHistory := repositories.NewWathcingRepository(infra.GormConnection)
	genreRepo := repositories.NewGenreRepository(infra.GormConnection)
	userRepo := repositories.NewUserRepo(infra.GormConnection)
	userUsecase, err := services.NewUserServices(&infra, movieRepo, castRepo, genreRepo, wHistory, userRepo)
	if err != nil {
		return err
	}

	controllers.NewUserController(engine, userUsecase, config, infra)

	return nil
}
