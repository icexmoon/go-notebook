package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html", "header.html"))
	msg := "hello world!"
	t.ExecuteTemplate(rw, "html", msg)
}

func index2(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	msg := "hello world!"
	t.ExecuteTemplate(rw, "html", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/index2", index2)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
