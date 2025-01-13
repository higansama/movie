package controllers

import (
	"movie-app/internal/config"
	"movie-app/internal/core/reqres"
	"movie-app/internal/core/services"
	"movie-app/internal/utils"
	"movie-app/utils/auth"
	"movie-app/utils/exception"
	"movie-app/utils/infra"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Config       config.Config
	UserServices services.UserServices
	Infra        infra.Infrastructure
}

func NewUserController(
	engine *gin.Engine,
	userService services.UserServices,
	cfg config.Config,
	infra infra.Infrastructure) error {

	handler := &UserController{
		Config:       cfg,
		UserServices: userService,
		Infra:        infra,
	}

	userRoute := engine.Group("user")
	userRoute.Use(infra.Middleware.UserMiddleware)
	userRoute.GET("watch/:id", handler.WathcMovie)
	userRoute.GET("search", handler.Find)
	userRoute.POST("register", handler.Register)
	userRoute.POST("login", handler.Login)

	return nil
}

func (uc *UserController) WathcMovie(c *gin.Context) {
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
	var payload reqres.WatchMovieReq
	result, err := uc.UserServices.WatchMovie(uuid, payload)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}

	reqres.JsonResponse(c, nil, result)
}

func (uc *UserController) Find(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		q = "*"
	}

	result, err := uc.UserServices.SearchMovies(q)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}

	reqres.JsonResponse(c, nil, result)
}

func (uc *UserController) Register(c *gin.Context) {
	var payload reqres.UserRegister
	// Bind JSON data
	if err := c.ShouldBindJSON(&payload); err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request form", err), nil)
		return
	}

	// Set the file path and ID in the payload
	err := uc.UserServices.Register(payload)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}
	reqres.JsonResponse(c, nil, gin.H{"message": "success"})
}

func (uc *UserController) Login(c *gin.Context) {
	var payload reqres.LoginRequest
	// Bind JSON data
	if err := c.ShouldBindJSON(&payload); err != nil {
		reqres.JsonResponse(c, exception.NewErrorMovie(400, "Bad request form", err), nil)
		return
	}

	// Set the file path and ID in the payload
	r, err := uc.UserServices.Login(payload)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}

	token, err := auth.GenerateAuthToken(uc.Config, r)
	if err != nil {
		reqres.JsonResponse(c, err, nil)
		return
	}

	reqres.JsonResponse(c, nil, gin.H{"token": token})
}
