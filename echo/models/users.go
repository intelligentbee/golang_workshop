package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Email     string     `json:”email”`
	FirstName string     `json:”firstName”`
	LastName  string     `json:”lastName”`
}
