package main

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	result := Add(a, b)
	if result != 3 {
		t.Fatal("1 + 2 is not 3")
	}
	a = -1
	b = -3
	result = Add(a, b)
	if result != -4 {
		t.Fatal("-1 + -3 is not -4")
	}
}
