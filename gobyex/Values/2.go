package main

import (
 "fmt"
 "io"
 "os"
)

func countChars(r io.Reader) int {
 buf := make([]byte, 16)
 total := 0
 for {
 n, err := r.Read(buf)
 if err != nil && err != io.EOF {
 return 0
 }
 if err == io.EOF {
 break
 }
 total = total + n
 }
 return total
}
