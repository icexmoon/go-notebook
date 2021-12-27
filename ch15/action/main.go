package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	p := Person{
		Name: "icexmoon",
		Age:  19,
	}
	t.Execute(rw, p)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
