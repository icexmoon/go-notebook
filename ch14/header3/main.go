package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		if len(value) > 1 {
			fmt.Fprintf(rw, "%s\n", key)
		}
	}
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
