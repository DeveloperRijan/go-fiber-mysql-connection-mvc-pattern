package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	HOST    string
	PORT    string
	USER    string
	PASS    string
	DB_NAME string
}

func Connect(config *Config) (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.USER, config.PASS, config.HOST, config.PORT, config.DB_NAME)

	fmt.Println(fmt.Sprintf("DB Connection String: %s", dsn))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
