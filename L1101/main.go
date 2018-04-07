package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 3)
	fmt.Println("sleeper() finished")
}

func main() {
	slowFunc()
	fmt.Println("I am not shown until slowFunc() conpletes")
}
