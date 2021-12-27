package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t.ParseFiles("index.html", "header.html", "footer.html")
	t.ExecuteTemplate(rw, "index.html", "<Hello world!>")
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
