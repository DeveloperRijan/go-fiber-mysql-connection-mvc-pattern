package config

import (
	"GoFiberMySQLApp/database"
	"fmt"

	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

//public func initiate all configs
func InitConfigs() (connectionStatus string) {
	//call the db connection func
	msg := dbConnection()
	//fmt.Println(msg)
	return msg
}

//private function -Config DB Connection
func dbConnection() (msg string) {
	config := &database.Config{
		HOST:    "localhost",    //os.Getenv("HOST"),
		PORT:    "3306",         //os.Getenv("PORT"),
		USER:    "root",         //os.Getenv("USER"),
		PASS:    "",             //os.Getenv("PASS"),
		DB_NAME: "gofiber_demo", //os.Getenv("DB_NAME"),
	}

	db, err := database.Connect(config)
	if err != nil {
		return fmt.Sprintf("Could not connect to DB, Error : %s", err)
	}

	c := &Config{}
	c.DB = db
	return "DB Connected Successfully"
}
