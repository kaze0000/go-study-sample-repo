// Presentersでは、OutputPortの実体を実装する

package presenters

import (
	"log"

	"echo-rest-clean/entities"
	"echo-rest-clean/usecases/ports"

	"github.com/labstack/echo"
)

type UserPresenter struct {
	ctx echo.Context // リクエストに関する情報や状態を格納している
}

func NewUserOutputPort(ctx echo.Context) ports.UserOutputPort{
	// [役割]ports.UserOutputPortインターフェースを実装したオブジェクトを返す
	return &UserPresenter{ctx: ctx}
}

//--- UserOutputPort interface をimplements
func (p *UserPresenter) OutputUsers(users []*entities.User) error {
	// [役割]ユースケース層からの出力を処理する
	// このメソッドは、ユーザーのリストを受け取り、HTTPレスポンスとしてJSON形式で返す
	return p.ctx.JSON(200, users)
}

func (p *UserPresenter) OutputError(err error) error {
	// [役割]ユースケース層からの出力を処理する
	// このメソッドは、エラーを受け取り、ログに記録した後、HTTPレスポンスとしてJSON形式で返す
	log.Fatal(err)
	return p.ctx.JSON(500, err)
}
//---
