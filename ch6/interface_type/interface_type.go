package main

import (
	"fmt"

	myinterface "github.com/icexmoon/go-notebook/ch6/my_interface"
)

type myByteArr []byte

func (myByteArr) Read([]byte) (length int, err error) {
	return
}

func main() {
	var mba myByteArr = nil
	if mba == nil {
		fmt.Println("mba == nil")
		// mba == nil
	}
	var reader myinterface.Reader = mba
	if reader == nil {
		fmt.Println("reader == nil")
	}
}
