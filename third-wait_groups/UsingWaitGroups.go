package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(2)
	go checkSomething()
	go checkSomethingElse()

	wg.Wait()
	fmt.Println("I'm done")
	elapsed := time.Since(start)
	fmt.Printf("The Processes took %s", elapsed)
}

func checkSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something")
	wg.Done()
}

func checkSomethingElse() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something else")
	wg.Done()
}
