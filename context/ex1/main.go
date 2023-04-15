package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type UserData struct {
	ID   int
	Name string
}

// Contextを使用する関数
func FetchUserData(ctx context.Context, userID int) (*UserData, error) {
	// 仮定：データを取得するのに時間がかかる場合がある
	time.Sleep(3 * time.Second) // タイムアウトする
	// time.Sleep(2 * time.Second) タイムアウトしない

	// タイムアウトが発生したかどうかを確認
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// データを返す
		return &UserData{
			ID:   userID,
			Name: "John Doe",
		}, nil
	}

}

// HTTPハンドラー
func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストからコンテキストを取得(リクエスト間での情報のやりとりに使われる)
	ctx := r.Context()

	// タイムアウトを設定
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	  // 以下だと、リクエスト固有の情報が保持できない
		// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userID := 1

	userData, err := FetchUserData(ctx, userID)

	// エラーチェック
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ユーザーデータを返す
	fmt.Fprintf(w, "User: %v", userData)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
