package db

import (
	"aion/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func init() {
	Connect()
}
