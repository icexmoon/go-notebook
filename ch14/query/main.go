package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(rw, "%#v", r.Form)
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
