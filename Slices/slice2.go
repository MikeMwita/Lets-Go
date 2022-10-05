package main

import "fmt"

func genSlice() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 30) //contains both length and capacity
	return s1, s2, s3
}
func main() {
	s1, s2, s3 := genSlice()
	fmt.Printf("s1:len=%v  cap =%v\n", len(s1), cap(s1))
	fmt.Printf("s2:len=%v  cap =%v\n", len(s2), cap(s2))
	fmt.Printf("S3: len=%v  cap =%v\n", len(s3), cap(s3))
}
