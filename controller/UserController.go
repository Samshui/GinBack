package controller

import (
	"Gin/common"
	"Gin/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(c *gin.Context) {

	// 获取数据库
	DB := common.GetDB()

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
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "telephone is exist"})
		return
	}

	// 新建用户
	newUser := model.User{
		Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 返回数据

	c.JSON(200, gin.H{
		"message": "register success!",
	})
}

// 查询手机号
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// RandomString 随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#$%&*_")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
