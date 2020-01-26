package models

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string `sql:"type:varchar(30)"`
}
