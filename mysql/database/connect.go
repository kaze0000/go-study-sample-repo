package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB{
	err := godotenv.Load() //こいつがいる！！
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") //dbではなく、localhost
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

	// 以下はubuntuでの設定
	// dbconf := "root@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4"

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
			fmt.Println(err.Error())
	}
	return db
}
