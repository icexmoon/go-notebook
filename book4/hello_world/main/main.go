// main.go
package main

import (
	"fmt"
	"hello_world/modules"
)

func main() {
	jack := modules.NewStudent("Jack Chen", 20)
	brus := modules.NewStudent("Brus Lee", 12)
	jack.Print()
	brus.Print()
	fmt.Print(jack.GetName())
}
