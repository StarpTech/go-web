package models

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string `sql:"type:varchar(30)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
