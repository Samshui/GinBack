package common

import (
	"Gin/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB 连接
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "GinBack"
	username := "root"
	password := "yhm012710"
	charset := "utf8"

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

	// 自动创建User表
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

// GetDB 获取DB
func GetDB() *gorm.DB {
	return DB
}
