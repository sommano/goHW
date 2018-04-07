package main

import (
	"fmt"
)

func main() {
	s := "After a backslash, certain single character escapes represent special values\nn is a line feed or new line \n\tt is a tab"
	fmt.Println(s)
}
