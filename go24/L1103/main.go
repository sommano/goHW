package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowFunc()
	fmt.Println("I am shown striaght away")
	time.Sleep(time.Second * 3)
}
