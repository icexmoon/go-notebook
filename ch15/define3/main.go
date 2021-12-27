package main

import (
	"html/template"
	"net/http"
	"time"
)

func index(rw http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")
	t := template.Must(template.ParseFiles("index.html", "footer.html", "header.html"))
	t.ExecuteTemplate(rw, "html", today)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
