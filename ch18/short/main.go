package main

import (
	"time"
)

func LongTimeAdd(a int, b int) int {
	time.Sleep(3 * time.Second)
	return a + b
}

func main() {
}
