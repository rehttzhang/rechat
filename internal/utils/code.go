package utils

//MyCode ...
type MyCode int64

const (
	CodeSuccess         MyCode = 200
	CodeInvalidParams   MyCode = 10001
	CodeUserExist       MyCode = 10002
	CodeUserNotExist    MyCode = 10003
	CodeInvalidPassword MyCode = 10004
	CodeServerBusy      MyCode = 10005

	CodeInvalidToken      MyCode = 10006
	CodeInvalidAuthFormat MyCode = 10007
	CodeNotLogin          MyCode = 10008
)

var msgFlags = map[MyCode]string{
	CodeSuccess:           "success",
	CodeInvalidParams:     "请求参数错误",
	CodeUserExist:         "用户名重复",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeServerBusy:        "服务繁忙",
	CodeInvalidToken:      "无效的token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
}

//GetMsg 根据code获得错误信息
func (c MyCode) GetMsg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
