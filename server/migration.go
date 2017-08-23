package server

import (
	"github.com/jinzhu/gorm"
	"github.com/starptech/go-web/models"
)

type Migration struct {
	Db *gorm.DB
}

func (m *Migration) Up() {
	m.Db.AutoMigrate(&models.User{})

	// test data
	model := models.User{Name: "thorsten"}
	m.Db.Create(&model)
}
