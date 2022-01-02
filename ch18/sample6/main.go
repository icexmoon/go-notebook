package main

import (
	"time"
)

func LangTimeAdd1(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}

func LangTimeAdd2(a int, b int) int {
	time.Sleep(2 * time.Second)
	return a + b
}

func LangTimeAdd3(a int, b int) int {
	time.Sleep(3 * time.Second)
	return a + b
}

func main() {
}
