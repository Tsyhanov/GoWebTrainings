package main

import (
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login handler")
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index handler")
	err := tmpl.ExecuteTemplate(w, "index.tmpl.html", nil)

	if err != nil {
		log.Println("index template execution error")
	}

}
