package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	err := godotenv.Load()
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") //dbではなく、localhost
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")
	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
			log.Fatal(err)
	}
	fmt.Println("connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
