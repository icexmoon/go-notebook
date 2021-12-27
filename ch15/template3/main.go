package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := "hello world!"
	t.Execute(rw, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
