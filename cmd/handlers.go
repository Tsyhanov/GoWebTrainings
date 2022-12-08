package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	fmt.Println("login handle")
	return c.Render(http.StatusOK, "index.tmpl.html", "index")

}
