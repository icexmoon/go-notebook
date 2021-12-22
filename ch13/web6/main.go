package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

func byeHandleFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "bye~~")
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", helloHandleFunc)
	serverMux.HandleFunc("/bye", byeHandleFunc)
	server := http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
