package main

import "fmt"

//go is a typecasted language

func main() {
	//declaring variables
	//NB-->a Var can be declared at package or function level
	var car string = "bmw"
	fmt.Println(car)

	//using the short variable declaration
	name := "mike"
	count := 2
	fmt.Println(name)
	fmt.Println(count)

	x := 10
	y := uint16(10)
	z := uint8(10)
	fmt.Println(x, y, z)

	//variable formatting
	//%v-->to print any value,%T-->to print the variables type ,--->%d->the decimal base,-->%s( string)
	fmt.Println("The vaue of name is :%d", name)
	//fmt.Println("The value of count is : %d\n", count)

}
