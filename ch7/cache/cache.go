package main

import "fmt"

func main() {
	cachedChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		cachedChan <- i
	}
	for num := range cachedChan {
		fmt.Printf("%d ", num)
		if len(cachedChan) <= 0 {
			break
		}
	}
	// 0 1 2 3 4 5 6 7 8 9
}
