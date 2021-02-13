package v1

import (
	"net/http"
	"rechat/internal/models"
	"rechat/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//UserRouter ...
type UserRouter struct {
	UserSrv service.UserService
}

//Register 注册
func (u *UserRouter) Register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		zap.L().Error("Error binding user for register: ", zap.Any("err", err))
		return
	}
	//先判断用户名是否已存在
	if ok := u.UserSrv.ExistUserByName(*user.Username); ok {
		zap.L().Info("用户已存在！")
		c.JSON(http.StatusFound, gin.H{
			"code":    302,
			"message": "用户已存在",
		})
	}
	//如果不存在再将用户密码加密
	passwordHash, err := service.HashPassword(string(*user.PasswordHash))
	if err != nil {
		zap.L().Error("密码加密失败: ", zap.Any("err", err))
		return
	}
	//再创建用户，存入数据库
	if err := u.UserSrv.Register(models.User{Username: user.Username, PasswordHash: &passwordHash}); err != nil {
		zap.L().Error("用户创建失败: ", zap.Any("err", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "注册成功！",
		"data":    nil,
	})
}

//Login 登录
func (u *UserRouter) Login(c *gin.Context) {}
