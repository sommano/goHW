package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdfl33l3"))
	fmt.Println(re.MatchString("roger"))

	fmt.Println(re.MatchString("iamthebestuseofthisappevaaaar"))

}
