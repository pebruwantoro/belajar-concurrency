package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {
	start := time.Now()
	go checkSomething()
	go checkSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("I'm done")
	elapsed := time.Since(start)
	fmt.Printf("The Processes took %s", elapsed)
}

func checkSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something")
	ch <- "checkSomething finished"
}

func checkSomethingElse() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something else")
	ch <- "checkSomethingElse finished"
}
