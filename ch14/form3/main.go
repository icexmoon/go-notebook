package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func userSubmit(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(rw, "%#v\n", r.MultipartForm)
	fmt.Fprintf(rw, "%s\n", r.MultipartForm.Value["name"][0])
	fmt.Fprintf(rw, "%s\n", r.MultipartForm.Value["age"][0])
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
