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

	if err = database.DB().Ping(); err != nil {
		log.Fatalf("gorm: could not ping database %q", err)
	}

	return database
}
