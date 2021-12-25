package main

import (
	"fmt"
	"net/http"
)

func writeCookie(rw http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "first_cookie",
		Value: "cookie1",
	}
	c2 := http.Cookie{
		Name:  "second_cookie",
		Value: "cookie2",
	}
	rw.Header().Set("Set-Cookie", c1.String())
	rw.Header().Add("Set-Cookie", c2.String())
}

func getCookie(rw http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for _, c := range cookies {
		fmt.Fprintf(rw, "%s:%s\n", c.Name, c.Value)
	}
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/write_cookie", writeCookie)
	serverMux.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
