package main

import "fmt"

type Celsius float64 //摄氏温度
func (c Celsius) String() string {
	return fmt.Sprintf("%.1fC", c)
}

type Fahrenheit float64 //华氏温度
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.1fF", f)
}

func main() {
	zero := Celsius(0)
	fmt.Println(zero)
	fmt.Println(changeC2F(zero))
	boil := Celsius(100)
	fmt.Println(boil)
	fmt.Println(changeC2F(boil))
}

func changeC2F(c Celsius) Fahrenheit {
	return Fahrenheit(9*c/5 + 32)
}

func changeF2C(f Fahrenheit) Celsius {
	c := Celsius(5 * (f - 32) / 9)
	return c
}
