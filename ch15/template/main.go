package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t.ParseFiles("index.html")
	msg := "hello world!"
	t.Execute(rw, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
