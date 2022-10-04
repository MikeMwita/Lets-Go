package main

import "fmt"

func change(value *int) {
	*value = 10
}

func modify2(arr *[3]int) {
	arr[0] = 50
}

func main() {
	//create a pointer
	b := 100
	var a *int = &b
	fmt.Printf("The type of a is :%T\n", a)
	fmt.Println("The address of variable b is :", a)

	//derefereencing
	m := 255
	n := &m

	fmt.Println("The address of m is :", n)
	fmt.Println("The value of n is :", *n)

}
