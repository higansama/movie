package mysql

import (
	"fmt"
	"movie-app/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(cfg config.Config) (*gorm.DB, error, func()) {
	fmt.Println("Creating MySQL connection...")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySqlConfig.User,
		cfg.MySqlConfig.Password,
		cfg.MySqlConfig.Host,
		cfg.MySqlConfig.Port,
		cfg.MySqlConfig.Database,
	)

	fmt.Printf("Connection URI: %s\n", uri)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               uri, // data source name
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening connection: %v\n", err)
		return db, err, func() {}
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting sqlDB: %v\n", err)
		return db, err, func() {}
	}

	fmt.Println("MySQL connection established successfully.")

	return db, nil, func() {
		fmt.Println("Closing MySQL connection...")
		_ = sqlDB.Close()
		fmt.Println("MySQL connection closed.")
	}
}
