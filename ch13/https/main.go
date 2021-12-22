package main

import "net/http"

func main() {
	http.ListenAndServeTLS("127.0.0.1:8080", "../ssl/cert.pem", "../ssl/key.pem", nil)
}
