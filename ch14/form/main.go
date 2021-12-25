package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func userSubmit(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(rw, "%s\n", r.Form.Get("name"))
	fmt.Fprintf(rw, "%s\n", r.Form.Get("age"))
}

func userInfo(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("user.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/user_submit", userSubmit)
	serverMux.HandleFunc("/user_info", userInfo)
	server.ListenAndServe()
}
