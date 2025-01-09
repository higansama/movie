package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	router := gin.Default()
	return &App{Router: router}
}

func (a *App) InitializeRoutes() {
	// Initialize routes here
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}
