package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	fmt.Println("home handle")
	return c.Render(http.StatusOK, "home.tmpl.html", "home")
}

func Login(c echo.Context) error {
	fmt.Println("home handle")
	return c.Render(http.StatusOK, "login.tmpl.html", "login")
}
