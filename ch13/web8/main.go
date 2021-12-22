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

func welcome(hf func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "welcome everyone!\n")
		hf(rw, r)
	}
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", welcome(helloHandleFunc))
	serverMux.HandleFunc("/bye", welcome(byeHandleFunc))
	server := http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
