package middleware

import (
	"Gin/common"
	"Gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authorization-header
		tokenString := context.GetHeader("Authorization")

		// 检查格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort() // 抛弃本次请求
			return
		}

		// 解析token
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort() // 抛弃本次请求
			return
		}

		// token通过，获取userID
		userID := claims.UserID
		DB := common.GetDB()

		var user model.User
		DB.First(&user, userID)

		// 用户不存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort() // 抛弃本次请求
			return
		}

		// 用户存在，user信息写入context
		context.Set("user", user)
		context.Next()
	}
}
