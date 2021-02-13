package utils

import "github.com/gin-gonic/gin"

//ResponseData ...
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//SuccessResponse ...
func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, &ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
