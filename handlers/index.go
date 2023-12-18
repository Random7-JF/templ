package handlers

import (
	"log"

	view "github.com/Random7-JF/templ/views/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func (h IndexHandler) IndexHandler(c echo.Context) error {
	log.Printf("Index Handler")
	return render(c, view.Index("Jon"))
}
