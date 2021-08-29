package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	dsn := ""
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database with err: " + err.Error())
	}

	return connection
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("Couldn't close connection: " + err.Error())
	}
}