package main

import (
	"fmt"
	myinterface "go-notebook/ch6/my_interface"
	sc "go-notebook/ch6/string_container"
	"log"
)

func readAndPrint(reader myinterface.Reader) {
	for {
		line := make([]byte, 0, 20)
		length, err := reader.Read(line)
		line = line[0:length]
		if err != nil {
			log.Fatalln(err)
		}
		if length == 0 {
			return
		}
		fmt.Print(string(line))
	}
}

func main() {
	var scontainer sc.StringContainer
	scontainer.SetStr("Hello!\nHow are you!\nI'm fine.")
	readAndPrint(&scontainer)
	// Hello!
	// How are you!
	// I'm fine.
}
