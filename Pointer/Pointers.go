package Pointer

import "fmt"

//a pointer holds a memory address  of a variable
//the & operator generates a pointer to its operand
//* denotes the pointers underlying variable
func main() {
	i, j := 42, 100
	p := *i         //points to i
	fmt.Println(*p) //reads i thro teh pointer
	*p = 21
	fmt.Println(i)

	p = &j //p pointing to j
	*p = *p / 7
	fmt.Println(j)
}
