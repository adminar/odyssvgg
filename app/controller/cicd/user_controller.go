package cicd

import (
	"fmt"
	"net/http"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/app/repository"
	"github.com/Riyoukou/odyssey/pkg/logger"
	"github.com/Riyoukou/odyssey/pkg/response"
	"github.com/gin-gonic/gin"
)

func HandleUserLogin(c *gin.Context) {
	var req model.UserTable
	err := c.ShouldBind(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	userInfo, err := repository.UserLogin(req.Name, req.Password)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		logger.Error(err)
		return
	}

	fmt.Println(userInfo)

	response.Success(c, map[string]interface{}{
		"id":    userInfo.ID,
		"name":  userInfo.Name,
		"token": userInfo.Token,
		"role":  userInfo.Role,
	}, "登录成功")
}

func HandleUserRegister(c *gin.Context) {
	var req model.UserTable
	err := c.ShouldBind(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	err = repository.UserRegister(req.Name, req.Password, req.Email, req.Phone)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		logger.Error(err)
		return
	}
	response.Success(c, nil, "注册成功")
}
