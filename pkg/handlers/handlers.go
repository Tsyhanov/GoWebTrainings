package handlers

import (
	"fmt"
	"net/http"
	"test-registration-form/config"
	"test-registration-form/pkg/auth"

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
	//TODO: check if user is available in db
	//....
	//Generate JWT
	storedUser := config.LoadTestUser() //for test only
	err := auth.GenerateTokensAndSetCookies(storedUser, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusSeeOther, "/restricted/home")

}

//show home page after login
func Home(c echo.Context) error {
	fmt.Println("home handle")
	return c.Render(http.StatusOK, "home.tmpl.html", "home")
}
