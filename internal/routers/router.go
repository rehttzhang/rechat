package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewRouter 路由初始化器
func NewRouter() *gin.Engine {
	r := gin.Default()

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
