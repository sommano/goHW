package main

import (
	"fmt"
)

var s = "hello world"

func main() {
	fmt.Printf("Print 's' variable from outer block %v\n", s)
	b := true
	if b {
		fmt.Printf("Printing 'b' variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Printing 'i' variable from outer block %v\n", i)
		}
	}
}
