package models

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	Client *redis.Client
)

//NewRedis redis数据库初始化器
func NewRedis() error {
	//
	Client = redis.NewClient(&redis.Options{
		Addr:     ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"), //默认db为0
	})

	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
