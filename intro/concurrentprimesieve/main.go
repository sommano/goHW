//concurrent prime sieve

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

//copy the values from chanel "in" to channel "out"
//removing those divisbble by "prime"
func Filter(in <-chan int, out chan<- int, prime int) {
  for {
    i := <-in //receive value from "in"
    if i%prime != 0 {
      out <- i //send "i" to "out"
    }
  }
}

//the prime sieve:daisychain filter processes
func main() {
  ch := make(chan int) //create a new channel
  go Generate(ch) //launch generate goroutine
  for i :=0; i<10; i++ {
    prime := <-ch
    fmt.Println(prime)
    ch1 := make(chan int)
    go Filter(ch, ch1, prime)
    ch = ch1
  }
}
