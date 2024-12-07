package main

// @title API Documentation
// @version 1.0
// @description This is a sample server.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1

import _ "atube/docs"
import (
	"atube/db"
	"atube/models"
	"atube/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func main() {

	connection := db.Connect{}
	connection.ConnectToDb()

	err := connection.DB.AutoMigrate(
		&models.User{},     
		&models.Video{},
		&models.Bookmark{},
		&models.Comment{},
		&models.Like{},
		&models.VideoView{},

	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
	log.Println("Миграция выполнена успешно!")

	router := gin.Default()
	routes.RegisterPingRoutes(router)
	routes.RegisterVideoRoutes(router, connection.DB)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
