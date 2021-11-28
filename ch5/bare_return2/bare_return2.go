package main

import (
	"fmt"
	"strconv"
)

func myFunc6(strNum string) (int, error) {
	var intNum int
	var err error
	intNum, err = strconv.Atoi(strNum)
	return intNum, err
}

func main() {
	intNum, _ := myFunc6("123")
	fmt.Println(intNum)
	// 123
}
