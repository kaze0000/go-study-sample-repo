// https://nishinatoshiharu.com/connect-go-database/
package main

import (
	"fmt"

	"github.com/kaze0000/go-study-youtube/mysql/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
    db := database.Connect()

    // 接続が終了したらクローズする
    defer db.Close()

    err := db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }

    err = database.CreateTable(db)
    if err != nil {
        fmt.Println("テーブル作成失敗")
        fmt.Println(err)
        return
    } else {
        fmt.Println("テーブル作成成功")
    }

    err = database.InsertData(db, "kaze", "ex@ex.com")
    if err != nil {
        fmt.Println("データ挿入失敗")
        return
    } else {
        fmt.Println("データ挿入成功")
    }
}
