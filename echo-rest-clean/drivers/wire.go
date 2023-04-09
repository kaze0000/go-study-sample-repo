//go:build wireinject
// +build wireinject

package drivers

import (
	"context"

	"echo-rest-clean/adapters/controllers"
	"echo-rest-clean/adapters/gateways"
	"echo-rest-clean/adapters/presenters"
	"echo-rest-clean/database"
	"echo-rest-clean/drivers"
	"echo-rest-clean/usecases/interactors"

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeUserDriver(ctx context.Context) (controllers.User, error) {
	wire.Build(NewFirestoreClientFactory, echo.New, NewOutputFactory, NewInputFactory, NewRepositoryFactory, controllers.NewUserController, NewUserDriver)
	return &drivers.UserDriver{}, nil
}

func NewFirestoreClientFactory() database.FirestoreClientFactory {
	return &database.MyFirestoreClientFactory{}
}

func NewOutputFactory() controllers.OutputFactory {
	return presenters.NewUserOutputPort
}

func NewInputFactory() controllers.InputFactory {
	return interactors.NewUserInputPort
}

func NewRepositoryFactory() controllers.RepositoryFactory {
	return gateways.NewUserRepository
}
