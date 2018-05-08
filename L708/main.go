package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	fmt.Printf(reflect.TypeOf(a))
	fmt.Printf(reflect.TypeOf(b))

}
