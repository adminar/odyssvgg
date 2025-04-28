package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors 允许跨域请求的中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin") // 获取请求中的 Origin 头

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin) // 动态设置允许的来源|
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Token")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带 cookie
		// 处理 OPTIONS 预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
