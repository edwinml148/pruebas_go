package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"`
	Task      []Task
}
