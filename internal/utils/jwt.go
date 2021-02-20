package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
//jwt-go自带的jwt.StandardClaims只包含了官方字段
//这里需要额外记录一个UserID字段，所以需要自定义结构体
//想要保存更多信息都可以添加到此结构体中
type MyClaims struct {
	UserID uint `json:"user_id"`
	//Username string `json:"username"`
	jwt.StandardClaims
}

var mySecret = []byte(viper.GetString("token.my_secret"))

var tokenExpireTime = time.Minute * time.Duration(viper.GetInt("token.expire_time"))

//var tokenRefreshTime = time.Hour * time.Duration(viper.GetInt("token.refresh_time"))

//GenToken 生成token
func GenToken(userID uint) (string, error) {
	//创建一个自己的声明
	c := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireTime).Unix(), //过期时间
			Issuer:    "dashu",                                //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获取完整的编码后的字符串token
	token, err := tokenClaims.SignedString(mySecret)
	if err != nil {
		zap.L().Error("Error getting token with secret", zap.Any("err", err))
		return "", err
	}
	return token, nil
}

//ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析鉴权声明
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if tokenClaims != nil {
		//从tokenClaims中获取Claims对象，并使用断言将对象转换为需要的*MyClaims
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
