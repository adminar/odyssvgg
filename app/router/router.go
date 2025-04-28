package router

import (
	"github.com/Riyoukou/odyssey/app/controller/cicd"
	"github.com/Riyoukou/odyssey/app/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine) {

	// 使用请求日志中间件
	r.Use(
		middleware.LogRequest(),
		middleware.Cors(),
	)

	// 健康检查
	r.GET("/man", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "what can i say! mamba out!",
		})
	})

	_user := r.Group("/user")
	{
		_user.POST("login", cicd.HandleUserLogin)
		_user.POST("register", cicd.HandleUserRegister)
	}

	_cicd := r.Group("/cicd")
	{
		_cicd.GET("fetch/repo/:type", cicd.HandleCICDFetchRepo)
		_cicd.GET("get/repo/:type", cicd.HandleCICDGetRepo)
		_cicd.POST("create/repo/:type", cicd.HandleCICDCreateRepo)
		_cicd.POST("update/repo/:type", cicd.HandleCICDUpdateRepo)
		_cicd.DELETE("delete/repo/:type/:id", cicd.HandleCICDDeleteRepo)
	}
}
