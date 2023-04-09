// HTTPリクエストを受け取り、Clean Architectureに従ってビジネスロジックを実行し、レスポンスを返す役割を担当している

package controllers

import (
	"context"

	"echo-rest-clean/adapters/gateways"
	"echo-rest-clean/entities"
	"echo-rest-clean/usecases/ports"

	"github.com/labstack/echo"
)

type User interface {
	AddUser(ctx context.Context) func(c echo.Context) error
	GetUsers(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInputPort
type RepositoryFactory func(gateways.FirestoreClientFactory) ports.UserRepository

type UserController struct {
	// Factory: オブジェクトの生成処理を専用のメソッドやクラスにまとめること
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	clientFactory     gateways.FirestoreClientFactory
}

func NewUserController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory, clientFactory gateways.FirestoreClientFactory) User {
	return &UserController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
		clientFactory:     clientFactory,
	}
}

func (u *UserController) AddUser(ctx context.Context) func(c echo.Context) error { //戻り値として、func(c echo.Context) errorを返す
	// [役割]ビジネスロジックを実行
	return func(c echo.Context) error {
		user := new(entities.User)
		if err := c.Bind(user); err != nil {
			return err
		}

		return u.newInputPort(c).AddUser(ctx, user)
	}
}

func (u *UserController) GetUsers(ctx context.Context) func(c echo.Context) error {
	// [役割]ビジネスロジックを実行
	return func(c echo.Context) error {
		return u.newInputPort(c).GetUsers(ctx)
	}
}

func (u *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	outputPort := u.outputFactory(c)
	repository := u.repositoryFactory(u.clientFactory)
	return u.inputFactory(outputPort, repository)
}
