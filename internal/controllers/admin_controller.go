package controllers

import (
	"movie-app/internal/config"
	"movie-app/internal/core/services"
)

type AdminController struct {
	Config  config.Config
	service services.AdminServices
}

func NewAdminController(service services.AdminServices, cfg config.Config) *AdminController {
	return &AdminController{service: service, Config: cfg}
}
