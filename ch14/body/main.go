package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, string(body))
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
