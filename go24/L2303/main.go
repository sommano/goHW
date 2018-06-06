package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("You have 2 seconds to calculate 19 *5")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Time's up! The answer is 74. Did u get it?")
			return
		}
	}
}
