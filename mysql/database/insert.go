package database

import (
	"database/sql"
)

func InsertData(db *sql.DB, name, email string) error {
	query := `INSERT INTO users(name, email) VALUES (?,?)`
	stmt, err := db.Prepare(query) // Prepare関数は、SQLインジェクション攻撃を防ぐためにも推奨される方法
																 // SQLクエリのパラメータをプレースホルダ（?）で表現し、後から値をバインドすることで、不正なSQLコードの埋め込みを防ぐ
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email)
	if err != nil {
		return err
	}

	return nil
}
