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
	switch read := read.(type) {
	case nil:
		fmt.Println("nil")
	case *stringcontainer.StringContainer:
		read.SetStr("test is changed")
	case *myReader:
		fmt.Println("myReader")
	default:
		fmt.Println("other")
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
