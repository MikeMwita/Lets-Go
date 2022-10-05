package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	var words []string = strings.Fields(s)

	for _, words := range words {
		result[words] += 1
	}
	return result
}
func main() {
	wc.Test(WordCount)
}
