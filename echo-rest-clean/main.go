package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users/:name", getUserName) // http://localhost:1323/users/xxx
	e.GET("/show", show) // http://localhost:1323/show?team=example&member=sasanori
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":1323"))
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
