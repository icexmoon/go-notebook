package main

import (
	"fmt"
	"log"

	myinterface "github.com/icexmoon/go-notebook/ch6/my_interface"
	stringcontainer "github.com/icexmoon/go-notebook/ch6/string_container"
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
