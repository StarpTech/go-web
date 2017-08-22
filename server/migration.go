package server

import (
	"github.com/starptech/go-web/db"
	"github.com/starptech/go-web/models"
)

type Migration struct{}

func (m *Migration) Up() {
	database := db.GetDB()
	database.AutoMigrate(&models.User{})

	// test data
	model := models.User{Name: "thorsten"}
	database.Create(&model)
}
