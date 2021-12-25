package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

type result struct {
	Result string `json:"result"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}

type person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func userSubmit(rw http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		log.Fatal(err)
	}
	p := new(person)
	if err := json.Unmarshal(body, p); err != nil {
		log.Fatal(err)
	}
	name := p.Name
	age := p.Age
	resp, err := json.Marshal(result{Result: "ok", Name: name, Age: age})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(rw, "%s", string(resp))
}

func userInfo(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("user.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/user_submit", userSubmit)
	serverMux.HandleFunc("/user_info", userInfo)
	server.ListenAndServe()
}
