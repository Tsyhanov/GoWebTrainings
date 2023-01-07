package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test-registration-form/config"
	"test-registration-form/models"
	"test-registration-form/pkg/auth"
	"test-registration-form/pkg/db"
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
	//create new user based on User struct
	u := new(models.User)
	storedUser := new(models.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//check email and get user from db
	//storedUser := config.LoadTestUser() //for debug only
	if !db.CheckIfEmailExist(u.Email) {
		return echo.NewHTTPError(http.StatusUnauthorized, "email does not exist")
	}
	err := db.GetUser(storedUser, u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Can not get user from db")
	}
	// Compare hashed password from db with the password from form
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
	}
	//password is ok - generate token
	err = auth.GenerateTokensAndSetCookies(storedUser, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusSeeOther, "/restricted/home")
}

//post handler for signup form
func PostSignup(c echo.Context) error {
	fmt.Println("postsignup handle")
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm_password")
	if password != confirmPassword {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password mismatch")
	}

	if db.CheckIfEmailExist(email) {
		return echo.NewHTTPError(http.StatusUnauthorized, "email already exist")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	user := models.User{Name: name, Email: email, Password: string(hashedPassword)}
	//add new user in db
	if !db.CreateUser(&user) {
		return echo.NewHTTPError(http.StatusUnauthorized, "could not create new user")
	}
	//generate token for new user
	err := auth.GenerateTokensAndSetCookies(&user, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusSeeOther, "/restricted/home")
}

//home page handle (after login or signup)
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

// GetComments godoc
// @Summary Get all comments
// @Produce json
// @Produce xml
// @Success 200 {array} models.Comment
// @Router /restricted/comments [get]
func GetComments(c echo.Context) error {
	var cmt []models.Comment
	cmt, err := db.GetComments()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, cmt)
}

// GetPosts godoc
// @Summary Get all posts
// @Produce json
// @Produce xml
// @Success 200 {array} models.Post
// @Router /restricted/posts [get]
func GetPosts(c echo.Context) error {
	var p []models.Post
	p, err := db.GetPosts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, p)
}

// GetPostById godoc
// @Summary Get post by id
// @Produce json
// @Produce xml
// @Param id path int true "Post Id"
// @Success 200 {object} models.Post
// @Router /restricted/posts/{id} [get]
func GetPostById(c echo.Context) error {
	var p models.Post
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := db.GetPostById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, p)
}

// AddPost godoc
// @Summary Add post
// @Produce json
// @Produce xml
// @Param userid formData int true "User Id"
// @Param title formData string true "Post Title"
// @Param body formData string true "Post Body"
// @Success 200 {object} models.Post
// @Router /restricted/posts/add [post]
func AddPost(c echo.Context) error {
	fmt.Println("addPost")
	userid := c.FormValue("userid")
	title := c.FormValue("title")
	body := c.FormValue("body")

	i, _ := strconv.Atoi(userid)
	p := models.Post{UserId: i, Title: title, Body: body}
	err := db.AddPost(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, p)
}

// AddComment godoc
// @Summary Add comment for given post Id
// @Produce json
// @Produce xml
// @Param postid formData int true "post Id"
// @Param name formData string true "Name"
// @Param email formData string true "E-mail"
// @Param body formData string true "Body"
// @Success 200 {object} models.Comment
// @Router /restricted/posts/{id}/comments/add [post]
func AddComment(c echo.Context) error {
	fmt.Println("addComment")
	postid, _ := strconv.Atoi(c.FormValue("postid"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	body := c.FormValue("body")

	cmt := models.Comment{PostId: postid, Name: name, Email: email, Body: body}
	err := db.AddComment(cmt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, cmt)
}
