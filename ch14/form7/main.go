package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func userSubmit(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		log.Fatal(err)
	}
	fh := r.MultipartForm.File["upload"][0]
	file, err := fh.Open()
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, string(contents))
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
