package main

import (
	"log"
	"net/http"
	"text/template"
)

func userSubmit(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(501)
}

func userInfo(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("user.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}

func redirect(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	header["location"] = append(header["location"], "/user_info")
	rw.WriteHeader(301)
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/user_submit", userSubmit)
	serverMux.HandleFunc("/user_info", userInfo)
	serverMux.HandleFunc("/redirect", redirect)
	server.ListenAndServe()
}
