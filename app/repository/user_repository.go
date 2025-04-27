package repository

import (
	"errors"
	"time"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(req model.UserTable) (string, error) {
	var user model.UserTable
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("用户名错误")
		}
		return "", errors.New("数据库查询错误")
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	// 生成 Token 并存入数据库
	user.AuthToken = utils.GenerateAuthToken(user.Username)
	user.LastLogin = time.Now().Format("2006-01-02 15:04:05")
	if err := DB.Save(&user).Error; err != nil {
		return "", errors.New("数据库保存 Token 失败")
	}
	return user.AuthToken, nil
}
