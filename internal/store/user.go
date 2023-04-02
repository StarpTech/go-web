package store

import "github.com/Stanxxy/stan-go-web/internal/models"

type User interface {
	First(m *models.User) error
	Find(m *[]models.User) error
	Create(m *models.User) error
	Ping() error
}
