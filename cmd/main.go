package main

import (
	"github.com/Random7-JF/templ/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	index := handlers.IndexHandler{}
	user := handlers.UserHandler{}
	e.GET("/", index.IndexHandler)
	e.GET("/user", user.User)
	e.Logger.Fatal(e.Start(":3000"))
}
