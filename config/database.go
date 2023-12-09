package config

import (
	"log"
	"url-shortener/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	var err error
	dsn := "host=localhost user=postgres password=abiyyucakra99 dbname=urlshortener port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	Db.AutoMigrate(&models.Request{}, &models.RequestCount{})
	// return db
	// return nil
}
