package handlers

import (
	"fmt"
	"net/http"
	"test-registration-form/config"
	"test-registration-form/models"
	"test-registration-form/pkg/auth"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

//get handler to show startup login page
func Login(c echo.Context) error {
	fmt.Println("login handle")
	return c.Render(http.StatusOK, "login.tmpl.html", "login")
}

//get handler to show signup page
func Signup(c echo.Context) error {
	fmt.Println("signup handle")
	return c.Render(http.StatusOK, "signup.tmpl.html", "signup")
}

//post handler. get email and password and create jwt
func PostLogin(c echo.Context) error {
	fmt.Println("PostLogin")
	//TODO: check if user is available in db
	//....
	//Generate JWT
	storedUser := config.LoadTestUser() //for test only
	//create new user based on User struct
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	fmt.Println(storedUser.Password)
	fmt.Println(u.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
	}
	//password is ok - generate token
	err := auth.GenerateTokensAndSetCookies(storedUser, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusSeeOther, "/restricted/home")
}

//post handler for signup form
func PostSignup(c echo.Context) error {
	fmt.Println("postsignup handle")
	name := c.FormValue("name")
	fmt.Println("name is ", name)
	return c.Redirect(http.StatusSeeOther, "/restricted/home")
}

//show home page after login
func Home(c echo.Context) error {
	fmt.Println("home handle")
	return c.Render(http.StatusOK, "home.tmpl.html", "home")
}

//Logout (to root page)
func Logout(c echo.Context) error {
	fmt.Println("logout handle")

	cookie := new(http.Cookie)
	cookie.Name = config.Config.TokenCookieName
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, "/")
}
