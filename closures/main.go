package main

import (
  "fmt"
)

func makeEvenGenerator() func() int {
  i := 0
  return func() (int) {
    ret := i
    i += 2
    return ret
  }
}

func main() {
  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
  fmt.Println(increment())

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())
}