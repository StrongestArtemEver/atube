package routes

import (
	"atube/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)


func RegisterVideoRoutes(router *gin.Engine, db *gorm.DB) {
	videoController, err := controllers.NewVideoController(db)
	if err != nil {
		log.Fatalf("Не удалось создать VideoController: %v", err)
	}

	router.POST("/video/upload", videoController.UploadVideo)
	router.GET("/videos", videoController.GetAllVideo)
}
