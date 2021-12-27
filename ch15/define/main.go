package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.ExecuteTemplate(rw, "html", nil)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
