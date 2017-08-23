package server

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDB(dialect, connString string) *gorm.DB {
	database, err := gorm.Open(dialect, connString)

	if err != nil {
		log.Fatalf("gorm: could not connect to database %q", err)
	}

	return database
}
