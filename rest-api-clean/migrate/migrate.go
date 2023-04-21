package main

import (
	"rest-api-clean/db"
	"rest-api-clean/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
