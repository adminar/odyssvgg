package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "操作成功"
	}
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

// OkWithData 成功响应，只返回数据
func OkWithData(c *gin.Context, data interface{}) {
	Success(c, data, "")
}

// OkWithMessage 成功响应，只返回消息
func OkWithMessage(c *gin.Context, message string) {
	Success(c, nil, message)
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// FailWithMessage 失败响应，使用默认错误码
func FailWithMessage(c *gin.Context, message string) {
	Fail(c, -1, message)
}

// FailWithCodeAndMessage 失败响应，使用自定义错误码和消息
func FailWithCodeAndMessage(c *gin.Context, code int, message string) {
	Fail(c, code, message)
}

// FailWithError 失败响应，使用错误对象
func FailWithError(c *gin.Context, err error) {
	if err != nil {
		FailWithMessage(c, err.Error())
		return
	}
	FailWithMessage(c, "未知错误")
}
