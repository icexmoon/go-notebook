package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%#v", r.URL)
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/hello/", helloHandleFunc)
	server.ListenAndServeTLS("../ssl/cert.pem", "../ssl/key.pem")
}
