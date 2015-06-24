package main
import "fmt"

type Number int

//method inc has a receiver of type pointer to Number
func (n *Number) inc() {
    *n++
}

//method print has a receiver of type Number
func (n Number) print() {
    fmt.Println("The number is equal to", n)
}

func main() {

    i := Number(10) //say that i is of type Number and is equal to 10
    fmt.Println("i is equal to", i)

    fmt.Println("Let's increment it twice")
    i.inc() //same as (&i).inc() method expects a pointer, but that's okay
    fmt.Println("i is equal to", i)
    (&i).inc() //this also works as expected
    fmt.Println("i is equal to", i)

    p := &i //p is a pointer to i

    fmt.Println("Let's print it twice")
    p.print() //same as (*p).print() method expects a value, but that's okay
    i.print() //this also works as expected
}