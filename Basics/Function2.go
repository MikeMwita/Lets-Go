package main

import "fmt"

//a function to swap strings
//we will use pointers to swap the string
func swapString(str1, str2 *string) {
	*str1, *str2 = *str1, *str2
}
func main() {
	a := "mike"
	b := "mwita"
	//we use the swap func to swap
	swapString(&a, &b)
	fmt.Printf("string a is %s\n  string b is %s\n", a, b)
}
