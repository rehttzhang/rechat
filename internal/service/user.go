package service

import (
	"errors"
	"rechat/internal/models"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserInterface ...
type UserInterface interface {
	Register(user models.User) error
	ExistUserByName(username string) bool
}

//UserService ...
type UserService struct {
	DB *gorm.DB
}

//ExistUserByName 根据用户名检验用户是否已存在
func (u *UserService) ExistUserByName(username string) bool {
	var user models.User
	result := u.DB.Table("user").Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 && result.Error == nil {
		return true
	}
	return false
}

//Register 用户注册
func (u *UserService) Register(user models.User) error {
	err := u.DB.Table("user").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//Login 登录
func (u *UserService) Login(user *models.User) (*models.User, error) {
	var userInfo models.User
	passwordHash, err := HashPassword(string(user.PasswordHash))
	if err != nil {
		return nil, err
	}
	err = u.DB.Table("user").Where("username = ? AND password_hash = ?", user.Username, &passwordHash).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

//UpdateUserPassword 更改密码
func (u *UserService) UpdateUserPassword(user *models.User, newPassword string) (*models.User, error) {
	var userInfo models.User
	passwordHash, err := HashPassword(newPassword)
	if err != nil {
		return nil, err
	}
	err = u.DB.Table("user").Where("username = ? AND password_hash = ?", user.Username, user.PasswordHash).
		First(&userInfo).
		Update("password", passwordHash).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

//CheckPassword 检验用户密码
func CheckPassword(user *models.User, password string) error {
	if user.PasswordHash != nil && len(user.PasswordHash) == 0 {
		return errors.New("密码未设置")
	}
	return bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
}

//HashPassword 将密码哈希加密
func HashPassword(password string) ([]byte, error) {
	//GenerateFromPassword 以给定的成本返回密码的bcrypt哈希值
	return bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("app.cost"))
}