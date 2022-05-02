package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("This the number %v", i)
		go printString(str)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func printString(str string) {
	fmt.Println(str)
	wg.Done()
}
