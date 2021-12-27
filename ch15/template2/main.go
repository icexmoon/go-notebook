package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	content := `
	<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<h1>{{ . }}</h1>
			<h1>index page.</h1>
		</body>
	</html>
	`
	t.Parse(content)
	msg := "hello world!"
	t.Execute(rw, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
