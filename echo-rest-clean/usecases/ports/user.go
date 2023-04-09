// firebaseにuserを登録する、addUserの実装
// 1.ユーザーデータを受け取り
// 2.Firestoreにユーザーデータを追加
// 3.全てのユーザーデータを返却する or エラーを返す

// Portではinterfaceのみを記述

package ports

import (
	"context"

	"echo-rest-clean/entities"
)

// 1.ユーザーデータを受け取り
type UserInputPort interface {
	// ユースケース層で使用される
	// [定義]アプリケーションの主要な操作（この場合はユーザーの追加と取得）を定義する
	// [役割]このインターフェースを実装した構造体は、ビジネスロジックを含むことが期待される
	AddUser(ctx context.Context, user *entities.User) error
	GetUsers(ctx context.Context) error
}

// 3.全てのユーザーデータを返却する or エラーを返す
type UserOutputPort interface {
	// プレゼンテーション層で使用される
	// [定義]ユースケース層からのデータを表示する方法を定義する
	// [役割]さまざまな種類のプレゼンテーション層がこのインターフェースを実装することができる
	OutputUsers([]*entities.User) error
	OutputError(error) error
}

// 2.Firestoreにユーザーデータを追加、全ユーザーを返す
type UserRepository interface {
	// データアクセス層で使用される
	// [定義]データベースやファイルなどの永続化されたデータを読み書きする操作を定義する
	// [役割]このインターフェースを実装した構造体は、データストレージとのやり取りを担当する
	// usecase層がデータアクセス層に依存しないようにするためのinterface
	AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error)
	GetUsers(ctx context.Context) ([]*entities.User, error)
}
