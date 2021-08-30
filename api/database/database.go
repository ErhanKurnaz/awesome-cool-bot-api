package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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