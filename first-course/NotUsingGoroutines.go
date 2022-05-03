package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	printSomething()
	printSomethingElse()

	fmt.Println("I'm done")
	elapsed := time.Since(start)
	fmt.Printf("The Processes took %s", elapsed)
}

func printSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something")
}

func printSomethingElse() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something else")
}
