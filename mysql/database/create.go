package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
