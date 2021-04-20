package common

import (
	"Gin/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt: json web token，加密协议token生成器

var jwtKey = []byte("BeiHang_software")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	// 过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Samshui",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
