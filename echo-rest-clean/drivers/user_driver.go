package drivers

import (
	"context"

	"echo-rest-clean/adapters/controllers"

	"github.com/labstack/echo"
)

type User interface {
	ServeUsers(ctx context.Context, address string)
}

type UserDriver struct {
	echo       *echo.Echo
	controller controllers.User
}

func NewUserDriver(echo *echo.Echo, controller controllers.User) User {
	return &UserDriver{
		echo:       echo,
		controller: controller,
	}
}

func (driver *UserDriver) ServeUsers(ctx context.Context, address string) {
	driver.echo.POST("/users", driver.controller.AddUser(ctx))
	driver.echo.GET("/users", driver.controller.GetUsers(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(address))
}
