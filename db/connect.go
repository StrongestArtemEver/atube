package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connect struct {
	DB *gorm.DB
}

func (c *Connect) ConnectToDb() {
	dsn := "host=localhost user=artem password=Dinis1981@ dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}
	c.DB = db
	log.Println("#####")
	log.Println("##### Подключение к базе данных успешно выполнено #####")
	log.Println("#####")
}
