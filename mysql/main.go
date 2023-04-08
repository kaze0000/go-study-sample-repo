// https://nishinatoshiharu.com/connect-go-database/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 波線でてるけど、使えている
)

func main() {
    // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
    dbconf := "webuser:webpass@tcp(localhost:3306)/go_mysql8_development?charset=utf8mb4" // データベース接続成功
    db, err := sql.Open("mysql", dbconf)

    // 接続が終了したらクローズする
    defer db.Close()

    if err != nil {
        fmt.Println(err.Error())
    }

    err = db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }
}
