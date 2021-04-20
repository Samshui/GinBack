package main

import (
	"Gin/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close() // 延迟关闭

	r := gin.Default()
	r = CollectRouter(r)

	panic(r.Run(":8010"))
}
