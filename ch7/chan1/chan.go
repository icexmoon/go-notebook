package main

import "fmt"

func main() {
	var chan1 chan int
	if chan1 == nil {
		fmt.Println("chan1 is nil")
		// chan1 is nil
	}
	chan1 = make(chan int)
	chan2 := make(chan int, 3)
	fmt.Println(cap(chan1))
	// 0
	fmt.Println(cap(chan2))
	// 3
	chan1 <- 1
	var num int = <-chan1
	close(chan1)
}
