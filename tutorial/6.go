package main

import (  
    "fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for i < n {
	sum += numbers[i]
	i++
	}
	fmt.Println("Sum is :: ", sum)
}
