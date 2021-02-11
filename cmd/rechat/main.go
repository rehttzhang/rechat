package main

import (
	"log"
	"net/http"
	"rechat/internal/routers"
	"rechat/internal/utils"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := utils.NewConfig(); err != nil {
		log.Println("Error reading config file")
		return
	}
	utils.NewLogger()
	zap.L().Info("config and logger all ready to use.")

	router := routers.NewRouter()

	server := &http.Server{
		Addr:         ":" + viper.GetString("server.port"),
		ReadTimeout:  time.Second * time.Duration(viper.GetInt("server.read_timeout")),
		WriteTimeout: time.Second * time.Duration(viper.GetInt("server.write_timeout")),
		Handler:      router,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error("Failed to connect server.", zap.Any("err", err))
	}
}
