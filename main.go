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
	routes.RegisterVideoRoutes(router,connection.DB)

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
}
}
