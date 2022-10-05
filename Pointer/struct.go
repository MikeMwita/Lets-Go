package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	Emp1 := &Employee{"Mike", "Mwita", 20, 1}
	//the prefix & returns a pointer to the struct value
	fmt.Println("The first name is :", (*Emp1).firstName)
	fmt.Println("age is :", (*Emp1).age)
	fmt.Println(Emp1)

}

//NB-->struct fiels can be accessed using a struct pointer

//to access teh x of a struct -->(*p).x without the explicit dereference.
