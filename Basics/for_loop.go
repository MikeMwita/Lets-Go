package main

import "fmt"

//the type of loop that go uses is the for loop structure
//also for is like the while loop in other languages
func main() {
	sum := 0
	sum1 := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum1 < 100 {
		sum1 += sum1
	}
	fmt.Println(sum1)
}
