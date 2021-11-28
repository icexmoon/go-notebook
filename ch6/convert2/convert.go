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
	sc := read.(*myReader)
	// panic: interface conversion: myinterface.Reader is *stringcontainer.StringContainer, not *main.myReader
	fmt.Println(sc)
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
