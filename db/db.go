package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/starptech/go-web/config"
)

var db *gorm.DB

func init() {
	config := config.GetConfig()
	database, err := gorm.Open(config.Dialect, config.ConnectionString)

	if err != nil {
		fmt.Printf("could not connect to database %q", err)
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}
