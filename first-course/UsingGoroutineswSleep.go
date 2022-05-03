package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go checkSomething()
	go checkSomethingElse()

	time.Sleep(time.Second * 5)
	fmt.Println("I'm done")
	elapsed := time.Since(start)
	fmt.Printf("The Processes took %s", elapsed)
}

func checkSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something")
}

func checkSomethingElse() {
	time.Sleep(1 * time.Second)
	fmt.Println("I've already print something else")
}

/*
Output

---wait---
I've already print something
I've already print something else
--wait---
I'm done
The Processes took 2.0175203s
*/
