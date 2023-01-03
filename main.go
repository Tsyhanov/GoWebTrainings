package main

import (
	"test-registration-form/config"
	"test-registration-form/pkg/concurrent"
	"test-registration-form/pkg/db"
	"test-registration-form/pkg/handlers"
	"test-registration-form/pkg/render"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.SetConfig()
}

func main() {
	//db
	db.Init()
	//TODO:
	//get posts and comments from external endpoint...
	concurrent.GetPostsAndComments()

	//web server
	e := echo.New()
	e.Static("/static", "web/static")
	//render
	templates := make(map[string]*template.Template)
	templates["login.tmpl.html"] = template.Must(template.ParseFiles("web/templates/login.tmpl.html", "web/templates/base.tmpl.html"))
	templates["signup.tmpl.html"] = template.Must(template.ParseFiles("web/templates/signup.tmpl.html", "web/templates/base.tmpl.html"))
	templates["home.tmpl.html"] = template.Must(template.ParseFiles("web/templates/home.tmpl.html", "web/templates/base.tmpl.html"))
	e.Renderer = &render.Template{
		Templates: templates,
	}
	//routes
	e.GET("/", handlers.Login)
	e.GET("/signup", handlers.Signup)
	e.POST("/login", handlers.PostLogin)
	e.POST("/signup", handlers.PostSignup)

	//routes restricted
	r := e.Group("/restricted")
	r.Static("/static", "web/static")
	r.Use(middleware.JWTWithConfig(config.AuthConfig))

	r.GET("/home", handlers.Home)
	r.GET("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}
