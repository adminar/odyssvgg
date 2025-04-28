package repository

import (
	"errors"
	"time"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserLogin(name, password string) (*model.UserTable, error) {
	var user model.UserTable
	if err := DB.Where("name = ?", name).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户名错误")
		}
		return nil, errors.New("数据库查询错误")
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	// 生成 Token 并存入数据库
	user.Token = utils.GenerateAuthToken(user.Name)
	user.LastLogin = time.Now().Format("2006-01-02 15:04:05")
	if err := DB.Save(&user).Error; err != nil {
		return nil, errors.New("数据库保存 Token 失败")
	}
	return &user, nil
}
