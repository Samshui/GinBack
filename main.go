package main

import (
	"Gin/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// 初始化读取配置
	InitConfig()

	db := common.InitDB()
	defer db.Close() // 延迟关闭

	r := gin.Default()
	r = CollectRouter(r)

	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

func InitConfig() {
	workDir, _ := os.Getwd()

	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
	}
}
