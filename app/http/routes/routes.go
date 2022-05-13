package routes

import (
	"com/denis/smarthome/app/http/controllers"
	"github.com/gin-gonic/gin"
)

// New 初始化路由
func New(e *gin.Engine) {
	contextPath := e.Group("/ladon")

	contextPath.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	contextPath.POST("/collect_photo", controllers.CollectPhoto)
}
