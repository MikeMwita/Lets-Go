package main

import (
	"fmt"
	"math/rand"
	"time"
)

//a go program that generates random numbers
//we use a built in package to generate random no's ==.which contains a rand.Intn() function

func main() {
	rand.Seed(time.Now().UnixNano()) //this sets the seed
	r := rand.Intn(10) + 1           //generates a random no btwn 0 thereafter adds 1 to the no to get the rand no btwn 0 and 10
	fmt.Println(r)
}
