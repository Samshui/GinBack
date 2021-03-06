package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null"`
	StudentID string `gorm:"varchar(8);not null"`
	Password  string `gorm:"size:255;not null"`
	Status    int    `gorm:"int;not null;default: 1"`
}
