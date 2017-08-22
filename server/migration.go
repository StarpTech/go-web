package server

import "github.com/jinzhu/gorm"
import "github.com/starptech/go-web/models"

func Up(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
