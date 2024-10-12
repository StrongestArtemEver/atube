package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"atube/services"
)

type PingController struct{
	pingService services.PingService
}

func(pc *PingController) Ping(context *gin.Context){
	message := pc.pingService.GetPingMessage()
	context.JSON(http.StatusOK,gin.H{
		"message" : message,
	})
}