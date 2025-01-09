package infra

import (
	"context"
	"errors"
	"fmt"
	"movie-app/internal/config"
	"movie-app/utils/mysql"

	"gorm.io/gorm"
)

type Infrastructure struct {
	Config         config.Config
	GormConnection *gorm.DB
	Ctx            context.Context
	ErrorCh        chan error
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

	return infra, nil, func() {
		fmt.Println("Infra clean up")
		for _, c := range cleanup {
			c()
		}
	}
}
