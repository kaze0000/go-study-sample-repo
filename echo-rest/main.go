// https://rightcode.co.jp/blog/information-technology/golang-clean-architecture-rest-api-syain

package main

import (
	"fmt"

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
	Status  string `json:"status"`
}

type Users *[]User

func main() {
	e := echo.New()
	e.GET("/users/:name", getUserName) // http http://localhost:8080/users/yoshi
	e.GET("/show", show) // http http://localhost:8080/show?team=example&member=yoshi
	e.POST("/save", save) // http --form POST http://localhost:8080/save name='yoshi' email='aa@aa.com'

	// jsonで返す場合
	e.POST("/users", saveUser) // http --form POST http://localhost:8080/users name='yoshi' email='aa@aa.com'
	e.POST("/send", sendMessage) // http --form POST http://localhost:8080/send name='yoshi' email='aa@aa.com' message='test message'
	e.Logger.Fatal(e.Start(":8080"))
}

// 以下handler/handler.goなどに切り出してもいい
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
}

func saveUser(c echo.Context) error {
	u := new(User) // uの中身、&{}
	if err := c.Bind(u); err != nil { //HTTPリクエストボディに含まれるJSONデータをパースし、uにマッピングする
		return err
	}
	return c.JSON(200, u)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}
	// MessageオブジェクトをResponseオブジェクトに変換
	r := new(Respose)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(200, r)
}

