package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go numbers(chan1)
	go quart(chan1, chan2)
	for {
		num, ok := <-chan2
		if !ok {
			break
		}
		fmt.Printf("%d ", num)
	}
}

func numbers(outChan chan<- int) {
	for i := 0; i < 10; i++ {
		outChan <- i + 1
	}
	close(outChan)
}

func quart(inputChan <-chan int, outChan chan<- int) {
	for {
		num, ok := <-inputChan
		if !ok {
			break
		}
		outChan <- num * num
	}
	close(outChan)
}
