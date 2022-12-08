package main

import (
	"fmt"
	"net/http"
	"test-registration-form/pkg/http/handlers"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	//file server
	fs := http.FileServer(http.Dir("./web/static"))
	//web server
	e := echo.New()

	//render
	templates := make(map[string]*template.Template)
	templates["login.tmpl.html"] = template.Must(template.ParseFiles("web/templates/login.tmpl.html", "web/templates/base.tmpl.html"))
	templates["home.tmpl.html"] = template.Must(template.ParseFiles("web/templates/home.tmpl.html", "web/templates/base.tmpl.html"))
	e.Renderer = &Template{
		templates: templates,
	}
	//routes
	e.GET("/", handlers.Login)

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	fmt.Println("web_test1: Echo Server started")
	e.Logger.Fatal(e.Start(":8080"))
}
