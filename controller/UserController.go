package controller

import (
	"Gin/common"
	"Gin/dto"
	"Gin/model"
	"Gin/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/** routers **/

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
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号位数不对")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码过短")
		return
	}

	if len(name) == 0 {
		// 姓名为空，生成随机初始名称
		name = RandomString(10)
	}

	log.Println(name, telephone, password)

	// 查找数据库，手机号是否已经存在
	if IsTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号重复存在")
		return
	}

	// 新建用户（密码加密）
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "密码加密失败")
	}
	newUser := model.User{
		Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  string(encryptedPassword),
	}
	DB.Create(&newUser)

	// 返回数据
	response.Success(c, nil, "注册成功！")
}

func Login(c *gin.Context) {
	// 获取数据（手机号 + 密码）
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		// gin.H = map[string]interface{}
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机位数不对")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码过短")
		return
	}

	// 查询手机号是否存在
	db := common.GetDB()
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 用户密码判断（加密判断）
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "用户密码错误")
	}

	// 密码通过（发放token）
	token, tokenErr := common.ReleaseToken(user)
	if tokenErr != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", tokenErr)
	}

	// 返回
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	// c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

/** 工具方法 **/

// IsTelephoneExist 查询手机号
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
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
