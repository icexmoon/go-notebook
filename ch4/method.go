package main

import "fmt"

type Celsius float64 //摄氏温度
func (c Celsius) String() string {
	return fmt.Sprintf("%.1fC", c)
}

//重新将温度设置为0摄氏度
func (c *Celsius) Reset() {
	*c = 0
}

type Fahrenheit float64 //华氏温度
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.1fF", f)
}

func main() {
	zero := Celsius(0)
	zeroString := Celsius.String(zero)
	fmt.Println(zeroString)
	// 0.0C
	t1 := Celsius(100)
	fmt.Println(t1)
	// 100.0C
	t1.Reset()
	fmt.Println(t1)
	// 0.0C

}

func changeC2F(c Celsius) Fahrenheit {
	return Fahrenheit(9*c/5 + 32)
}

func changeF2C(f Fahrenheit) Celsius {
	c := Celsius(5 * (f - 32) / 9)
	return c
}
