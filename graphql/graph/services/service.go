package services

import (
	"context"
	"gql-server/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type IServices interface {
	IUserService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func NewServices(exec boil.ContextExecutor) IServices {
	// ファクトリ関数という
	return &services{
		userService: &userService{exec: exec},
	}
}
