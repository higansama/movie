package controllers

import (
	"movie-app/internal/config"
	"movie-app/internal/core/reqres"
	"movie-app/internal/core/services"
	"movie-app/internal/utils"
	"movie-app/utils/exception"
	"movie-app/utils/infra"
	"movie-app/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Config       config.Config
	AdminService services.AdminServices
	Infra        infra.Infrastructure
}

func NewAdminController(
	engine *gin.Engine,
	adminService services.AdminServices,
	cfg config.Config,
	infra infra.Infrastructure) error {
	handler := &AdminController{
		Config:       cfg,
		AdminService: adminService,
		Infra:        infra,
	}

	adminRoute := engine.Group("admin/movie")
	adminRoute.POST("list", handler.ListMovies)
	adminRoute.POST("create", handler.UploadMovie)
	adminRoute.GET("/detail/:id", handler.UploadMovie)
	adminRoute.PUT("/edit/:id", handler.EditMovie)

	return nil
}

func (ac *AdminController) UploadMovie(c *gin.Context) {
	// Implement the logic to upload a movie
	var payload reqres.CreateMovieRequest
	err := c.ShouldBind(&payload)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request", err), nil)
		return
	}

	response, err := ac.AdminService.CreateMovie(payload)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request", err), nil)
		return
	}
	reqres.JsonResponse(c, nil, response)
}

func (ac *AdminController) EditMovie(c *gin.Context) {
	id := c.Param("id")
	uuid, err := utils.StringToUUID(id)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "invalid id", err), nil)
		return
	}

	var payload reqres.EditMovieRequest
	err = c.ShouldBind(&payload)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "bad request", err), nil)
		return
	}

	err = ac.AdminService.EditMovie(uuid, payload)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}

	reqres.JsonResponse(c, nil, gin.H{"message": "success"})
}

func (ac *AdminController) DeleteMovie(c *gin.Context) {
	// Implement the logic to delete a movie
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

func (ac *AdminController) ListMovies(c *gin.Context) {
	var pagination pagination.Pagination
	err := c.ShouldBind(&pagination)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request", err), nil)
		return
	}
	pagination.Validate()
	list, err := ac.AdminService.GetAllMovies(pagination)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(500, err.Error(), err), nil)
		return
	}

	reqres.JsonResponse(c, nil, list)
}

func (ac *AdminController) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request", nil), nil)
		return
	}
	uuid, err := utils.StringToUUID(id)
	if err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "ID is invalid", nil), nil)
		return
	}
	response, err := ac.AdminService.GetMovie(uuid)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}
	reqres.JsonResponse(c, nil, response)

}
