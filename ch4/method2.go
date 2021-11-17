package main

import "fmt"

type Pointer struct {
	x int
	y int
}

func (p Pointer) add(other Pointer) Pointer {
	return Pointer{x: p.x + other.x, y: p.y + other.y}
}

func (p *Pointer) selfAdd(other Pointer) {
	p.x += other.x
	p.y += other.y
}

func main() {
	p1 := Pointer{x: 1, y: 1}
	p2 := Pointer{x: 5, y: 5}
	fmt.Println(p1.add(p2))
	// {6 6}
	fmt.Println(p1)
	// {1 1}
	p1.selfAdd(p2)
	fmt.Println(p1)
	// {6 6}
}
