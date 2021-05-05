package common

import (
	"Gin/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// InitDB 连接
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("数据库连接失败，错误信息：" + err.Error())
	}

	// 自动创建表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Experiment{})
	db.AutoMigrate(&model.Record{})

	DB = db
	return db
}

// GetDB 获取DB
func GetDB() *gorm.DB {
	return DB
}
