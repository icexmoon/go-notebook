package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	changeSlice(slice)
	fmt.Println(slice)
	// [1 2 3]
}

func changeSlice(slice []int) {
	slice = append(slice, 4)
	slice = append(slice, 5)
}
