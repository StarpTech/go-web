package store

import "github.com/starptech/go-web/models"

type User interface {
	First(m *models.User) error
	Find(m *[]models.User) error
	Create(m *models.User) error
	Ping() error
}
