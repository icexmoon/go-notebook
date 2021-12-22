package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

type bye struct{}

func (bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "bye~~")
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.Handle("/hello", hello{})
	serverMux.Handle("/bye", bye{})
	server := http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
