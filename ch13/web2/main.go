package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

func main() {
	my_server := http.Server{
		Addr:    ":8080",
		Handler: hello{},
	}
	my_server.ListenAndServe()
}
