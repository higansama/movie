package migration

import (
	"context"
	"fmt"
	"movie-app/internal/config"
	"movie-app/internal/models"
	"movie-app/utils/auth"
	"movie-app/utils/infra"
	"movie-app/utils/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate() {
	// Create a context that cancels on interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load configuration
	err := config.LoadConfig("")
	if err != nil {
		fmt.Println("error disini 3", err.Error())
		panic(err)
	}
	logger.InitLogger(config.Cfg)
	config.Cfg.MySqlConfig.ShowLog = true
	// init infra
	infra := infra.NewInfrastructure(config.Cfg)
	infra, err, infraCleanUp := infra.InitInfrastructure(ctx)
	if err != nil {
		fmt.Println("error disini  4")
		panic(err)
	}
	defer infraCleanUp()
	fmt.Println("migrating . . .")
	// Perform migration
	ModelsToMigrate(infra.GormConnection)
}

func ModelsToMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Movie{}, &models.Casting{}, &models.Genre{})
	if err != nil {
		fmt.Println("error migrate models movie ", err.Error())
	}

	err = AlterTable(db, "users", models.User{})
	if err != nil {
		panic("Migration User failed: " + err.Error())
	}

	err = AlterTable(db, "actors", models.Actor{})
	if err != nil {
		panic("Migration Actor failed")
	}

	err = db.AutoMigrate(&models.Casting{}, &models.Actor{})
	if err != nil {
		panic("Migration Casting failed")
	}

	err = db.AutoMigrate(&models.Genre{})
	if err != nil {
		panic("Migration Casting failed")
	}

	err = AlterTable(db, "voting_histories", &models.VotingHistory{})
	if err != nil {
		panic("alter voting_histories failed")
	}

	err = db.AutoMigrate(&models.WathcingHistory{})
	if err != nil {
		panic("Migration Casting failed")
	}

}
func AlterTable(db *gorm.DB, tableName string, newModel interface{}) error {
	if db.Migrator().HasTable(tableName) {
		err := db.Migrator().AutoMigrate(newModel)
		if err != nil {
			return err
		}
	} else {
		err := db.AutoMigrate(newModel)
		if err != nil {
			return err
		}
	}
	return nil
}

func SeedActors() {
	// Create a context that cancels on interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load configuration
	err := config.LoadConfig("")
	if err != nil {
		panic(err)
	}
	logger.InitLogger(config.Cfg)
	config.Cfg.MySqlConfig.ShowLog = true
	// init infra
	infra := infra.NewInfrastructure(config.Cfg)
	infra, err, infraCleanUp := infra.InitInfrastructure(ctx)
	if err != nil {
		panic(err)
	}
	defer infraCleanUp()

	actors := []models.Actor{
		{ID: uuid.New(), Name: "Leonardo DiCaprio"},
		{ID: uuid.New(), Name: "Robert Downey Jr."},
		{ID: uuid.New(), Name: "Scarlett Johansson"},
		{ID: uuid.New(), Name: "Brad Pitt"},
		{ID: uuid.New(), Name: "Johnny Depp"},
		{ID: uuid.New(), Name: "Tom Hanks"},
		{ID: uuid.New(), Name: "Meryl Streep"},
		{ID: uuid.New(), Name: "Natalie Portman"},
		{ID: uuid.New(), Name: "Denzel Washington"},
		{ID: uuid.New(), Name: "Morgan Freeman"},
		{ID: uuid.New(), Name: "Christian Bale"},
		{ID: uuid.New(), Name: "Emma Stone"},
		{ID: uuid.New(), Name: "Anne Hathaway"},
		{ID: uuid.New(), Name: "Matt Damon"},
		{ID: uuid.New(), Name: "Angelina Jolie"},
		{ID: uuid.New(), Name: "Chris Hemsworth"},
		{ID: uuid.New(), Name: "Chris Evans"},
		{ID: uuid.New(), Name: "Chris Pratt"},
		{ID: uuid.New(), Name: "Jennifer Lawrence"},
		{ID: uuid.New(), Name: "Daniel Radcliffe"},
		{ID: uuid.New(), Name: "Emma Watson"},
		{ID: uuid.New(), Name: "Hugh Jackman"},
		{ID: uuid.New(), Name: "Ryan Reynolds"},
		{ID: uuid.New(), Name: "Gal Gadot"},
		{ID: uuid.New(), Name: "Keanu Reeves"},
		{ID: uuid.New(), Name: "Benedict Cumberbatch"},
		{ID: uuid.New(), Name: "Mark Ruffalo"},
		{ID: uuid.New(), Name: "Tom Holland"},
		{ID: uuid.New(), Name: "Zendaya"},
		{ID: uuid.New(), Name: "Joaquin Phoenix"},
		{ID: uuid.New(), Name: "Robert Pattinson"},
		{ID: uuid.New(), Name: "Timothée Chalamet"},
		{ID: uuid.New(), Name: "Florence Pugh"},
		{ID: uuid.New(), Name: "Millie Bobby Brown"},
	}

	err = infra.GormConnection.CreateInBatches(actors, len(actors)).Error
	if err != nil {
		fmt.Println("err ", err.Error())
	}

}

func SeedGenre() {
	// Create a context that cancels on interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load configuration
	err := config.LoadConfig("")
	if err != nil {
		panic(err)
	}
	logger.InitLogger(config.Cfg)
	config.Cfg.MySqlConfig.ShowLog = true
	// init infra
	infra := infra.NewInfrastructure(config.Cfg)
	infra, err, infraCleanUp := infra.InitInfrastructure(ctx)
	if err != nil {
		panic(err)
	}
	defer infraCleanUp()

	actors := []models.Genre{
		{Title: "Action"},
		{Title: "Adventure"},
		{Title: "Comedy"},
		{Title: "Crime"},
		{Title: "Drama"},
		{Title: "Fantasy"},
		{Title: "Historical"},
		{Title: "Horror"},
		{Title: "Mystery"},
		{Title: "Romance"},
		{Title: "Science Fiction"},
		{Title: "Thriller"},
		{Title: "War"},
		{Title: "Western"},
		{Title: "Animation"},
		{Title: "Biography"},
		{Title: "Documentary"},
		{Title: "Family"},
		{Title: "Musical"},
		{Title: "Sport"},
	}

	err = infra.GormConnection.CreateInBatches(actors, len(actors)).Error
	if err != nil {
		fmt.Println("err ", err.Error())
	}

}

func CreateAdmin(username, password string) error {
	salt := auth.GenerateSalt()
	p := auth.GeneratePassword(salt, password)
	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Password: p,
		Salt:     salt,
		Role:     "admin",
	}

	// Create a context that cancels on interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load configuration
	err := config.LoadConfig("")
	if err != nil {
		panic(err)
	}
	logger.InitLogger(config.Cfg)
	config.Cfg.MySqlConfig.ShowLog = true
	// init infra
	infra := infra.NewInfrastructure(config.Cfg)
	infra, err, infraCleanUp := infra.InitInfrastructure(ctx)
	if err != nil {
		panic(err)
	}
	defer infraCleanUp()

	err = infra.GormConnection.Create(user).Error
	if err != nil {
		fmt.Println("err ", err.Error())
	}

	return nil
}
