package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
