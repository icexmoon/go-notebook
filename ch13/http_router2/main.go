package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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
	http.ListenAndServe(":8080", router)
}
