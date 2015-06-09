package main

import "fmt"

func first() {
  fmt.Println("I'm the first function")
}
func second() {
  fmt.Println("I'm the second function.")
}
func main() {
  defer second()
  first()
}