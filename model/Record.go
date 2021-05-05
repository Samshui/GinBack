package model

import (
	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model

	SelectedExperiment Experiment `gorm:"ForeignKey:ExperimentID"`
	Selector           User       `gorm:"ForeignKey:UserID"`
	Date               string     `gorm:"type:varchar(255)"`
	Site               int        `gorm:"type:int;default:-1;not null"`
	Time               int        `gorm:"type:int;default:-1;not null"`

	ExperimentID string
	UserID       string
}
