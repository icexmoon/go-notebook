package main

import (
	"fmt"
	myinterface "go-notebook/ch6/my_interface"
	stringcontainer "go-notebook/ch6/string_container"
	"log"
)

func dealSC(read myinterface.Reader) {
	var sc *stringcontainer.StringContainer = read.(*stringcontainer.StringContainer)
	sc.SetStr("test is changed")
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
