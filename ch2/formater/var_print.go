package formater

import "fmt"

func PrintVariable(variable interface{}) {
	fmt.Printf("%#v %T %v\n", variable, variable, variable)
}
