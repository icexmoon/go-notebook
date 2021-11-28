package main

import "fmt"

type Pointer struct {
	x int
	y int
}

type Pointers struct {
	pointers []Pointer
}

func newPointer(x, y int) Pointer {
	return Pointer{x: x, y: y}
}

func main() {
	var p1 Pointer
	fmt.Println(p1)
	// {0 0}
	p1 = Pointer{x: 1, y: 1}
	fmt.Println(p1)
	// {1 1}
	var pts Pointers = Pointers{pointers: []Pointer{p1, {x: 1, y: 2}}}
	fmt.Println(pts)
	// {[{1 1} {1 2}]}

}
