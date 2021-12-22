package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
)

type book struct {
	Name string
	Id   int
	Desc string
}

var books = map[int]book{
	1: {Name: "哈利波特", Id: 1, Desc: "小说"},
	2: {Name: "时间简史", Id: 2, Desc: "科普读物"},
	3: {Name: "Go程序设计语言", Id: 3, Desc: "程序设计"},
}

func bookHandleFunc(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Proto)
	id := p.ByName("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	book, ok := books[bookId]
	if !ok {
		fmt.Fprintf(rw, "not find the book.")
	} else {
		fmt.Fprintf(rw, "Book details:\n")
		fmt.Fprintf(rw, "name: %s\n", book.Name)
		fmt.Fprintf(rw, "id: %d\n", book.Id)
		fmt.Fprintf(rw, "description:%s\n", book.Desc)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/book/:id", bookHandleFunc)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("../ssl/cert.pem", "../ssl/key.pem")
}
