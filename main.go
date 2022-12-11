package main

import (
	"test-registration-form/config"
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
	//web server
	e := echo.New()
	e.Static("/static", "web/static")
	//render
	templates := make(map[string]*template.Template)
	templates["login.tmpl.html"] = template.Must(template.ParseFiles("web/templates/login.tmpl.html", "web/templates/base.tmpl.html"))
	templates["home.tmpl.html"] = template.Must(template.ParseFiles("web/templates/home.tmpl.html", "web/templates/base.tmpl.html"))
	e.Renderer = &render.Template{
		Templates: templates,
	}
	//routes
	e.GET("/", handlers.Login)
	e.POST("/login", handlers.PostLogin)
	//routes restricted
	r := e.Group("/restricted")
	r.Static("/static", "web/static")
	r.Use(middleware.JWTWithConfig(config.AuthConfig))

	r.GET("/home", handlers.Home)

	e.Logger.Fatal(e.Start(":8080"))
}
