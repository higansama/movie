package infra

import (
	"context"
	"errors"
	"fmt"
	"movie-app/internal/config"
	"movie-app/utils/middleware"
	"movie-app/utils/mysql"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Infrastructure struct {
	Config         config.Config
	GormConnection *gorm.DB
	Ctx            context.Context
	ErrorCh        chan error
	Middleware     Middleware
}

func NewInfrastructure(cfg config.Config) *Infrastructure {
	return &Infrastructure{Config: cfg}
}

func (infra *Infrastructure) InitInfrastructure(ctx context.Context) (*Infrastructure, error, func()) {
	var cleanup []func()
	infra.ErrorCh = make(chan error)

	infra.Ctx = ctx
	if infra.Config.AppAttribute.Name == "" {
		return nil, errors.New("application name is empty"), func() {}
	}

	mysqlConn, err, mysqlCleanUp := mysql.NewMysqlConnection(infra.Config)
	if err != nil {
		return nil, err, func() {}
	}
	infra.GormConnection = mysqlConn
	infra.Ctx = ctx
	cleanup = append(cleanup, mysqlCleanUp)

	Middleware, err := infra.SetupMiddleware()
	if err != nil {
		return nil, err, nil
	}
	infra.Middleware = Middleware

	return infra, nil, func() {
		fmt.Println("Infra clean up")
		for _, c := range cleanup {
			c()
		}
	}
}

type Middleware struct {
	AdminMiddleware func(ctx *gin.Context)
	UserMiddleware  func(ctx *gin.Context)
}

func (infra *Infrastructure) SetupMiddleware() (Middleware, error) {
	adminAuth, err := middleware.NewAdminMiddleware(infra.Config)
	if err != nil {
		return Middleware{}, err
	}

	userAuth, err := middleware.NewUserMiddleware(infra.Config)
	if err != nil {
		return Middleware{}, err
	}

	return Middleware{
		AdminMiddleware: adminAuth.Handle,
		UserMiddleware:  userAuth.Handle,
	}, nil
}
