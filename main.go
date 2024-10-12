package main

import (
	"log"
	"atube/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := os.Getenv("DBCONNECTIONSTRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil){
		log.Fatal("##### No DB #####",err)
	}

	log.Println("##### DB OK #####",db)
	router := gin.Default()

	routes.RegisterPingRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
}
}
