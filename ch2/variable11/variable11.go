package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := "11"
	b := "1.5"
	intA, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatal(err)
	}
	if intA < 100 {
		fmt.Println("a<100")
	}
	if floatB < 5 {
		fmt.Println("b<5")
	}
}
