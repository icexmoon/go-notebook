package main

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	result := Add(a, b)
	if result != (a + b) {
		t.Error("result is not a + b")
	}
}
