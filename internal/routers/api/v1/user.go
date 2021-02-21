package v1

import (
	"net/http"
	"rechat/internal/models"
	"rechat/internal/service"
	"rechat/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//UserRouter ...
type UserRouter struct {
	UserSrv service.UserService
}

//Register 注册
func (u *UserRouter) Register(c *gin.Context) {
	userInfo := models.User{}
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		zap.L().Error("Error binding user for register: ", zap.Any("err", err))
		utils.ErrorResponse(c, utils.CodeInvalidParams)
	}
	//先判断用户名是否已存在
	if ok := u.UserSrv.ExistUserByName(userInfo.Username); ok {
		zap.L().Info("用户已存在！")
		// c.JSON(http.StatusFound, gin.H{
		// 	"code":    302,
		// 	"message": "用户已存在",
		// })
		utils.ErrorResponse(c, utils.CodeUserExist)
	}
	//如果不存在再将用户密码加密
	passwordHash, err := service.HashPassword(userInfo.PasswordHash)
	if err != nil {
		zap.L().Error("密码加密失败: ", zap.Any("err", err))
		return
	}
	//再创建用户，存入数据库
	if err := u.UserSrv.Register(models.User{Username: userInfo.Username, PasswordHash: string(passwordHash)}); err != nil {
		zap.L().Error("用户创建失败: ", zap.Any("err", err))
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "注册成功！")
}

//Login 登录
func (u *UserRouter) Login(c *gin.Context) {
	//1. 获取登录请求携带的用户名，密码数据
	var userInfo models.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		zap.L().Error("Error binding user when login", zap.Any("err", err))
		utils.ErrorResponse(c, utils.CodeInvalidParams)
	}
	//2. 根据用户名获取用户
	user, err := u.UserSrv.GetUserByName(userInfo.Username)
	if err != nil {
		zap.L().Error("Can't get user by username", zap.Any("err", err))
		return
	}
	//3. 去数据库校验用户密码是否和输入密码一致
	err = service.CheckPassword(user, userInfo.PasswordHash)
	if err != nil {
		zap.L().Error("Can't login in", zap.Any("err", err))
		utils.ErrorResponse(c, utils.CodeInvalidPassword)
	}
	//4. 生成JWT Token
	token, err := utils.GenToken(user.ID)
	if err != nil {
		zap.L().Error("Error getting token with loginUser", zap.Any("err", err))
		utils.ErrorResponse(c, utils.CodeInvalidToken)
	}
	//5. 返回响应
	utils.SuccessResponse(c, http.StatusOK, map[string]interface{}{"登录成功！": token})
}
