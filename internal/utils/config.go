package utils

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//NewConfig 配置初始化器
func NewConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/internal/configs/")
	viper.AddConfigPath("./configs")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	return viper.ReadInConfig()
}
