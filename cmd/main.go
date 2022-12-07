package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("web/templates/*"))
}

func main() {

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//	e := echo.New()
	//	e.GET("/", Index)

	http.HandleFunc("/", Index)

	fmt.Println("web_test1: Echo Server started")
	http.ListenAndServe(":8080", nil)
	//	e.Logger.Fatal(e.Start(":8080"))
}
