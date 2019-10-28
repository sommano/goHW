package main

import (  
    "fmt"
)

func isEven(value int) {
	switch {
	case value%2 == 0:
	fmt.Println("I is even")
	default:
	fmt.Println("I is odd")
	}
}
