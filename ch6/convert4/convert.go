package main

import (
	"fmt"
	"log"

	myinterface "github.com/icexmoon/go-notebook/ch6/my_interface"
	stringcontainer "github.com/icexmoon/go-notebook/ch6/string_container"
)

type myReader struct {
}

func (mr *myReader) Read(container []byte) (length int, err error) {
	return
}

func dealSC(read myinterface.Reader) {
	if read == nil {
		fmt.Println("nil")
	} else if sc, ok := read.(*stringcontainer.StringContainer); ok {
		sc.SetStr("test is changed")
	} else if _, ok := read.(*myReader); ok {
		fmt.Println("myReader")
	}
}

func main() {
	var sc stringcontainer.StringContainer
	sc.SetStr("test")
	dealSC(&sc)
	line := make([]byte, 0, 20)
	length, err := sc.Read(line)
	if err != nil {
		log.Fatalln(err)
	}
	line = line[:length]
	fmt.Println(string(line))
	// test is changed
}
