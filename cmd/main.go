package main

import (
	"context"
	"movie-app/internal/config"
	"movie-app/utils/infra"
	logger "movie-app/utils/logger"

	"github.com/gin-contrib/cors"
	"github.com/rs/zerolog/log"

	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a context that cancels on interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load configuration
	err := config.LoadConfig("")
	if err != nil {
		panic(err)
	}
	logger.InitLogger(config.Cfg)

	// init infra
	infra := infra.NewInfrastructure(config.Cfg)
	infra, err, infraCleanUp := infra.InitInfrastructure(ctx)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	defer infraCleanUp()

	// Initialize the Gin router
	engine := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true                                             // Mengizinkan semua asal permintaan
	corsCfg.AllowCredentials = true                                            // Mengizinkan pengiriman kredensial (cookie, authorization header)
	corsCfg.AllowHeaders = []string{"*"}                                       // Mengizinkan semua header
	corsCfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Mengizinkan semua metode
	engine.Use(cors.New(corsCfg))                                              // Allow all headers

	// Initialize the application

	go func() {
		// Run server on separate go routine for Go < 1.18 to make sure another
		// deffered func in main working.
		err = engine.Run(config.Cfg.AppAttribute.Host + ":" + config.Cfg.AppAttribute.Port)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	for {
		select {
		case err = <-infra.ErrorCh:
			log.Panic().Err(err).Send()
			return
		case <-ctx.Done():
			log.Info().Msg("Server exiting")
			return
		}
	}
}
