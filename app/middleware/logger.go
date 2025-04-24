package middleware

import (
	"time"

	"github.com/Riyoukou/odyssey/pkg/logger"
	"github.com/gin-gonic/gin"
)

// LogRequest 记录 HTTP 请求日志
func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求的开始时间
		start := time.Now()

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// 记录请求的基本信息，使用格式化的日志输出
		logger.Infof("[Request] %s %s | IP: %s | UA: %s",
			method, path, clientIP, userAgent)

		// 调用后续的处理函数
		c.Next()

		// 记录响应信息
		latency := time.Since(start)
		status := c.Writer.Status()

		// 根据状态码选择日志级别
		if status >= 400 {
			logger.Errorf("[Response] %s %s | Status: %d | Latency: %v",
				method, path, status, latency)
		} else {
			logger.Infof("[Response] %s %s | Status: %d | Latency: %v",
				method, path, status, latency)
		}
	}
}
