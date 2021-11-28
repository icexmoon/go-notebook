package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	checkStringNumber("12")
	checkStringNumber("2.5")
	checkStringNumber("lalala")
}

func checkStringNumber(strA string) {
	if floatA, err := strconv.ParseFloat(strA, 64); err != nil {
		log.Fatal(err)
	} else if floatA > 10 {
		fmt.Println("a>10")
	} else {
		fmt.Println("a<=10")
	}
}

// a>10
// a<=10
// 2021/11/07 15:38:50 strconv.ParseFloat: parsing "lalala": invalid syntax
