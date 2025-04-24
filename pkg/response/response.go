package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 结构体定义 API 返回的标准格式
type Response struct {
	Status  int         `json:"status"`  // 状态码
	Message string      `json:"message"` // 消息
	Result  interface{} `json:"result"`  // 数据
}

// Success 成功响应
func Success(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Result:  result,
	})
}

// Error 错误响应
func Error(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: err.Error(),
		Result:  nil,
	})
}

func FailWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Result:  nil,
	})
}
func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "success",
		Result:  data,
	})
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "success",
		Result:  nil,
	})
}
