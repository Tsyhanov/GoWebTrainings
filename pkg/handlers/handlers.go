package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//get handler to show startup login page
func Login(c echo.Context) error {
	fmt.Println("login handle")
	return c.Render(http.StatusOK, "login.tmpl.html", "login")
}

//post handler. get email and password and create jwt
func PostLogin(c echo.Context) error {
	fmt.Println("postlogin handle")
	return nil
}

//show home page after login
func Home(c echo.Context) error {
	fmt.Println("home handle")
	return c.Render(http.StatusOK, "home.tmpl.html", "home")
}
