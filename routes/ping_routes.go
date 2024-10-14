package routes

import (
	"github.com/gin-gonic/gin"
	"atube/controllers"
)

func RegisterPingRoutes(router *gin.Engine) {
	pingController := controllers.PingController{}

	router.GET("/ping",pingController.Ping)
}