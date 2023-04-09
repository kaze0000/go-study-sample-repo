// https://rightcode.co.jp/blog/information-technology/golang-clean-architecture-rest-api-syain

package main

import (
	"context"
	"fmt"
	"os"

	"echo-rest-clean/drivers"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}


type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Respose struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Status  string    `json:"status"`
}

type Users *[]User

func main() {
	// e := echo.New()
	// e.GET("/users/:name", getUserName) // http://localhost:1323/users/xxx
	// e.GET("/show", show) // http://localhost:1323/show?team=example&member=sasanori
	// e.POST("/save", save)

	// e.POST("/users", saveUser) //jsonで返す
	// e.POST("/send", sendMessage) //jsonで受け取って、jsonで返す
	// e.Logger.Fatal(e.Start(":1323"))

	ctx := context.Background()
	userDriver, err := drivers.IndirectUserDriver(ctx)
	if err != nil {
		fmt.Printf("failed to create UserDriver: %s\n", err)
		os.Exit(2)
	}
	userDriver.ServeUsers(ctx, ":8000")
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, name)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(200, fmt.Sprintf("team: %s, member: %s", team, member))
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(200, fmt.Sprintf("name: %s, email: %s", name, email))
	// http --form POST http://localhost:1323/save name='kazuya' email='aa@aa.com'
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(200, u)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}
	r := new(Respose)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(200, r)
}

