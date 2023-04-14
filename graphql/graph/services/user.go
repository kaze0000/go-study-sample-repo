package services

import (
	"context"
	"database/sql"
	"gql-server/graph/db"
	"gql-server/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct { //非公開型の構造体なので、他のパッケージからはアクセスできない
	exec boil.ContextExecutor //ContextExecutorインターフェイスを実装したオブジェクトを格納するfield
}

type ContextExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func converUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)
	// SQLBoilerで生成されたORMコードを呼び出して得られるのは、SQLBoilerコマンドにて自動生成されたdb.User型
	// なのでリゾルバで使うためにはmodel.User型にconvertする必要がある

	// 2. エラー処理
	if err != nil {
		return nil, err
	}

	// 3. 戻り値の*model.User型を作る
	return converUser(user), nil
}
