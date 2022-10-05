package main

import "fmt"

//maps are created using built in make func that only accepts one argument-which is built in map type

func main() {

	myMap := make(map[string]int)
	//we assign values to the map

	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Printf("mymap is :%v/n", myMap)
	fmt.Printf("map one :%v   map two :%v", myMap["one"], myMap["two"])
	//deleting an element from a map
	delete(myMap, "one")
	fmt.Println("the value", myMap)

	//in order to test whether a key is present --we use a two value assignment

	v, ok := myMap["one"]
	fmt.Println("The value :", v, "Present?", ok)

}
