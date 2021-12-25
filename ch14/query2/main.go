package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(rw, "%s\n", r.Form["name"][0])
	fmt.Fprintf(rw, "%s\n", r.Form["id"][0])
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/hello/", helloHandleFunc)
	server.ListenAndServe()
}
