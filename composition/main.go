package main

import (
  "fmt"
)

type User struct{
  name, email string
  age int
}

type Student struct{
  User
  speciality string
}

func main(){
  student := Student{User{"Julio", "julio@example.com", 23}, "Engineer"}
  fmt.Println(student)
  fmt.Println(student.name)
  fmt.Println(student.speciality)
  student.age = 25
  fmt.Println(student.age)
}