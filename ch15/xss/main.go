package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.New("submit.html")
	t.ParseFiles("submit.html")
	t.Execute(rw, nil)
}

func submit(rw http.ResponseWriter, r *http.Request) {
	comment := r.PostFormValue("comment")
	t := template.Must(template.ParseFiles("show.html"))
	t.Execute(rw, comment)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
