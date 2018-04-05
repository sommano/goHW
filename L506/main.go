package main

import "fmt"

func main() {
	s := "c"

	switch s {
	case "a":
		fmt.Println("the letter a!")
	case "b":
		fmt.Println("the letter b!")
	case "c":
		fmt.Println("the letter c!")
	default:
		fmt.Println("I don't recognize that letter")
	}
}
