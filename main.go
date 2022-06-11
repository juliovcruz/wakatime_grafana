package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main.go/api"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5555/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can not be connect to database")
	}

	api.Api(db)
}
