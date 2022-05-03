package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "Channel 1, sending every one second"
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			ch2 <- "Channel 2, sending every 5 second"
		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			ch3 <- "Channel 3, we're done"
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg + " ,something cool happened")
		case msg := <-ch3:
			fmt.Println(msg)
			os.Exit(0)
		}
	}
}
