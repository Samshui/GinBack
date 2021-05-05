package model

import "github.com/jinzhu/gorm"

type Experiment struct {
	gorm.Model
	Eid         string `gorm:"type:varchar(7);not null"`
	ELabel      int    `gorm:"type:int;default:0;not null"`
	Lab         string `gorm:"size:1000;not null"`
	Name        string `gorm:"size:1000;not null"`
	TimeMorning int    `gorm:"type:tinyint(1);default:0;not null"`
	TimeNoon    int    `gorm:"type:tinyint(1);default:0;not null"`
	TimeEvening int    `gorm:"type:tinyint(1);default:0;not null"`
	SiteSize    int    `gorm:"type:int;default:30;not null"`
}
