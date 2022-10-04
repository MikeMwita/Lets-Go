package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	Emp1 := &Employee{"Mike", "Mwita", 20, 1}

	fmt.Println("The first name is :", (*Emp1).firstName)
	fmt.Println("age is :", (*Emp1).age)

}
