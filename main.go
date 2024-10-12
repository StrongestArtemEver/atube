package main

import (
	"log"
	"atube/routes"
	"github.com/gin-gonic/gin"
	"atube/db"
)

func main() {
	connection := db.Connect{}
	connection.ConnectToDb()

	router := gin.Default()

	routes.RegisterPingRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
}
}
