package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseData ...
type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//SuccessResponse ...
func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.GetMsg(),
		Data:    data,
	})
}

//ErrorResponse 根据错误状态码返回响应
func ErrorResponse(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.GetMsg(),
		Data:    nil,
	}
	ctx.JSON(int(c), rd)
}

//ErrorWithMsgResponse 自定义错误响应
func ErrorWithMsgResponse(ctx *gin.Context, code MyCode, errMsg string) {
	ctx.JSON(http.StatusBadRequest, &ResponseData{
		Code:    code,
		Message: code.GetMsg(),
		Data:    nil,
	})
}
