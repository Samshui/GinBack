package model

import (
	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model

	ExperimentName string `gorm:"type:varchar(255)"`
	Lab            string `gorm:"type:varchar(255)"`
	Date           string `gorm:"type:varchar(255)"`
	Site           int    `gorm:"type:int;not null"`
	Time           int    `gorm:"type:int;not null"`
	ExperimentID   string
	UserID         string
}
