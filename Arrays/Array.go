package main

import "fmt"

func main() {
	//arrays cannot be resized
	var a [2]string
	a[0] = "am"
	a[1] = "mike"
	fmt.Println(a[0], a[1])
	fmt.Print(a)

	names := [3]string{"am ", "mike", "mwita"}
	fmt.Println(names)

}
