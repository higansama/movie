package entrypoint

import (
	"movie-app/internal/config"
	"movie-app/internal/controllers"
	"movie-app/internal/repositories"
	"movie-app/internal/services"
	"movie-app/utils/infra"

	"github.com/gin-gonic/gin"
)

func NewAdminModule(engine gin.Engine, config config.Config, infra infra.Infrastructure) error {
	movieRepo := repositories.NewMovieRepository(infra.GormConnection)
	castRepo := repositories.NewCastingImplementation(infra.GormConnection)
	adminService, err := services.NewAdminServices(&infra, movieRepo, castRepo)
	if err != nil {
		return err
	}

	route := controllers.NewAdminController(
		&engine,
		adminService,
		config,
		infra,
	)

	return route
}
