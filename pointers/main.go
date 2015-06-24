package main

import "fmt"

type User struct{
  name string
}

func change_without_pointer(user User){
  user.name = "Piero"
}

func change_with_pointer(user *User){
  user.name = "Piero"
}

func main() {
  user := User{"Julio"}
  fmt.Println("Say hello to the pointer!: ", &user.name)
  fmt.Println(user)
  change_without_pointer(user)
  fmt.Println(user)
  change_with_pointer(&user)
  fmt.Println(user)
}
