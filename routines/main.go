package main

import (
  "fmt"
  "time"
)

func f(n int) {
  for i := 0; i < 10; i++ {
    time.Sleep(800 * time.Millisecond)
    fmt.Println(n, ":", i)
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
