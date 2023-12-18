package handlers

import (
	"log"

	userview "github.com/Random7-JF/templ/views/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h *UserHandler) User(c echo.Context) error {
	log.Printf("Index Handler")
	return render(c, userview.UserDisplay(4, "Jon", "jon@jon.com"))
}
