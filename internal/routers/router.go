package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//NewRouter 路由初始化器
func NewRouter() *gin.Engine {
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()

	//r.POST("/register", v1.UserRouter)
	//r.POST("/login", v1.UserRouter)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "go",
			})
		})
	}

	return r
}
