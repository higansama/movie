package main

import (
	"context"
	"fmt"
	"movie-app/cmd/migration"
	"movie-app/internal/config"
	"movie-app/internal/entrypoint"
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
	if len(os.Args) < 2 {
		fmt.Println("expected 'runserver' or 'migrate' subcommands")
		os.Exit(1)
	}
	fmt.Println("os.Args[1] ", os.Args[1])

	switch os.Args[1] {
	case "runserver":
		runServer()
	case "migrate":
		fmt.Println("error disini  2")
		migration.Migrate()
	case "seed-actor":
		fmt.Println("seed actor")
		migration.SeedActors()
	case "seed-genre":
		migration.SeedGenre()
	case "createadmin":
		var username, password string
		fmt.Print("Enter username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)
		fmt.Printf("Creating admin with username: %s and password: %s\n", username, password)
		// Call a function to create the admin user with the provided username and password
		migration.CreateAdmin(username, password)

	default:
		fmt.Println("expected 'runserver' or 'migrate' subcommands")
		os.Exit(1)
	}
}

func runServer() {
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
	engine.Static("/assets", "./assets")

	// Initialize the application
	entrypoint.NewAdminModule(engine, config.Cfg, *infra)

	entrypoint.NewUserModule(engine, config.Cfg, *infra)

	// closing state
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
