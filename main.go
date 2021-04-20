package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	defer db.Close() // 延迟关闭

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		// 检查数据
		if len(telephone) != 11 {
			// gin.H = map[string]interface{}
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "phone's len != 11"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "password's len < 6"})
			return
		}

		if len(name) == 0 {
			// 姓名为空，生成随机初始名称
			name = RandomString(10)
		}

		log.Println(name, telephone, password)

		// 查找数据库，手机号是否已经存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "telephone is exist"})
			return
		}

		// 新建用户
		newUser := User{
			Model:     gorm.Model{},
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		// 返回数据

		c.JSON(200, gin.H{
			"message": "register success!",
		})
	})

	panic(r.Run(":8010"))
}

// 查询手机号
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


// 随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#$%&*_")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string((result))
}

// 连接
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
	db.AutoMigrate(&User{})

	return db;
}
