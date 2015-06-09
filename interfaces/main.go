package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
}

type Rectangle struct {
  x, y float64
}

type Circle struct {
  r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
  return r.x * r.y
}

func SumTotalAreas(shapes ...Shape) float64{
  var totalarea float64
  for _, shape := range shapes{
    totalarea += shape.area()
  }
  return totalarea
}

func main() {
  c := Circle{r: 5} 
  r := Rectangle{x:3, y:6}
  fmt.Println("area of circle: ", c.area()) 
  fmt.Println("area of rectangle: ", r.area()) 
  fmt.Println("Sum of Areas: ", SumTotalAreas(&c, &r)) 
}