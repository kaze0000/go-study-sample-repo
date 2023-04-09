// Interactorでは、ユーザーデータを受け取るInputPortの実体と
// portsにかいた、技術的要素を省いた1~3の手順を実装

package interactors

import (
	"context"

	"echo-rest-clean/entities"
	"echo-rest-clean/usecases/ports"
)

// UserInteractorを使って、アプリケーションのビジネスロジックを実装することで、プレゼンテーション層とデータアクセス層の両方と疎結合になる->拡張性とメンテナンス性が向上する

type UserInteractor struct {
	// [役割]ユースケース層のinteractorとして機能する
	// interactor: ビジネスロジックを実装し、プレゼンテーション層とデータアクセス層をつなぐ役割を担う
	// UserInteractor は ports.UserInputPortインターフェースを実装しており、ビジネスロジックをカプセル化する
	OutputPort ports.UserOutputPort
	Repository ports.UserRepository
}

func NewUserInputPort(outputPort ports.UserOutputPort, repository ports.UserRepository) ports.UserInputPort {
	// [役割]ports.UserInputPortインターフェースを実装したオブジェクトを返す
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

//--- UserInputPort interface をimplements
func (u *UserInteractor) AddUser(ctx context.Context, user *entities.User) error {
	// [役割]ユーザーを追加するためのビジネスロジックを実装
	// このメソッドは、データアクセス層の Repository を使用してユーザーを追加し、結果をプレゼンテーション層のOutputPortに出力する
	users, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}

func (u *UserInteractor) GetUsers(ctx context.Context) error {
	// [役割]ユーザーを追加するためのビジネスロジックを実装
	// このメソッドは、データアクセス層の Repository を使用してユーザーを取得し、結果をプレゼンテーション層のOutputPortに出力する
	users, err := u.Repository.GetUsers(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}
//---
