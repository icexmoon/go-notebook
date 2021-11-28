package main

import "fmt"

func main() {
	var a = 1
	var b int
	a, b := 2, 3
	fmt.Println(a, b)
}

// Build Error: go build -o C:\Users\70748\AppData\Local\Temp\__debug_bin3053900173.exe -gcflags all=-N -l .\variable9.go
// # command-line-arguments
// .\variable9.go:8:7: no new variables on left side of := (exit status 2)
