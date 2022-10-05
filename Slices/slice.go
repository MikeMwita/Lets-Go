package main

import "fmt"

//a slice has a dynamic size
//syntax:[]T
//a[low;high]-->contains 2 indices (low and high bound)
func main() {

	primes := [5]int{2, 3, 5, 7, 11}
	//we  write the values in a message string
	fmt.Sprintln(primes[0], primes[0:1], primes[:1])

	var s []int = primes[1:4]
	fmt.Println(s)
	var s1 []int = primes[2:]
	fmt.Println(s1)

	var s2 []int = primes[:2]
	fmt.Println(s2)
}
