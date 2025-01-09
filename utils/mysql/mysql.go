package mysql

import (
	"fmt"
	"movie-app/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(cfg config.Config) (*gorm.DB, error, func()) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySqlConfig.User,
		cfg.MySqlConfig.Password,
		cfg.MySqlConfig.Host,
		cfg.MySqlConfig.Port,
		cfg.MySqlConfig.Database,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               uri, // data source name
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})
	if err != nil {
		return db, err, func() {}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return db, err, func() {}
	}

	return db, nil, func() {
		_ = sqlDB.Close()
	}
}
