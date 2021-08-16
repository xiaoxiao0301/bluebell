package controller

import (
	"go_web/web_app/dict"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义接口返回的格式信息
type Response struct {
	Code    dict.ResponseCode `json:"code"`
	Message interface{}       `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
}

// ReturnOk 成功返回
func ReturnOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    dict.CodeSuccess,
		Message: dict.CodeSuccess.Message(),
		Data:    data,
	})
}

// ReturnErr 错误信息返回
func ReturnErr(ctx *gin.Context, code dict.ResponseCode) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	})
}

// ReturnErrWithMessage 自定义错误信息返回
func ReturnErrWithMessage(ctx *gin.Context, code dict.ResponseCode, message interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
