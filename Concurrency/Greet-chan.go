package main

import "log"

func greeting(ch chan string) {
	ch <- "hello mike"

}
func main() {
	ch := make(chan string)
	go greeting(ch)
	log.Print(<-ch)

}
