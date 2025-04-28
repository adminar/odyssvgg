package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// Hash 密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 生成 AuthToken
func GenerateAuthToken(username string) string {
	data := username + ":" + GenerateSecret(8)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func GenerateSecret(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length] // 截取指定长度
}
