package main

import "fmt"

func main() {
	login("lalala")
	login("root")
	login("apple")
}

func login(name string) {
	switch name {
	case "apple":
		fmt.Println("welcome")
	case "root":
		fmt.Println("don't permit use root account login")
	default:
		fmt.Println("it's a unregistry account")
	}
}

// it's a unregistry account
// don't permit use root account login
// welcome
