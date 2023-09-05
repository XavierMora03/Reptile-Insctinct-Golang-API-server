package controllers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := "host=localhost user=xavier password=storepass dbname=reptileinstinct " //port=9920" // sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=storeadmin password=storepass dbname=reptileinstinct " //port=9920" // sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	log.Default().Output(0, "DATABASE CONECTED SUCCESSFULLY")

	// err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}

// */
