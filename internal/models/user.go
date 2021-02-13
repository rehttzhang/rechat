package models

//User 用户
type User struct {
	Model
	Username     *string `json:"username" gorm:"not null;comment:用户名"`
	PasswordHash *[]byte `json:"passwordHash" gorm:"not null;comment:密码"`
}
