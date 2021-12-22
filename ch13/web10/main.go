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

func welcome(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "welcome everyone!\n")
		hf(rw, r)
	}
}

func userValidate(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//登录验证模块
		hf(rw, r)
	}
}

func log(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//日志记录模块
		hf(rw, r)
	}
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", log(userValidate(helloHandleFunc)))
	serverMux.HandleFunc("/bye", log(userValidate(byeHandleFunc)))
	server := http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
