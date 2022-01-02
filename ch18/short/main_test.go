package main

import "testing"

func TestLongTimeAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping LongTime func test in short mode")
	}
	if LongTimeAdd(1, 5) != 6 {
		t.Error("result of LongTimeAdd(1, 5) is not 6")
	}
}
