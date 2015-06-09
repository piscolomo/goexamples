package main
import (
  "fmt"
)

type SliceOfints []int
type AgesByNames map[string]int

func (s SliceOfints) sum() int {
    sum := 0
    for _, value := range s {
        sum += value
    }
    return sum
}

func (people AgesByNames) average() int{
  sum := 0
  for _, v := range people{
    sum = sum + v
  }
  return sum/len(people)
}

func main(){
    s := SliceOfints{1, 2, 3, 4, 5}
    people := AgesByNames {
        "Bob": 36,
        "Mike": 44,
        "Jane": 30,
        "Popey": 100,
    }
    fmt.Println("The sum of ints in the slice s is: ", s.sum())
    fmt.Println("The older in the map folks is:", people.average())
}