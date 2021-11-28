package main

import (
	"fmt"
	"strconv"
)

func myFunc5(strNum string) (intNum int, err error) {
	intNum, err = strconv.Atoi(strNum)
	return
}

func main() {
	intNum, _ := myFunc5("123")
	fmt.Println(intNum)
	// 123
}
