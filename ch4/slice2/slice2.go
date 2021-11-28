package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	changeSlice(slice)
	fmt.Println(slice)
	// [99 2 3]
}

func changeSlice(slice []int) {
	slice[0] = 99
}
