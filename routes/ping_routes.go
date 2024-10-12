package routes

import (
	"github.com/gin-gonic/gin"
	"atube/controllers"
)

func RegisterPingRoutes(router *gin.Engine) {
	pingControlle := controllers.PingController{}

	router.GET("/ping",pingControlle.Ping)
}