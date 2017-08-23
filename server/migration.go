package server

import (
	"github.com/starptech/go-web/models"
)

type Migration struct{}

func (m *Migration) Up() {
	db.AutoMigrate(&models.User{})

	// test data
	model := models.User{Name: "thorsten"}
	db.Create(&model)
}
