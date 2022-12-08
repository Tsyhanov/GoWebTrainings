package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {

	fs := http.FileServer(http.Dir("./web/static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	//web server
	e := echo.New()
	//render
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("web/templates/*")),
	}
	e.Renderer = tmpl
	//routes
	e.GET("/", Index)
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	fmt.Println("web_test1: Echo Server started")
	e.Logger.Fatal(e.Start(":8080"))
}
