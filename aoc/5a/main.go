package main

import (
	"fmt"
	"io/ioutil"
)

func opposite(a, b byte) bool {
	return a == b+32 || a == b-32
}

func react(polymer []byte, remove byte) []byte {
	var result []byte
	for _, unit := range polymer {
		switch {
		case unit == remove || unit == remove-32:
			continue
		case len(result) != 0 && opposite(result[len(result)-1], unit):
			result = result[:len(result)-1]
		default:
			result = append(result, unit)
		}
	}
	return result
}

func main() {
	polymer, err := ioutil.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}

	polymer = react(polymer, 0)
	fmt.Println(len(polymer))

	lengths := make(chan int)
	for unitType := 'a'; unitType <= 'z'; unitType++ {
		go func(unitType byte) { lengths <- len(react(polymer, unitType)) }(byte(unitType))
	}
	min := len(polymer)
	for unitType := 'a'; unitType <= 'z'; unitType++ {
		length := <-lengths
		if length < min {
			min = length
		}
	}
	fmt.Println(min)
}
