package main

import (
	"github.com/Random7-JF/templ/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	index := handlers.IndexHandler{}
	e.GET("/", index.IndexHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
