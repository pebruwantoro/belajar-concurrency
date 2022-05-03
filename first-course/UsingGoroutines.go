package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go doSomething()
	go doSomethingElse()

	fmt.Println("I'm done")
	elapsed := time.Since(start)
	fmt.Printf("The Processes took %s", elapsed)
}

func doSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something")
}

func doSomethingElse() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something else")
}

/*
Output
I'm done
The Processes took 536.7µ
*/
