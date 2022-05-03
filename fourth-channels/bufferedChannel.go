package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "Hello "
	c <- "World "
	c <- "!"
	// c <- "I'm excited" //if you don't comment, it will be error caused by deadlock, channel only have size 3 but you used 5
	// c <- "How about you?"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
