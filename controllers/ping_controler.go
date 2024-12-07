package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"atube/services"
)

type PingController struct{
	pingService services.PingService
}
// @Summary      Ping endpoint
// @Description  Проверяет работоспособность сервера
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func(pc *PingController) Ping(context *gin.Context){
	message := pc.pingService.GetPingMessage()
	context.JSON(http.StatusOK,gin.H{
		"message" : message,
	})
}