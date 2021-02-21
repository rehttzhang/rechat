package middleware

import (
	"rechat/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
基于JWT实现登录认证中间件
对于需要登录才能访问的API来说，该中间件需要从请求头中获取JWT Token
如果没有Token --> /login
如果token过期 --> /login
从JWT中解析UserID字段 --> 根据userID就能从数据库中查询到当前请求的用户是谁
*/

//JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//从请求头中获取token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorWithMsgResponse(c, utils.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}

		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.ErrorWithMsgResponse(c, utils.CodeInvalidToken, "Token格式有误")
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenString，使用定义好的解析JwT函数解析它
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.ErrorResponse(c, utils.CodeInvalidToken)
			zap.L().Error("invalid JWT token", zap.Error(err))
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求上下文c中
		c.Set("userID", mc.UserID)
		c.Next() //后续的处理函数可以通过c.Get("userID")来获取当前请求的用户信息
	}
}
