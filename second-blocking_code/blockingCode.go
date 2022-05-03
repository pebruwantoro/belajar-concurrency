package main

import (
	"fmt"
	"time"
)

func main() {
	doTheFirstThing()                          // do the first thing
	doTheSecondThing()                         // do the second thing
	fmt.Println("No more blocking.  I'm done") // last step
}
func doTheFirstThing() {
	fmt.Println("FirstThing 'blocking' for 2 seconds")
	time.Sleep(time.Second * 2) // blocking
}

func doTheSecondThing() {
	fmt.Println("SecondThing 'blocking' for 3 seconds")
	time.Sleep(time.Second * 3) //blocking
}
