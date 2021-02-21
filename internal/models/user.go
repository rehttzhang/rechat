package models

//User 用户
type User struct {
	Model
	Username     string `json:"username" gorm:"type:varchar(100);not null;comment:用户名"`
	PasswordHash string `json:"passwordHash" gorm:"type:varchar(100);not null;comment:密码"`
	//Email        string `json:"email" gorm:"type:varchar(100);comment:注册邮箱"`
	//Phone        string `json:"phone" gorm:"type:varchar(100);comment:注册手机号"`
}

//UserInfo 注册登录用模块
